package process

import (
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/client/utils"
	"go-base/study_project/go_web/chatroom/common/message"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-------------恭喜您登录成功!-----------------")
	fmt.Println("1 显示在线用户列表")
	fmt.Println("2 发送消息")
	fmt.Println("3 退出系统")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("你想对大家说什么:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("您输入的选线不正确...")
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个Transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息!")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf ReadPkg err = ", err)
			continue
		}
		//判断消息的类型
		switch mes.Type {
		//有人上线了
		case message.NotifyUserStatusMesType:
			fmt.Println("===========有人上线了!================")
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: //群聊
			outputGroupMes(&mes)
		default:
			fmt.Println("服务端返回了未知的消息类型!")
		}

	}

}
