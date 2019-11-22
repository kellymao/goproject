package common

import (
	"fmt"
	"net"
)


var usermanger *UserProcess


type UserProcess struct{

	OnlineUsers map[string]net.Conn


}


func (u *UserProcess) Login(user map[string]string,conn net.Conn) bool {

	if user["userid"] == "1" && user["password"] == "2" {
		fmt.Println("登录成功了:",user["username"])

		u.OnlineUsers[user["userid"]] = conn

		return true


	}else{
		fmt.Println("用户名或密码不对")
		return false

	}

}


func (u *UserProcess) Logout(id string){
	delete(u.OnlineUsers,id)



}

func  (u *UserProcess) Sendmsg(id string,data []byte) error {

	connection ,ok := u.OnlineUsers[id]
	if ok {
		fmt.Println("向用户发送信息",id,string(data))
		_, err := connection.Write(data)
		if err!=nil{
			return err
		}
		return nil
	}else
	{
		return fmt.Errorf("用户已退出或者不存在 %s ",id)
	}
}



func NewUserProcess()*UserProcess{

	return &UserProcess{
		make(map[string]net.Conn ),
	}

}


