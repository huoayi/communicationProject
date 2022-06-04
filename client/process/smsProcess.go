package process

import (
	"communicationProject/client/utils"
	common "communicationProject/common/message"
	"encoding/json"
	"fmt"
)

type sysProcess struct {
}

//发送群聊消息
func (this *sysProcess) SendGroupMes(content string) (err error) {
	//1 创建一个mes
	var mes common.Message
	mes.Type = common.SmsMesType
	var smsMes common.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserStatus

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败", err.Error())
		return
	}
	mes.Data = string(data)
	//对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err.Error())
		return
	}
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg错误,", err.Error())
		return
	}
	return
}
