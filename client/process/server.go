package process

import (
	"communicationProject/client/utils"
	"communicationProject/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-------------登陆成功页面----------------")
	fmt.Println("\t\t1.显示在线用户列表")
	fmt.Println("\t\t2.发送消息")
	fmt.Println("\t\t3.消息列表")
	fmt.Println("\t\t4.退出系统")
	fmt.Println("请选择 1～4")
	var key int
	var content string
	fmt.Scanf("%d\n", &key)
	sysProcess := &sysProcess{}
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		OutputOnlineUsers()
	case 2:
		fmt.Println("发送消息")
		fmt.Scanf("%s\n", &content)
		sysProcess.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误")

	}
}

//和服务器保持通讯
func ServerProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发来的消息
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for true {
		mes, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("ReadPkg错误 ", err)
			return
		}
		fmt.Printf("mes = %v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线
			//1 取出NotifyUserStatusMes

			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2 把用户的信息和状态保存在map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: //有人群发消息

		default:
			fmt.Println("服务器端返回了未知类型，暂时不能处理")
		}
	}
}
