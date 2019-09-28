package main

import "net"
import "fmt"
import "os"

import "io/ioutil"
import "io"
import "bufio"

import "net/http"
import "time"

import "text/template"

// 打印ip
func print_ip(){

	if len(os.Args) !=2{

		fmt.Println("参数输入错误")
		return
	}
	addr :=net.ParseIP(os.Args[1])

	if addr ==nil{

		fmt.Println("ip转换错误。")
		return
	}

	fmt.Println("ip:",addr.String())



}

//DialTCP函数  tcp client 访问 qbox.me:80
//实现初步的 HTTP 协议，通过向网络主机发送 HTTP Head 请求，读取网络主机返回的信息


func tcp_client(){

	if len(os.Args)!=2{

		fmt.Printf("USAGE: %s host:port \n",os.Args[0])
		os.Exit(1)
	}
	//www.baidu.com:80
	name:= os.Args[1]

	//(*TCPAddr, os.Error)
	tcpaddr,err:=net.ResolveTCPAddr("tcp4",name)

	if err!=nil{

		fmt.Println("tcpaddr resolve err")
		return
	}


	// (c *TCPConn, err os.Error)
	tcpconn,err:= net.DialTCP("tcp4", nil, tcpaddr)



	if err!=nil{

		fmt.Println("tcpaddr resolve err")
		return
	}

	defer tcpconn.Close()

	_, err = tcpconn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

	if err!=nil{

		fmt.Println("tcpaddr resolve err")
		return
	}


	result, err := ioutil.ReadAll(tcpconn)

	if err!=nil{

		fmt.Println("tcpaddr resolve err")
		return
	}

	fmt.Println(string(result))

	os.Exit(0)

	/*

	root@pts/1 # ./example-socket qbox.me:80
	HTTP/1.1 200 OK
	Server: nginx
	Date: Wed, 18 Sep 2019 08:36:46 GMT
	Content-Type: text/html
	Content-Length: 612
	Last-Modified: Wed, 24 Jul 2019 02:28:35 GMT
	Connection: close
	ETag: "5d37c253-264"
	Accept-Ranges: bytes

	 */
}

// ping

func ping(){

	var (
		//icmp   ICMP
		laddr  = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		raddr, _ = net.ResolveIPAddr("ip", os.Args[1])
	)
	fmt.Println(raddr)
	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)
	fmt.Println(conn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(conn)
	defer conn.Close()


}


// Dial 函数
func dial(){


	client,err :=net.Dial("tcp4","0.0.0.0:16000")  // example-channel1 的socket服务器

	if err!=nil {

		fmt.Println(err)
		return
	}

	defer client.Close()

	fmt.Println(client.LocalAddr())   // 127.0.0.1:9630 客户端地址
	fmt.Println(client.RemoteAddr())  // 127.0.0.1:16000 服务器地址



	// 写数据到服务器


	//方式一： io 包的方式
	io.WriteString(client,"hello i come in .\n")
	io.WriteString(client,"hello i come in .\n")


	//方式二： fmt格式化
	fmt.Fprintf(client,"hello i come in .\n")
	fmt.Fprintf(client,"hello i come in .\n")
	fmt.Fprintln(client,"hello i come in .")  // 自动换行




	//方式三：conn自带的方法 Write

	client.Write([]byte("hello i client come in .\n"))

	// 从终端获取，写数据到服务器

	// 方式一： io.Copy

	//io.Copy(client,os.Stdin)  // 写一行文字后，敲回车，服务端就会显示这一行 ，然后一直卡在前台


	// 方式二： fmt 从终端读取数据，发到服务器

	var line string
	fmt.Scanln(&line)     // 从终端读取一行文件，中间不可以有空格

	if err!=nil{

		fmt.Println(err)
	}

	fmt.Fprintln(client,line)


	// 方式三： 缓存读

	input:= bufio.NewReader(os.Stdin)


    /*
	for {

		// 缓存一行读
		data ,err:= input.ReadString('\n')

		if err!=nil{
			fmt.Println(err)
		}
		client.Write([]byte(data))


	}

     */

    go func(){

		data:=make([]byte,5)

		for {
		n,err:=client.Read(data)

		if err!=nil{


			fmt.Println(err)
		}


		fmt.Print(string(data[0:n]))

		}


	}()

	// 缓存字节读, 结果和上面一样，输入回车后才往服务端发送

	buf := make([]byte,5)



	for{
		n,err:= input.Read(buf)

		if err!=nil{

			break

		}

		fmt.Println(n)
		client.Write(buf[0:n])
	}


	// 接受服务器发来的数据


	// 方式一： io.Copy 输出到客户端终端

	io.Copy(os.Stdout,client)  // 会等待接收数据，卡在终端上


	// 方式二：conn.Read, fmt


		data:=make([]byte,5)

		for {                       //通过for循环实现等待接收数据
		n,err:=client.Read(data)

		if err!=nil{


			fmt.Println(err)
		}


		fmt.Print(string(data[0:n]))

		}




	/*
	f:=os.OpenFile("w.log",O_CREATE|O_APPEND,0644)

	out:=[]*os.File{os.Stdout,f}
	for i:= range out {

		writer := bufio.NewWriter(i)

		writer.Write([]byte("hello i come in .\n"))

	}
	writer := bufio.NewWriter()
	*/

}

func dial_baidu(){


	conn ,err:= net.Dial("tcp","www.baidu.com:80")

	if err!=nil{

		return
	}
	defer conn.Close()

	msg:="GET / HTTP/1.1\r\n"
	msg +="Host: www.baidu.com\r\n"
	msg +="Connection: close\r\n"
	msg +="\r\n\r\n"

	conn.Write([]byte(msg))

	io.Copy(os.Stdout,conn)

}

// Listen 函数

func listen(){

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

		go sub_telnet(conn)

	}

}

func sub_telnet(conn net.Conn){


}



// ListenTCP 函数 tcp server


// http client

func http_client(){

	resp,err:= http.Get("www.baidu.com")

	if err!=nil{

		return
	}


	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}


func http_head(){


	var url=[]string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.taobao.com",

	}

	for _,i :=  range url{

		resp,err:=http.Head(i)
		if err!=nil{
			fmt.Printf("head %s failed ,err: %s \n", i ,err)
			continue
		}

		fmt.Printf("head %s success,status:%s \n",i,resp.Status)

	}


/*

   root@pts/4 # ./example-socket
   head http://www.baidu.com success,status:%!d(string=200 OK)
   head http://www.google.com failed ,err: Head http://www.google.com: dial tcp 31.13.69.86:80: connect: connection timed out
   head http://www.taobao.com success,status:%!d(string=200 OK)
*/

}

func http_client1(){

	var url=[]string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.taobao.com",

	}

	for _,i :=  range url{
		c:=http.Client{

			Transport: &http.Transport{

				Dial: func(network,addr string)(net.Conn,error){

						timeout := time.Second*5
						return net.DialTimeout(network,addr,timeout)
				},

			},


		}

		resp,err:=c.Head(i)
		if err!=nil{
			fmt.Printf("head %s failed ,err: %s \n", i ,err)
			continue
		}

		fmt.Printf("head %s success,status:%s \n",i,resp.Status)

	}



}

// http server


// template



func http_template(){


	type Person struct {

		Name string
		Age  int
	}


	t,err:=template.ParseFiles("/tmp/index2.html")

	if err!=nil{

		return
	}


	p :=Person{

		"zhagnsan",
		10,
	}


	/*
	if err:=t.Execute(os.Stdout,p); err!=nil{

		fmt.Println(err)
	}

	p1:=make(map[string]interface{})

	p1["Name"] = "zhagnsan"
	p1["Age"] = 10

	if err:=t.Execute(os.Stdout,p1); err!=nil{

		fmt.Println(err)
	}

	*/


	p2 := Person{
		"liwu",
		 11,
	}


	var plist []Person
	plist = append(plist,p)
	plist = append(plist,p2)

	if err:=t.Execute(os.Stdout,plist); err!=nil{

		fmt.Println(err)
	}



}


func main(){

	//print_ip()

	//tcp_client()

	//ping()

	//dial()

	//dial_baidu()

	//http_client1()

	http_template()
}