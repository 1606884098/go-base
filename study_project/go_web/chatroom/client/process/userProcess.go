package process

import (
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/client/utils"
	"go-base/study_project/go_web/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

//用户注册
func (this *UserProcess) Register(userId int, userPwd, userName string) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	//处理错误
	if err != nil {
		fmt.Println("连接服务器失败, err = ", err)
		return
	}
	defer conn.Close()

	//2.通过conn给服务器发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType

	//创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json marshal err = ", err)
		return
	}
	//将注册信息赋值给mes
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err = ", err)
		return
	}

	//创建一个Transfer对象
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送消息错误, err = ", err)
	}
	//3接收服务端返回的消息并解析
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err = ", err)
		return
	}

	var registerResMes message.RegisterResMes
	//将服务端返回的消息反序列化为registerResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,请重新登录!")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
}

//登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接服务端失败, err = ", err)
		return
	}
	defer conn.Close()

	var mes message.Message

	mes.Type = message.LoginMesType

	//创建一个登录类型的变量
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//对loginMes做序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err = ", err)
		return
	}
	//将data转换为string类型
	mes.Data = string(data)

	//对mes做序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err = ", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//将登录消息写入连接中
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("conn write fail, err = ", err)
		return
	}

	//等待服务端返回
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("conn read fail, err = ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		return
	}
	//登录成功
	if loginResMes.Code == 200 {
		fmt.Println("登录成功,当前在线用户列表如下:")

		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//用户登录后，通知那些人在线
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:", v)

			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}

		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	}

	return
}
