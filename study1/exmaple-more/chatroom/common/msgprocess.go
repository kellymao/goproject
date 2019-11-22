package common

import (
	"encoding/json"
	"fmt"
	"net"
)

type Msginfo struct{

	Msgtype string  `json:"msgtype"`
	Msgtxt string `json:"msgtxt"`
	Sendto string `json:"sendto"`
	From string   `json:"from"`
	Stauts bool `json:"status"`

}

type MsgProcess struct{


	Msginfo Msginfo
	Connection net.Conn
	Ownid   string


}

func (m MsgProcess) Login(){

		var msgtxt string
		var user map[string]string = map[string]string{"userid":"","password":""}
		_ = json.Unmarshal([]byte(m.Msginfo.Msgtxt), &user)

		ok := usermanger.Login(user,m.Connection)
		if ok {
			m.Ownid = user["userid"]

			msgtxt = "登录成功：" + user["userid"]

		}else{
			msgtxt = "登录失败，用户名或密码不对：" + user["userid"]

		}

	   msg := Msginfo{
		Msgtype:"loginresp",
		Msgtxt:msgtxt ,
		From:"server",
	   }

		data, _ := json.Marshal(msg)
	_ = usermanger.Sendmsg(m.Ownid, data)



}

func (m MsgProcess) Logout(){
	usermanger.Logout(m.Ownid)
}

func (m MsgProcess) Sendmsg(data []byte){
	_, err := m.Connection.Write(data)
	if err!=nil{
		fmt.Println("消息发送失败",data)
	}



}

func (m MsgProcess) Handle(){
	callmsg := Msginfo{
		Msgtype:m.Msginfo.Msgtype,
		Msgtxt:m.Msginfo.Msgtxt ,
		From:m.Ownid,
	}

	data ,_ := json.Marshal(callmsg)
	sendto_id := m.Msginfo.Sendto

	err := usermanger.Sendmsg(sendto_id, data)

	if err!=nil{

		callmsg = Msginfo{
			Msgtype:m.Msginfo.Msgtype,
			Msgtxt:err.Error() ,
			From:"server",
		}

		data ,_ = json.Marshal(callmsg)
		_ = usermanger.Sendmsg(m.Ownid, data)

	}




}

func (m MsgProcess) HandleMsg(){

	switch m.Msginfo.Msgtype{
	case "login":
		m.Login()
	case "logout":
		m.Logout()
	case "message":
		m.Handle()






	}


}