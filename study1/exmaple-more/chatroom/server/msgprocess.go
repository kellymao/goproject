package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"study1/exmaple-more/chatroom/common"
)

type Msginfo struct{

	Msgtype string  `json:"msgtype"`
	Msgtxt string `json:"msgtxt"`
	Sendto string `json:"sendto"`
	From string   `json:"from"`
	Status bool `json:"status"`

}

type MsgProcess struct{


	Msginfo Msginfo
	Connection net.Conn
	Ownid   string


}

func (m *MsgProcess) Login(){

		var msgtxt string
		var status bool
		var user map[string]string = map[string]string{"userid":"","password":""}
		_ = json.Unmarshal([]byte(m.Msginfo.Msgtxt), &user)

		ok := usermanger.Login(user,m.Connection)
		if ok {
			m.Ownid = user["userid"]

			msgtxt = "登录成功：" + user["userid"]

			var k_list = make([]string,0)
			for k := range usermanger.OnlineUsers{
				k_list = append(k_list,k)
			}

			msgtxt+="当前在线客户列表："
			msgtxt+= strings.Join(k_list,"、")
			status = true

		}else{
			msgtxt = "登录失败，用户名或密码不对：" + user["userid"]
			status = false

		}



	   msg := Msginfo{
		Msgtype:"loginresp",
		Msgtxt:msgtxt ,
		From:"server",
		Status:status,
	   }

		data, _ := json.Marshal(msg)
	    returnmsg, err := common.Encode(string(data))
	    if err!=nil{
	    	fmt.Println("encode data err",err)
		}

		fmt.Println("登录的信息",string(returnmsg))
		_, _ = m.Connection.Write(returnmsg)



}

func (m *MsgProcess) Logout(){
	usermanger.Logout(m.Ownid)
}

/*
func (m *MsgProcess) Sendmsg(data []byte){

	_, err := m.Connection.Write(data)
	if err!=nil{
		fmt.Println("消息发送失败",data)
	}



}
*/

func (m *MsgProcess) Handle(){


	fmt.Printf("本次的信息%+v \n",m)
	var err error
	callmsg := Msginfo{
		Msgtype:m.Msginfo.Msgtype,
		Msgtxt:m.Msginfo.Msgtxt ,
		From:m.Ownid,

	}



	sendto_id := m.Msginfo.Sendto


	if sendto_id == "all"{

		callmsg.Sendto = "all"
		fmt.Println("要发送给客户端的数据:",callmsg)
		data ,_ := json.Marshal(callmsg)

		for uid := range usermanger.OnlineUsers{

				go  usermanger.Sendmsg(uid, data)



		}

	}else{

		callmsg.Sendto = sendto_id
		fmt.Println("要发送给客户端的数据:",callmsg)
		data ,_ := json.Marshal(callmsg)

		err = usermanger.Sendmsg(sendto_id, data)

	}

	if err!=nil{

		callmsg = Msginfo{
			Msgtype:m.Msginfo.Msgtype,
			Msgtxt:err.Error() ,
			From:"server",
		}

		data ,_ := json.Marshal(callmsg)
		_ = usermanger.Sendmsg(m.Ownid, data)

	}




}

func (m *MsgProcess) HandleMsg(){


	switch m.Msginfo.Msgtype{
	case "login":
		m.Login()
		fmt.Println("----login after")
		fmt.Printf("%+v \n",m)

	case "logout":
		m.Logout()


	case "message":
		m.Handle()






	}


}