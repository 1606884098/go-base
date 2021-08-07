package main

import (
	"encoding/binary"
	"fmt"
	"go-base/study_project/go_web/chatroom/client/process"
)

func main() {
	fmt.Println("------------欢迎登陆聊天室-----------------")
	fmt.Println("\t\t\t1 登陆聊天室")
	fmt.Println("\t\t\t2 用户注册")
	fmt.Println("\t\t\t3 退出系统")
	fmt.Println("\t\t\t请选择(1-3):")
	//用于接收用户的输入
	var key int
	//用户id
	var userId int
	//用户密码
	var userPwd string
	//用户名
	var userName string
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("登陆聊天室")
		fmt.Println("请输入用户的id:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码:")
		fmt.Scanf("%s\n", &userPwd)
		up := &process.UserProcess{}
		up.Login(userId, userPwd)
	case 2:
		fmt.Println("注册用户")
		fmt.Println("请输入用户id:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入登陆密码:")
		fmt.Scanf("%s\n", &userPwd)
		fmt.Println("请输入用户名:")
		fmt.Scanf("%s\n", &userName)
		up := &process.UserProcess{}
		up.Register(userId, userPwd, userName)
	case 3:
		fmt.Println("3 退出系统")
	}
}

func main123() {
	i := uint32(12345)
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, i)
	fmt.Println(b)

	fmt.Println(uint32(binary.BigEndian.Uint32(b)))

}
