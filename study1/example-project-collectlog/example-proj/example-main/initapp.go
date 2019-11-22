package main

import "fmt"
import "github.com/astaxie/beego/logs"

type Appconf struct{

	log_level string
	log_path string
	etcd_addr string



}

var appconf Appconf


type Collectconf struct{

	Topic string
	Logfile string
	Exitchan chan string
}

var collects []*Collectconf


type Message struct{

	topic string
	msgtxt string

}

var msgchan  chan Message = make(chan Message)


func server(){

	for {

		msg := <-msgchan

		logs.Debug(fmt.Sprintf("-----%s",msg))
		send_to_kafka(msg.topic,msg.msgtxt)



	}



}




func producer(){

	for _,c := range collects{

		//fmt.Printf("%s",c)
		go tailfile(c)


	}


}
