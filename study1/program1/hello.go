package main

import(

	"fmt"
	"study1/student"
)


func compute(n int){

	for i:=0 ; i <= n ;i++   {
		fmt.Printf("%d + %d = %d \n" ,i,n-i,n)
	}

}

func getdata(a,b int)(int,int){

	return a+b, a-b
}

func main(){
	var a,b int


	a = 10
	b = 30
	fmt.Println(a + b)

	fmt.Println("name is :"  , student.Name)

	compute(5)

	_,subval := getdata(30,20)
	addval,_ := getdata(30,20)
	fmt.Println(addval,subval)
	fmt.Println(student.FEMALE)
}