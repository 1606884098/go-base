package process

import (
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}
	info := fmt.Sprintf("用户id:%d对大家说:%s\n", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
