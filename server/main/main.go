package main

import (
	"communicationProject/server/model"
	"fmt"
	"io"
	"net"
)

func main() {
	model.InitPool("localhost:6379", 8, 0, 10)
	initUserDao()
	lis, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听错误")
		return
	}
	defer lis.Close()
	//监听成功，等待客户端来连接服务器
	for {
		fmt.Println("监听成功")
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("连接错误")
			return
		}
		//连接成功则立即启动协程与客户端保持通讯
		go Process(conn)
	}

}

func Process(conn net.Conn) {
	defer conn.Close()
	//读取客户端发送的信息
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process()
	if err != nil && err != io.EOF {
		fmt.Println("process出了点小问题，客户端与服务器通讯协程错误", err)
		return
	}
}

//这里编写一个函数完成对userDao的初始化任务
func initUserDao() {
	//这里的model.Pool本身是一个全局变量
	//这里需要注意一个初始化顺序问题
	//initPool，在initUserDao
	model.MyUserDao = model.NewUserDao(model.Pool)
}
