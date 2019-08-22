package main

import "fmt"

// 函数也是一种类型: type 定义的函数类型

type add_func func( a int, b int) int

func add(a int ,b int) int {
	return a+b
}

func operate(op add_func , a int , b int) int {
	return op(a,b)
}


// 函数也是一种类型： 匿名函数做类型

var no_name func( int ,int ) int

func operate_noname(op func( a int, b int) int, a int ,b int) int {
	return op(a,b)
}


// map 默认就是引用传递

type mapstr map[string]int


func op_map(a mapstr){

	fmt.Printf("%v",a)
	fmt.Printf("%v", a)
}

// 可变参数列表

func add1(a int,arg ...int) int {

	var num int = a
	for _,v := range arg{
		num +=v
	}
	return num
}


// len(str1) 是打印的字节长度，不是字符长度 中文长度是3 字节 ,  len([]rune(str1)) 是字符长度


// 终端读取一行内容，输出中文，字母，空格的个数


func total_count(str string)(wordcount,spacecount,numbercount,charcount int){


	for _,v := range str {

		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <='Z':
			wordcount++
		case v >=  0 && v <=9:
			numbercount++
		case v ==' ':
			spacecount++
		default:
			charcount++
		}


	}

		return
	}




func main(){



	c:=add

	fmt.Println(operate(c,1,2))


	var d  add_func

	d = add

	fmt.Println(operate_noname(d,3,5))


	fmt.Println(operate_noname(add,3,7))

	//var a map[string]int
	//op_map(a)

	fmt.Println(add1(1,2,3,5,6))

	var str1 string="hello world 中古人你好"
	wordcount,spacecount,numbercount,charcount := total_count(str1)

	fmt.Println(wordcount,spacecount,numbercount,charcount)


}
