package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"study1/exmaple-more/chatroom/common"
)




func init(){

	usermanger = NewUserProcess()
}



func handle(conn net.Conn){
	reader := bufio.NewReader(conn)
	var msghander *MsgProcess  = &MsgProcess{}


	var msginfo Msginfo

	defer func(){
		conn.Close()
		if msghander.Ownid != "" {
				delete(usermanger.OnlineUsers,msghander.Ownid)
		}

	}()

	/*
	for {

		line,err:=reader.ReadString('\n')
		if err!=nil{

			fmt.Println(err)
			if err != io.EOF{
				fmt.Println("数据接收错误",err)
				return
			} else {
				fmt.Println("读取数据结束",err) // 客户端取消时，会收到EOF
				return
			}

		}


		line = strings.TrimSuffix(line,"\n")
		fmt.Println("读取到的数据",line)

		err = json.Unmarshal([]byte(line), &msginfo)
		if err!=nil{
			fmt.Println("数据unmarshal 失败了",err)
		}

		msghander.Connection = conn
		msghander.Msginfo = msginfo

		fmt.Printf("%+v \n",msghander)
		msghander.HandleMsg()
		fmt.Println("-------------------")
		fmt.Printf("%+v \n",msghander)



	}
	*/


	for {

		data,err:=common.Decode(reader)
		if err!=nil{
			fmt.Println(err)
			return
		}

		err = json.Unmarshal([]byte(data), &msginfo)
		if err!=nil{
			fmt.Println("数据unmarshal 失败了",err)
		}

		msghander.Connection = conn
		msghander.Msginfo = msginfo

		msghander.HandleMsg()
		fmt.Println("-------------------")
		fmt.Printf("%+v \n",msghander)


	}


}

func main(){

	listener, err := net.Listen("tcp", "127.0.0.1:10086")

	if err!=nil{
		fmt.Println("监听端口失败",err)
		return
	}



	for {
		conn,err:=listener.Accept()

		if err!=nil{
			fmt.Println("accept 失败",err)
			continue
		}

		go handle(conn)



	}




}
