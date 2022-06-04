package utils

import (
	common "communicationProject/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时使用的缓冲
}

func (this *Transfer) ReadPkg() (message common.Message, err error) {

	fmt.Println("读取客户端发来的数据")
	//conn.Read在conn没有被关闭的情况下才会阻塞
	//如果客户端关闭了conn则不会出现阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read错误\t", err)
		return
	}
	//根据buf[:4]转化成1个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("handler错误")
		return
	}
	//把pkgLen反序列化成->common.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &message)
	if err != nil {
		err = errors.New("body错误")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail")
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("+++++++++++++++++++++")
		fmt.Println("conn.Write fail")
		return
	}
	return
}
