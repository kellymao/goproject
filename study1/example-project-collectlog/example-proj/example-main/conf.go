package main


import (
	"fmt"
	"github.com/astaxie/beego/config"
)




func Readconf() {
	conf, err := config.NewConfig("ini", "/workspace/src/study1/example-project-collectlog/example-proj/conf/logagent.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		panic(err)
	}

	/*
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		panic(err)
	}

	fmt.Println("Port:", port)

	 */
	log_level := conf.String("logs::log_level")
	if len(log_level) == 0 {
		log_level = "debug"
	}


	appconf.log_level = log_level

	log_path := conf.String("logs::log_path")

	appconf.log_path = log_path

	etcd_addr:= conf.String("etcd::etcd_addr")

	appconf.etcd_addr = etcd_addr
}

