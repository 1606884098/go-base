package Bridge

import "fmt"

type UrgencyMessage struct {
	method MessageImlementer
}

func NewUrencyMessage(method MessageImlementer) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}
func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("发送到[%s]", text), to) //h很快速度发送
}
