package main

import (
	"fmt"
	"strconv"
	"encoding/json"
)



var Name string
var Age int




const(
	MALE string ="男"
	FEMALE string = "女"
)

func init(){

	Name = "zhangsan"
	Age = 10
	fmt.Println(MALE)


	var c1 []byte
	c1 = []byte("hello world!")
	fmt.Println(string(c1))

	var c2 []byte
	c2 = []byte("中国")
	fmt.Println(c2)
	fmt.Println(string(c2))

	var c3 = fmt.Sprintf("str-%d", 123)
	fmt.Println(c3)

	var v5 = strconv.Itoa(12356)
	fmt.Println(v5)

}

func main(){

	var a map[string]string 
	fmt.Println(a==nil) //true
	var b map[string]string = map[string]string{"name":"alex"}
	data,_:=json.Marshal(b)
	json.Unmarshal(data,&a)
	fmt.Println(a)
	fmt.Println(a==nil) //false  
	a = make(map[string]string)
	fmt.Println(a)
	fmt.Println(a==nil) //false

}
