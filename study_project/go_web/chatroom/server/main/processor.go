package main

import (
	"fmt"
	"go-base/study_project/go_web/chatroom/common/message"
	"go-base/study_project/go_web/chatroom/server/process"
	"go-base/study_project/go_web/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) process2() (err error) {
	//循环的读取客户端发送的消息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出..")
				return err
			} else {
				fmt.Println("readPkg err = ", err)
				return err
			}
		}
		//fmt.Println("mes = ", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}

//根据客户端发送的消息种类的不同，决定调用哪个函数进行处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes = ", mes)
	switch mes.Type {
	//处理登录的业务逻辑
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		up.ServerProcessLogin(mes)
	//处理注册的业务逻辑
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不正确，无法处理....")
	}
	return
}
