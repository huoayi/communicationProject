package process

import (
	"communicationProject/client/utils"
	"communicationProject/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//..暂时不需要字段
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的map
	//将消息转发
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &mes)
	if err != nil {
		fmt.Println("反序列化出错", err)
		return
	}
	data, err := json.Marshal(smsMes.Content)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMessageToEachOnlineUser(data, up.Conn)
	}
}
func (this *SmsProcess) SendMessageToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个tf实例发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发失败", err)
		return
	}

}
