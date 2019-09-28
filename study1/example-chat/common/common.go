package common


import "os"
import "fmt"

const (

	Online = "online"
	Offline = "offline"



)

const(
	Msgtype_login="login"
	Msgtype_loginresp="loginresp"

	Msgtype_getonlinelist = "getonlinelist"

	Msgtype_getfriends = "getfriends"


	Msgtype_allmessage="allmessage"
	Msgtype_singlemessage = "singlemessage"

	Msgtype_notice="notice"

)


func Check_err(err error){

	if err!=nil{

		fmt.Println(err)
		os.Exit(1)
	}


}


type Message struct{

	Msgtype string
	Msginfo string
	Msgsend string
	Msgrecv  string

}


//登录的返回 status 是是否成功， respbody是user 的信息

type Resp struct {

	Status bool
	RespBody string



}



