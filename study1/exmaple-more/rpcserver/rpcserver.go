package main

import (
	"fmt"
	"os"
	"time"
)

func rpcclient(msgchan chan string,name string){

	str:= fmt.Sprintf(" i'm is client  %s" ,name)
	msgchan <- str

	go func(){

		for {
			txt :=<- msgchan
			fmt.Println(name," recieve the server message ",txt)

		}

	}()



}


func main(){


	msgchan := make(chan string)

	go rpcclient(msgchan,"client1")

	//go rpcclient(msgchan,"client2")

	for {


		select {
		    case msg := <-msgchan:
				fmt.Println("服务器收到了消息", msg)

				msgchan <- "recv ok"

			case <- time.After(time.Second*10):

				os.Exit(0)



		}

	}




}