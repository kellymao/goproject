package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"study1/exmaple-more/chatroom/common"
)

var msgstdinchan chan string = make(chan string)
var msgsendchan chan string = make(chan string)
var msgrecvchan = make(chan string)

var login_resp = make(chan bool)

var my_userid string


func handle(conn net.Conn){




}


func client() (net.Conn, bool) {



	conn,err:=net.Dial("tcp","127.0.0.1:10086")

	if err!=nil{

		fmt.Println("连接服务器失败")
		return nil,false
	}

	return conn,true




}


func handle_stdin(){


	reader :=bufio.NewReader(os.Stdin)

	for {

		line ,err := reader.ReadString('\n')

		if err !=nil{
			fmt.Println("从终端读取数据失败")
			os.Exit(1)
		}

		fmt.Println(line)
		msgstdinchan <- strings.TrimSuffix(line,"\n")

	}
}

func recvmsg(conn net.Conn){

	var msg bytes.Buffer
	var buf  [7]byte

	for {
		n ,err := conn.Read(buf[0:])
		if err!=nil{


			if err == io.EOF{
				fmt.Println("服务器已断开连接")
				return
			}
			fmt.Println("从服务器接收数据失败",err)
			continue

		}
		fmt.Println("收到了数据长度为",n)

		msg.Write(buf[0:n])

		if buf[n-1] == '\n'{
			msgrecvchan <- msg.String()
			msg = bytes.Buffer{}
		}


	}


}

func handle_message(data string){

	var msg common.Msginfo
	err := json.Unmarshal([]byte(data),&msg)
	if err!=nil{
		fmt.Println("handle_message unmarshal err",err)
	}

	switch msg.Msgtype {

	case "loginresp":
		fmt.Println(msg.Msgtxt)
			if msg.Stauts {
				login_resp <- true

			}else{
				login_resp <- false
			}




	case "message":

		if msg.Sendto == "all"{

			if msg.From == my_userid{
				fmt.Printf("我对大家说：%s \n",msg.Msgtxt)
				return
			}
			fmt.Printf("%s 对大家说：%s \n",msg.From,msg.Msgtxt)

		}else{
			fmt.Printf("%s 对我说：%s \n",msg.From,msg.Msgtxt)

		}

	}


}

func handleconn(conn net.Conn)  {

	for {
		select {

		case data := <-msgsendchan:

			fmt.Println("before send", data)
			_, err := conn.Write([]byte(data + "\n"))
			if err != nil {
				fmt.Println("数据发送失败", err)

			}

		case data1 := <-msgrecvchan:

			//fmt.Println(data1)

			handle_message(data1)


		}

	}





}

func login() bool {

	fmt.Println("请输入登录信息")
	fmt.Println("请输入用户id：")
	userid := <-msgstdinchan
	fmt.Println("请输入密码：")
	pwd:= <-msgstdinchan

	if len(userid ) == 0 || len(pwd) ==0 {

		fmt.Println("用户名密码不可以为空")
		login()
	}

	/*
		Msgtype string  `json:"msgtype"`
		Msgtxt string `json:"msgtxt"`
		Sendto string `json:"sendto"`
		From string   `json:"from"`
	 */
	Msgtxt := map[string]string{"userid":userid ,"password":pwd}

	fmt.Println(Msgtxt)

	txt,err:=json.Marshal(Msgtxt)

	if err!=nil{
		fmt.Println("marshal err",err)
	}
	msg:= common.Msginfo{
		Msgtype:"login",
		Msgtxt:string(txt),
		Sendto:"server",
	}
	fmt.Println(msg)

	data,err := json.Marshal(msg)
	if err!=nil{
		fmt.Println("marshal err",err)
	}
	fmt.Println(string(data))
	msgsendchan <- string(data)

	status:= <- login_resp

	if status{

		fmt.Printf("登录成功了，[%s]\n" ,userid)
		my_userid =userid

	}else{
		fmt.Printf("登录失败了，[%s]\n" ,userid)


	}

	return status




}

func logicview(){

	fmt.Println("欢迎来到聊天室：")

	fmt.Println("公共聊天@all + 聊天内容， 单聊@userid + 聊天内容，退出 logout")

	for{

		message  := <-msgstdinchan

		if len(message ) == 0 {
			fmt.Println("输入内容不可以为空")
			continue
		}

		if strings.HasPrefix(message,"@all"){

			msg:= common.Msginfo{
				Msgtype:"message",
				Msgtxt:message,
				Sendto:"all",
			}


			data,err := json.Marshal(msg)
			if err!=nil{
				fmt.Println("marshal err",err)
			}

			msgsendchan <- string(data)


		}else if strings.HasPrefix(message,"@"){

			msglist :=strings.Split(message," ")
			txt:=strings.Join(msglist[1:],"")

			msg:= common.Msginfo{
				Msgtype:"message",
				Msgtxt:txt,
				Sendto:strings.TrimPrefix(msglist[0],"@"),
			}


			data,err := json.Marshal(msg)
			if err!=nil{
				fmt.Println("marshal err",err)
			}

			msgsendchan <- string(data)

		}else if strings.HasPrefix(message,"logout"){


			msg:= common.Msginfo{
				Msgtype:"logout",
				Msgtxt:"",
				Sendto:"",
			}


			data,err := json.Marshal(msg)
			if err!=nil{
				fmt.Println("marshal err",err)
			}

			msgsendchan <- string(data)

			fmt.Println("退出客户端")
			wg.Done()
			return

		} else{
			fmt.Println("输入的格式不对")
			fmt.Println("公共聊天@all + 聊天内容， 单聊@userid + 聊天内容：")
			continue

		}

	}



}
