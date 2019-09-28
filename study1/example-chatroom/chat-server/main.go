package main


import "time"


func main(){


	InitRedis("localhost:6379",10,1024,time.Second*100)
	initUsermgr()

	initServer("0.0.0.0:10000")

}