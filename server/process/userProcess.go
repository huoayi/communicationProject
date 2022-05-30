package process

import (
	common "communicationProject/common/message"
	"communicationProject/server/model"
	"communicationProject/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn net.Conn
}

func (this *UserProcessor) ServerProcessLogin(mes *common.Message) (err error) {
	//核心代码
	//先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}
	//先声明一个resMes
	var resMes common.Message
	resMes.Type = common.LoginResMesType
	//再声明一个LoginResMes,并完成赋值
	var LoginResMes common.LoginResMes
	//需要去redis数据库进行校验
	//使用model.MyUserDao到redis进行验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			LoginResMes.Code = 500
			LoginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			LoginResMes.Code = 300
			LoginResMes.Error = err.Error()
		} else {
			LoginResMes.Code = 404
			LoginResMes.Error = "服务器内部错误"
		}

		//这里测试成功后返回具体错误信息
	} else {
		LoginResMes.Code = 200
		fmt.Println(user.UserName, "登陆成功")
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg err", err)
		return

	}
	return
}
