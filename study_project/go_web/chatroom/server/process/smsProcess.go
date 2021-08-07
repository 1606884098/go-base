package process2

import (
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/client/utils"
	"go-base/study_project/go_web/chatroom/common/message"
	"net"
)

type SmsProcess struct {
}

//转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err = ", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}

		this.SnedMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SnedMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败, err = ", err)
	}
}
