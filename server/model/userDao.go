package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个userDao结构体
//完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//服务器启动后就初始化一个userDao
//把他做成全局变量，在需要和redis操作时就直接使用
var (
	MyUserDao *UserDao
)

//使用工厂模式创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//第一个方法，根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id来redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示users哈希中没有找到对应的id
			err = ERROR_USER_NOTEXISTS

		}
		fmt.Println(err)
		return
	}
	//这里需要把res反序列化user实例
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json反序列化出错")
		return
	}
	return
}

//完成登录校验
//Login 完成对用户的校验
//如果用户的id和pwd都正确，则返回一个user实例
//如果id或pwd有错误，则返回对应信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("getUserById失败")
		return
	}
	//这个时候证明用户已经获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return

}
