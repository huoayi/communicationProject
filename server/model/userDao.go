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

//第一个方法，根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id来redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示users哈希中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//这里需要把res反序列化user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json反序列化出错")
		return
	}
	return
}
