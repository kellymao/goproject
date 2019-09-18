package main

import "time"
import "fmt"
import "math/rand"
import "net"
import "bufio"
import "io"
import "strings"

import "sync"

var wg sync.WaitGroup


func time_test(){


	var chanint chan int

	go func(){

		defer func(){

			fmt.Println(recover())
			if e:=recover();e !=nil {

				fmt.Println(e)
			}

		}()


		var chanint1 chan int
		chanint1 <-0

		time.Sleep(time.Second)
		chanint<- 0

		time.Sleep(time.Second)
		chanint<- 1

	}()


	go func() {
		for {
			fmt.Println(1)
			data, ok := <-chanint
			fmt.Println(2)
			fmt.Println(data)
			fmt.Println(ok)

		}

	}()


	time.Sleep(15*time.Second)
}


func tim_tt(){


	v:=<-time.After(time.Second)
	fmt.Println(v)  // 2019-09-12 17:55:29.969551877 +0800 CST m=+1.000694811

}


// 模拟一个打棒球程序

func main_play(){

	var chanball chan int = make(chan int)

	var exit chan int = make(chan int)



	go sub_play("张三",chanball,exit)

	go sub_play("王五",chanball,exit)


	chanball <- 1
	<-exit


}

func sub_play(name string,ball,exit chan int){


	for {

		data:=<-ball

		fmt.Printf("%s 第%d 个回合\n",name,data)


		randnum:=rand.Intn(100)
		if randnum%13 == 0 {


			fmt.Printf("%s 没接到,输了\n",name)


			exit<-0
			return


		}

		time.Sleep(time.Second)
		fmt.Printf("%s 接%d球success\n",name,data)
		data+=1

		ball<-data


	}



}

// 模拟一个接力跑
func main_runnning(){



	var chanint chan int = make(chan int )

	var begin_time time.Time = time.Now()
	fmt.Println("比赛开始的时间",begin_time)

	wg.Add(1)
	go sub_running(chanint,&begin_time)

	chanint<-1
	wg.Wait()

}


func sub_running(chanint chan int,begin_ptr *time.Time){

	defer wg.Done()
	num:=<-chanint

	fmt.Printf("我是第%d棒\n" ,num )

	time.Sleep(time.Second)

	if num == 4 {

		end_time:=time.Now()

		total_time:=end_time.Sub(*begin_ptr)
		fmt.Println(end_time)
		fmt.Println("比赛结束了，最后的成绩是 ",total_time)
		return
	}

	num+=1

	wg.Add(1)
	go sub_running(chanint,begin_ptr)

	chanint<-num           // 先建立个goroutine 在传值。通道是两个线程之间传数据才可以

}


// channel 超时机制

/*

time.After(time.Second)
time.NewTicker() *Ticker
time.NewTimer
time.Afterfunc(time.Second,func)
 */


func main_timeout(){

	var chanint chan int = make(chan int )

	go sub_timeout(chanint)

	time.Sleep(time.Second*5)
	//chanint<- 0

}

func sub_timeout(chanint chan int){


	select {

		case ch:=<-chanint:

			fmt.Println(ch)
		case ch:=<-time.After(time.Second):
			fmt.Println("timeout",ch)


	}



}

// range 循环 time.NewTicker()


func main_ticker(){

	t:=time.NewTicker(time.Second)

	for newtime:= range t.C {

		fmt.Println(newtime) //每秒钟执行一次

	}


}


// 计时器 time.NewTicker() 和 time.NewTimer

func main_timer(){

	t:=time.NewTicker(time.Second *3)

	t1:=time.NewTimer(time.Second*16)

	num:=1

	for {
		select {

		case <-t.C:
			fmt.Println("调用的次数", num)
			num++

		case <-t1.C:
			fmt.Println("结束")
			return

		}
	}

}


// time.AfterFunc 函数用来做超时退出

func main_afterfunc(){


	var exit chan int = make(chan int)

	time.AfterFunc(time.Second * 3,func(){


		fmt.Println("three Seconds，timeout exit。 ")

		exit<-0


	})

	<-exit

}


// telnet 服务器

func main_telnet(){

	var exit chan int = make(chan int )

	go server_telnet(exit)


	<-exit
	fmt.Println("Server is shuttind down.")


}

func server_telnet(exit chan int){
	server ,err :=net.Listen("tcp","0.0.0.0:16000")
	if err!=nil {

		fmt.Println("server listen error")
		return
	}

	defer server.Close()


	for {

		conn,err:=server.Accept()
		if err!=nil {

			fmt.Println("server accept error")
			continue
		}

		go sub_telnet(conn,exit)

	}

}

func sub_telnet(conn net.Conn, exit chan int){

	defer conn.Close()

	//buf:=make([]byte,512)

	//conn.Read(buf)

	inputreader:=bufio.NewReader(conn)


	for {
		input, err := inputreader.ReadString('\n')

		if err != nil {
			if err == io.EOF {

				fmt.Println("read finished")
			} else {
				fmt.Println(err)
				return
			}
		}

		fmt.Println(input)

		if !process_telnet(input,exit) {



			conn.Close()
			return
		}

		conn.Write([]byte(input + "\n"))
	}


}

func process_telnet(input string,exit chan int) bool {


	if strings.HasPrefix(input,"@close"){

		return false
	}

	if strings.HasPrefix(input,"@shutdown"){

		exit<-1
		return false
	}

	return true

}

// 单向通道的使用

/*

1. 定义了一个双向的通道 chinint

2. 声明一个只能收数据的通道类型 chinrev , 并赋值为chinint

3. 声明一个只能取数据的通道类型 chinsed , 并赋值为chinint

3. 定义一个goroutine  sub_rev_channel 往里面chinrev 传值

4. 从chinsed 和 chinint 两个通道都可以取值

 */

func main_single_channel(){


	var chinint chan int = make(chan int)


	var chinrev chan<- int  = chinint

	go sub_rev_channel(chinrev)



	var chinsed <-chan int  = chinint

	fmt.Println("first from chinsed :", <-chinsed)

	for data:= range chinint{


		fmt.Println(data)
	}








}

func sub_rev_channel(chanint chan<- int){

	fmt.Println(chanint)
	chanint<-3

	for i:=0;i<30;i++{

		chanint<-i

	}


}


/*

自定义一个type，类型为通道。 client 里面存的是通道


 */


func main_type_channel(){

	type client chan<- string

	var client_send chan client = make(chan client)


	ch:=make(chan string)
	//ch<-"job"


	go func(){

		fmt.Println(client_send)
		fmt.Println(<-client_send)

	}()

	client_send <- ch






}

func main (){

	//main_play()

	//main_runnning()

	//main_timeout()

	//main_ticker()

	//main_timer()

	//main_afterfunc()

	//main_telnet()

	//main_single_channel()

	main_type_channel()






}

