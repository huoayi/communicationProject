package process

import (
	"communicationProject/client/model"
	common "communicationProject/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int]*common.User = make(map[int]*common.User, 10)
var curUser model.CurUser //我们在用户登陆成功后完成对curUser的初始化

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(mes *common.NotifyUserStatusMes) {
	user, ok := onlineUsers[mes.UserId]
	if !ok {
		user = &common.User{
			UserId: mes.UserId,
		}
	}
	user.UserStatus = mes.Status
	onlineUsers[mes.UserId] = user
	OutputOnlineUsers()
}

//显示当前客户端在线用户
func OutputOnlineUsers() {
	fmt.Println("当前用户列表")
	for id := range onlineUsers {
		fmt.Println("用户id\t", id)

	}
}
