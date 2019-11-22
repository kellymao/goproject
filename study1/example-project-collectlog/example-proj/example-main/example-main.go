package main
import "github.com/astaxie/beego/logs"


func main(){




	Readconf()  // 读取配置文件
	err:=initLogger()  //初始化日志
	if err!=nil{
		panic(err)
	}

	logs.Debug("init logger success!")

	initetcd() // 里面启动了子线程tail  到管道 msg

	//producer()
	server()  // for循环 取管道msg的数据，发送到kafka




}
