package main

import "time"

func main(){

	init_redis(10,1024,time.Second*10,"0.0.0.0:6379")

	init_usermgr()

	run_server()




}
