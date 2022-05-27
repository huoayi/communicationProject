package process

import (
	"communicationProject/client/utils"
	common "communicationProject/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//制定协议
	//连接到服务器
	conn, err := net.Dial("tcp", "192.168.31.101:8888")
	if err != nil {
		fmt.Println("连接失败\t", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//准备通过conn发送消息给服务
	var mes common.Message
	mes.Type = common.LoginMesType
	//创建1个loginMes结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化出错")
		return
	}
	//把data赋给mes.Data
	mes.Data = string(data)
	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	//data就是我们要发送的消息
	//先把data的长度发送给服务器
	//先获取data的长度，转化成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail")
		return
	}
	fmt.Println(string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("con.Write(data)失败")
	}
	time.Sleep(5 * time.Second)
	fmt.Println("休眠5秒")
	//这里处理服务端返回的消息
	transfer := &utils.Transfer{
		Conn: conn,
	}
	mes, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg错误", err)
		return
	}
	//将mes的data部分反序列化成LoginResMes
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
		//起一个协程为了保持与服务器端端通讯
		//如果服务器有数据推送可以接受并显示在客户端的终端
		go ServerProcessMes(conn)
		//1.显示登陆成功的菜单
		for true {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return

}
