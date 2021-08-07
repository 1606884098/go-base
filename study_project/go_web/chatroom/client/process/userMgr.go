package process

import (
	"fmt"
	"go-base/study_project/go_web/chatroom/client/model"
	"go-base/study_project/go_web/chatroom/common/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CruUser

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status

	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}

//在客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:", id)
	}
}
