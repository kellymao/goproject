package main

import "flag"


func main(){


	var confpath string
	var loglevel int

	flag.StringVar(&confpath,"c","","config file name")
	flag.IntVar(&loglevel,"d",0,"log level")


	flag.Parse()

	fmt.Println("path",confpath)
	fmt.Println("log level",loglevel)



}
