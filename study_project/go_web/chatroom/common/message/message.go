package message

const (
	RegisterMesType         = "RegisterMes"
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//用户状态
const (
	UserOnline = iota
	UserOffline
)

type Message struct {
	//消息的类型
	Type string `json:"type"`
	//消息的内容
	Data string `json:"data"`
}

type RegisterMes struct {
	User User `json:"user"`
}

//注册返回结构体
type RegisterResMes struct {
	Code  int    `json:"code"`  //状态码
	Error string `json:"error"` //错误信息
}

type LoginMes struct {
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
}

type LoginResMes struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	UsersId []int  //保存用户id的切片
}

//用于推送用户状态
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//群聊结构体
type SmsMes struct {
	Content string `json:"content"` //内容
	User
}
