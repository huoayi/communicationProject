package model

import (
	common "communicationProject/common/message"
	"net"
)

//因为在客户端很多地方会用到curUser,我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	common.User
}
