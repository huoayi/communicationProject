package process

import "fmt"

//因为userMgr实例在服务器有且仅有一个
//因为在很多地方都会使用到，因此将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcessor
}

//完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

//完成对onlineUsers的添加
func (this *UserMgr) AddOnlineUser(up *UserProcessor) {

	this.onlineUsers[up.UserId] = up
}

//完成对onlineUsers的删除
func (this *UserMgr) DeleteOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcessor {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcessor, err error) {

	//如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
