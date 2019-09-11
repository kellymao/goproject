package main

import "flag"
import "fmt"
import "encoding/json"


func jsontest(){

	var m map[string]interface{} = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 10
	m["sex"] ="man"

	data,err:= json.Marshal(m)
	if err!=nil{

		fmt.Println(err)
		return
	}

	fmt.Println(string(data))



	var s []map[string]interface{}

	m1:=make(map[string]interface{})
	m1["name"] = "zhagnsan"
	m1["age"] = 10

	s = append(s,m1)
	m2:= make(map[string]interface{})
	m2["name"] = "san"
	m2["age"] = 30

	s = append(s,m2)

	data1,err:= json.Marshal(s)
	if err!=nil{

		fmt.Println(err)
		return
	}

	fmt.Println(string(data1))

	var  s1 []map[string]interface{}
	json.Unmarshal(data1,&s1)

	fmt.Println(s1)


}

func main(){


	var confpath string
	var loglevel int

	flag.StringVar(&confpath,"cnf","","config file name")
	flag.IntVar(&loglevel,"l",0,"log level")


	flag.Parse()

	fmt.Println("path:",confpath)
	fmt.Println("log level:",loglevel)

	jsontest()

}
