package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main(){

	wg.Add(1)
	conn,ok :=client()
	if !ok{

		fmt.Println("聊天室无法连接")
		return
	}

	fmt.Println("服务器登录成功")
	go handle_stdin()
	go handleconn(conn)
	go recvmsg(conn)


	// 开始验证登录
	status:=login()
	if !status{

		fmt.Println("登录失败请重新登录！")
		wg.Done()
	}


	logicview()


	wg.Wait()



}
