package main

import "fmt"
import "study1/example-chat/common"
import "encoding/json"
import "net"
import "io"
import "os"
import "strings"
import "bufio"


var userid string

var pwd string



var userinfo common.User


var chan_write chan common.Message = make(chan  common.Message)

var chan_exit chan int = make(chan int)

func prompt(){


	fmt.Println("欢迎来到聊天室，请先登录聊天室：")

	fmt.Println("请输入你的聊天id：")

	fmt.Scanln(&userid)


	fmt.Println("请输入你的密码：")

	fmt.Scanln(&pwd)




}

/*
func process_talk(){


	fmt.Println("选择哪个聊天室，1：公共聊天室 2. 找好友单聊")

	var sel int
	fmt.Scanln(&sel)

	if sel == "1" {

		fmt.Printf("%s : 欢迎 %s 来到公用聊天室",time.Now().Format("02/1/2006 15:04"),userinfo.Name)

		// 获取在线人的列表

		var msg common.Message =  common.Message{

			Msgtype:"getonlinelist",
			Msginfo:"",
		}


		chan_write<-msg


		process_stdin("allmessage")

	}

	if sel == "2" {

		fmt.Printf("%s : 欢迎 %s 来单聊",time.Now().Format("02/1/2006 15:04"),userinfo.Name)

		// 显示好友的列表

		data,err:= json.Marshal(userinfo.Friends.User_id)
		common.Check_err(err)

		var msg common.Message =  common.Message{}

		msg.Msgtype = "getfriends"
		msg.Msginfo = string(data)



		chan_write<-msg


		process_stdin("singlemessage")


	}

}
*/

func process_stdin(msgtype string){

	var msg common.Message = common.Message{}
	msg.Msgtype=msgtype

	for{
	inputReader := bufio.NewReader(os.Stdin)

	input,err := inputReader.ReadString('\n')
	if err!=nil{
		fmt.Println(err)
		chan_exit<-1
	}

	if input == "close"{

		chan_exit<-1
	}

	msg.Msginfo = input
	msg.Msgsend = userinfo.Id

	if msgtype == "singlemessage"{

		msg.Msgrecv = strings.TrimLeft(strings.Fields(input)[0],"@")

	}else{

		msg.Msgrecv = "all"
	}
	 // @13 你好 === 取出来13
	chan_write<-msg

	}



}
/*
type User struct {

	Id int
	Pwd string
	Name string
	Status string
	Friends Friend

*/


func logininfo() (data []byte){

	var user common.User = common.User{
		Id: userid,
		Pwd:pwd,


	}

	user_str,err := json.Marshal(user)
	common.Check_err(err)





	var msg common.Message  = common.Message{

		Msgtype : "login",
		Msginfo: string(user_str),


	}


	data,err = json.Marshal(msg)

	common.Check_err(err)


	return




}


func send_data(conn net.Conn, data []byte){


	_, err := conn.Write(data)
	common.Check_err(err)


}


func rev_data(conn net.Conn) []byte {


	var buf [512]byte

	var bufdata []byte
	for {
		n, err := conn.Read(buf[0:])  //读取内容时 没有内容会一直等待

		if err!=nil{

			fmt.Println(err)
			break


		}

		bufdata = append(bufdata,buf[0:n]...)

		if err == io.EOF {

			break


		}

		if n < 512 {              // 小于512 说明接受的内容已经接受完毕，可以退出循环

			break
		}



		//rAddr := conn.RemoteAddr()
		//fmt.Println("Receive from server", rAddr.String(), string(buf[0:n]))

		//common.Check_err(err)






	}


	return bufdata


}

