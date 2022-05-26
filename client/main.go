package main

import (
	"communicationProject/client/login"
	"fmt"
)

var userId int
var userPwd string

func main() {
	var sel int          //选项
	var flag bool = true //判断是否继续循环
	fmt.Println("------------欢迎使用通信系统---------------")
	fmt.Println("\t\t1.登陆")
	fmt.Println("\t\t2.注册")
	fmt.Println("\t\t3.退出")
	fmt.Println("\t\t请选择1～3")
	fmt.Scanf("%d", &sel)
	for flag {
		switch sel {
		case 1:
			fmt.Println("登陆中")
			flag = false
		case 2:
			fmt.Println("注册中")
			flag = false
		case 3:
			fmt.Println("退出系统")
			flag = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	if sel == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)
		login.Login(userId, userPwd)

	} else if sel == 2 {
		fmt.Println("开始注册")
	}
}
