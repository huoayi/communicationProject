package main

import (
	"communicationProject/client/utils"
	common "communicationProject/common/message"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

func main() {
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
		go process(conn)
	}

}

//
//func readPkg(conn net.Conn) (message common.Message, err error) {
//	buf := make([]byte, 8096)
//	fmt.Println("读取客户端发来的数据")
//	//conn.Read在conn没有被关闭的情况下才会阻塞
//	//如果客户端关闭了conn则不会出现阻塞
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//		fmt.Println("conn.Read错误\t", err)
//		return
//	}
//	//根据buf[:4]转化成1个uint32类型
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4])
//	//根据pkgLen读取消息内容
//	n, err := conn.Read(buf[:pkgLen])
//	if n != int(pkgLen) || err != nil {
//		err = errors.New("handler错误")
//		return
//	}
//	//把pkgLen反序列化成->common.Message
//	err = json.Unmarshal(buf[:pkgLen], &message)
//	if err != nil {
//		err = errors.New("body错误")
//		return
//	}
//	return
//}

func process(conn net.Conn) {
	defer conn.Close()
	//读取客户端发送的信息
	for {

		//将读取数据包直接封装成1个函数
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出，服务端也退出..")
				return
			} else {
				err = errors.New("read package err")
				fmt.Println(err)
				return
			}
		}
		fmt.Println("mes=", mes)
		err = serverProcessLogin(conn, &mes)
		if err != nil {
			return
		}
	}
}

//编写一个serverProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		//处理登陆
	case common.LoginResMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

func serverProcessLogin(conn net.Conn, mes *common.Message) (err error) {
	//核心代码
	//先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unarshal fail err = ", err)
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
	err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("writePkg err", err)
		return

	}
	return
}

//func writePkg(conn net.Conn, data []byte) (err error) {
//
//	//先发送一个长度给对方
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var byte [4]byte
//	binary.BigEndian.PutUint32(byte[0:4], pkgLen)
//	n, err := conn.Write(byte[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write fail")
//		return
//	}
//	//发送data本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write fail")
//		return
//	}
//	return
//}
