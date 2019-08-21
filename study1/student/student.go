package student

import (
	"fmt"
	"strconv"
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
