package main

import "fmt"
import "math/rand"

type Student struct{

	Name string
	Age int
	Score float32

}

type Student_single_list struct{

	Name string
	Age int
	Score float32
	Next *Student_single_list

}

func insertbefore(head *Student_single_list){


	for j:=0 ; j<2;j++{

		/*
		变量不可以重复定义， 但是for 循环里面可以重复定义，应该是每次for循环里面的同名变量都是一个局部变量 。
		 */
		var a *int = new(int)
		*a = 3
		fmt.Println(a)
	}


	for i:=0 ;i<100;i++{
		var p *Student_single_list = new(Student_single_list)
		p.Name = fmt.Sprintf("stu_%d",i)
		p.Age = rand.Intn(100)
		p.Score = float32(rand.Float64() * 100)
		head.Next = p
		head = p

	}



}
func lianbiao(){

	var head Student_single_list
	head.Age = 18
	head.Score = 80
	head.Name = "zhuzhu"

	var stu1 Student_single_list = Student_single_list{"zhagnsan",90,100,nil}
	var stu2 *Student_single_list = &Student_single_list{"wangwu",90,100,nil}


	head.Next = &stu1
	stu1.Next = stu2


	var p *Student_single_list = &head

	for p != nil{
		fmt.Println(p.Name)

		if p.Next == nil {

			break
		}
		p = p.Next
	}

}

func main(){

	var stu Student
	stu.Age = 18
	stu.Score = 80
	stu.Name = "zhuzhu"

	fmt.Println(stu)
	fmt.Printf("name:%p\n",&stu.Name)
	fmt.Printf("age:%p\n",&stu.Age)
	fmt.Printf("score:%p\n",&stu.Score)

	var stu1 Student = Student{"zhagnsan",90,100}

	var stu2 *Student = &Student{}
	stu2.Name = "wang"
	fmt.Println(stu1,stu2)

	lianbiao()


	var head *Student_single_list = new(Student_single_list)
	head.Age = 18
	head.Score = 80
	head.Name = "zhuzhu"
	insertbefore(head)

	for head != nil {
		fmt.Println(head)

		if head.Next == nil{
			break
		}

		head = head.Next

	}


}