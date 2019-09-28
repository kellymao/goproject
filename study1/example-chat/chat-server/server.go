package main


import "fmt"
import "net"
import "study1/example-chat/common"

func run_server(){



	server,err:=net.Listen("tcp","0.0.0.0:19000")

	common.Check_err(err)

	defer server.Close()

	for {

		conn,err:=server.Accept()
		if err!=nil {

			fmt.Println("server accept error")
			continue
		}

		go handleClient(conn)

	}






}


func handleClient(conn net.Conn){


	var handle Handle = Handle{

		conn: conn,

	}

	handle.process()






}

/*
func rev_msg(conn net.Conn) []byte {

	var bufdata []byte

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err!=nil {

			fmt.Println("recive  data error")
			return nil
		}


		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))

		bufdata = append(bufdata, buf[0:n]...)

		break

	}


	return bufdata



}


func send_msg(conn net.Conn, data []byte){


	_,err:=conn.Write(data)
	common.Check_err(err)

}



	_, err = conn.Write([]byte("Welcome client\n"))
	if err!=nil {

		fmt.Println("send  data error")
		return
	}

	 */


