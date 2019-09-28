package main
import "net"
import "study1/example-chat/common"
import "fmt"


func run_client(){



	conn,err :=net.Dial("tcp4","0.0.0.0:19000")

	common.Check_err(err)

	defer conn.Close()

	fmt.Println(conn.LocalAddr())   // 127.0.0.1:9630 客户端地址
	fmt.Println(conn.RemoteAddr())  // 127.0.0.1:16000 服务器地址



	handleClient(conn)






}



func handleClient(conn net.Conn){


	defer conn.Close()


	var handle *Handle = &Handle{
		conn,
		1,
	}


	handle.run()

	/*
	// 发送登录登录信息
	for i:=0;i<3;i++{

		prompt()

		fmt.Println("登录信息：", string(logininfo()))
		send_data(conn, logininfo())

		status, resptype, respbody := process_login_resp(rev_data(conn))

		if !status {

			fmt.Println("登录失败！,请重新登录")

		} else {

			fmt.Println("登录成功！")

			fmt.Println(resptype)

			fmt.Println(respbody)
			return
		}
	}

	*/

}