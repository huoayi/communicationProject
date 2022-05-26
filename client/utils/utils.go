package utils

import (
	common "communicationProject/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func ReadPkg(conn net.Conn) (message common.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发来的数据")
	//conn.Read在conn没有被关闭的情况下才会阻塞
	//如果客户端关闭了conn则不会出现阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read错误\t", err)
		return
	}
	//根据buf[:4]转化成1个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("handler错误")
		return
	}
	//把pkgLen反序列化成->common.Message
	err = json.Unmarshal(buf[:pkgLen], &message)
	if err != nil {
		err = errors.New("body错误")
		return
	}
	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var byte [4]byte
	binary.BigEndian.PutUint32(byte[0:4], pkgLen)
	n, err := conn.Write(byte[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail")
		return
	}
	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write fail")
		return
	}
	return
}
