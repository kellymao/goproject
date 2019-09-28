package main

import "net"
import "fmt"

func initServer(addr string)error{

	l,err:=net.Listen("tcp",addr)

	if err!=nil{

		fmt.Println("listen failed",err)
		return err
	}


	for {

		conn,err:=l.Accept()


		if err!=nil{

			fmt.Println("accept failed, ",err)
			continue
		}

		go process(conn)



	}




}


func process(conn net.Conn){


	defer conn.Close()

	client := &Client{
		conn:conn,
	}

	err








}
