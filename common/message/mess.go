package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMeType  = "RegisterMes"
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
	Code  int    //状态码
	Error string //返回错误信息
}
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}
type RegistResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
