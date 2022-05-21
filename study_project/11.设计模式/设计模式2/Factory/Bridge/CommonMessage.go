package Bridge

type ComonMessage struct {
	method MessageImlementer
}

func NewComonMessage(method MessageImlementer) *ComonMessage {
	return &ComonMessage{method: method}
}
func (com *ComonMessage) SendMessage(text, to string) {
	com.method.Send(text, to)
}
