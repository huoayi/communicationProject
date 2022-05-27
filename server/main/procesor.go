package main

import (
	common "communicationProject/common/message"
	"communicationProject/server/process"
	"communicationProject/server/utils"
	"errors"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//编写一个serverProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) ServerProcessMes(mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		//处理登陆
		userProcessor := &process.UserProcessor{
			Conn: this.Conn,
		}
		err = userProcessor.ServerProcessLogin(mes)
	case common.LoginResMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

func (this *Processor) process() (err error) {
	for {
		UT := utils.Transfer{
			Conn: this.Conn,
		}
		//将读取数据包直接封装成1个函数
		mes, err := UT.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出，服务端也退出..")
				return err
			} else {
				err = errors.New("read package err")
				fmt.Println(err)
				return err
			}
		}
		fmt.Println("mes=", mes)

		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
	return nil
}
