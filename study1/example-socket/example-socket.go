package main

import "net"
import "fmt"
import "os"

import "io/ioutil"

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

// tcp client 访问 qbox.me:80
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

func main(){

	//print_ip()

	//tcp_client()

	ping()

}