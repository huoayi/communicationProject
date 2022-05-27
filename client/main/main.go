package main

import (
	"communicationProject/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	var sel int //选项
	fmt.Println("------------欢迎使用通信系统---------------")
	fmt.Println("\t\t1.登陆")
	fmt.Println("\t\t2.注册")
	fmt.Println("\t\t3.退出")
	fmt.Println("\t\t请选择1～3")
	fmt.Scanf("%d", &sel)
	for true {
		switch sel {
		case 1:

			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("登陆中")
			//完成登陆
			//创建1个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册中")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
