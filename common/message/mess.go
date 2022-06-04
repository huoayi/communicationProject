package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMeType          = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//这里定义登陆状态常量
const (
	UserOnline = iota
	UserOffice
	UserBusys
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//定义两个消息，后面可自由拓展
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}
type LoginResMes struct {
	Code    int    //状态码
	UserIds []int  //增加字段，保存用户id的切片
	Error   string //返回错误信息
}
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}
type RegistResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//为了配合服务器端推送用户变化状态的信息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户ID
	Status int `json:"status"` //用户状态
}

//增加一个SmsMes
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体
}
