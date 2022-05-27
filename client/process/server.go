package process

import (
	"communicationProject/client/utils"
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
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
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
	}
}
