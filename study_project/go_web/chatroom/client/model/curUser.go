package model

import (
	"go-base/study_project/go_web/chatroom/common/message"
	"net"
)

type CruUser struct {
	Conn net.Conn
	message.User
}
