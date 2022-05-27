package process

import (
	common "communicationProject/common/message"
	"communicationProject/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
}

func (this *UserProcessor) ServerProcessLogin(mes *common.Message) (err error) {
	//核心代码
	//先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}
	//先声明一个resMes
	var resMes common.Message
	resMes.Type = common.LoginResMesType
	//再声明一个LoginResMes,并完成赋值
	var LoginResMes common.LoginResMes

	//如果用户id = 100 密码 = 123456认为合法，否则非法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		LoginResMes.Code = 200
	} else {
		//非法
		LoginResMes.Code = 500
		LoginResMes.Error = "表示该用户不存在"
	}

	//将loginResMes序列化
	data, err := json.Marshal(LoginResMes)
	if err != nil {
		fmt.Println("json.Marshal失败")
		return
	}
	//将Data赋值给resMes
	resMes.Data = string(data)
	//对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//发送data 我们将他封装到writePkg函数中
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg err", err)
		return

	}
	return
}
