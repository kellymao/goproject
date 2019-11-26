package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"study1/exmaple-more/chatroom/common"
)



type ChatClient struct{

	Chatsocket net.Conn
	Msg_stdin chan string
	Msg_send  chan string
	Msg_recv  chan string
	Login_status chan bool

}

func (chat *ChatClient) connect_server(addr string) bool {


	client,err:=net.Dial("tcp",addr)
	if err!=nil{
		log.Error("connect_server 连接服务器失败！",err)
		return false

	}

	chat.Chatsocket = client
	log.Debug(chat.Chatsocket)
	return true

}

func (chat *ChatClient) handle_stdin(){

	reader := bufio.NewReader(os.Stdin)
	for {

		line ,err:= reader.ReadString('\n')
		if err!=nil{
			log.Error("handle_stdin err:",err)
		}

		/*
		从终端接收输入， 去掉尾部的空格
		 */
		log.Debug(line)
		chat.Msg_stdin <- strings.TrimSuffix(line,"\n")
	}
}

func (chat *ChatClient) handle_message(userid string,ctx context.Context){

	for{

		select{
			case message := <-chat.Msg_send:
				chat.send_message(message)

			case message := <-chat.Msg_recv:
				chat.handle_stdout(message,userid)
			case <-ctx.Done():
				return
		}
	}

}

func (chat *ChatClient) send_message(message string){
	log.Debug("Msg_send收到了数据，准备发送数据", message)
	data, err := common.Encode(message)
	if err != nil {
		log.Error("encode data err:", err)
	}

	_, err = chat.Chatsocket.Write(data)

	if err != nil {
		log.Error("handle_send err:", err)
	}
}

func (chat *ChatClient) handle_stdout(message string,userid string){


	/*
		Msgtype string  `json:"msgtype"`
		Msgtxt string `json:"msgtxt"`
		Sendto string `json:"sendto"`
		From string   `json:"from"`
		Stauts bool `json:"status"`

	*/

	var msg common.Msginfo

	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Error("handle_msg error:", err)
	}

	switch msg.Msgtype {

	case "loginresp":

		log.Debug(msg.Msgtxt)
		if msg.Stauts {

			chat.Login_status <- true
		}
	case "message":
		log.Debug("sendto ", msg.Sendto)
		log.Debug("from ", msg.From)
		log.Debug("text ", msg.Msgtxt)

		if msg.Sendto == "all" {
			if msg.From == userid {
				fmt.Println("我对大家说：", msg.Msgtxt)

			} else {
				fmt.Println(msg.From, "对大家说：", msg.Msgtxt)

			}
		} else {
			fmt.Println(msg.From, "对我说：", msg.Msgtxt)
		}

	}
}





func (chat *ChatClient) handle_recv(ctx context.Context,cancal context.CancelFunc){

	log.Debug("socket client ",chat.Chatsocket)
	reader := bufio.NewReader(chat.Chatsocket)
	for {
		data, err := common.Decode(reader)
		if err != nil {
			log.Error("decode data err:", err)
			cancal()
			return
		}
		log.Debug("处理过的数据",data)
		chat.Msg_recv <- data
	}

}