package process2

import (
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/common/message"
	"go-base/study_project/go_web/chatroom/server/models"
	"go-base/study_project/go_web/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

//注册
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	//取出mes的data字段并反序列化成RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmarshal fail, err = ", err)
		return
	}

	var resMes message.Message
	resMes.Type = "RegisterResMes"
	var registerResMes message.RegisterResMes

	//MVC
	//注册
	err = models.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == models.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = models.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
		fmt.Println("注册失败!")
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功!")
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json Marshal fail, err = ", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail, err = ", err)
		return
	}

	//创建一个Transfer对象
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	//发送data给服务器
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//将客户端传递过来的数据发序列化成loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshal fail err = ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	//登录
	_, err = models.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == models.ERROR_USER_NOEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == models.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误!"
		}
	} else {
		loginResMes.Code = 200

		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//将当前在线用户的id追加到loginResMes.UsersId
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		//当一个新的用户上线后，其他已经登录的用户也能获取最新在线用户列表
		this.NotifyOtherOnlineUser(loginMes.UserId)
	}
	//对loginResMes进行序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json Marshal fail, err = ", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail, err = ", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	return
}

//通知所有其他在线用户
//userId是上线的用户
func (this *UserProcess) NotifyOtherOnlineUser(userId int) {
	//遍历onlineUsers，然后一个个的去通知
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

//UserProcess.UserId = 5
//userId = 6
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes进行序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json Marshal fail, err = ", err)
		return
	}
	mes.Data = string(data)

	//对mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal fail, err = ", err)
		return
	}

	//创建Transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err = ", err)
		return
	}
}
