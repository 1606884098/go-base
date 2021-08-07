package process

import (
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/client/utils"
	"go-base/study_project/go_web/chatroom/common/message"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message

	mes.Type = message.SmsMesType

	//定义一个群聊的实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json marshal fail, err = ", err)
		return
	}
	mes.Data = string(data)

	//对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes mes json marshal fail, err = ", err)
		return
	}

	//将mes发送到服务端
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send group message fail, err = ", err)
		return
	}
	return
}
