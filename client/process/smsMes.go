package process

import (
	common "communicationProject/common/message"
	"encoding/json"
	"fmt"
)

func OutputGroupMes(mes *common.Message) { //一定要是SmsMes
	//显示即可
	//1 反序列化
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), smsMes)
	if err != nil {
		fmt.Println("Unmarshal失败", err)
		return
	}
	info := fmt.Sprintf("用户\t%d\t发送了\t%s\n", smsMes.UserId, smsMes.Content)
	fmt.Println(info)

}
