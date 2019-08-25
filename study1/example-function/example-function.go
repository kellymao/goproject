package main

import "fmt"

import "bufio"
import "os"
import "strings"

import "errors"

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

	str_rune := []rune(str)

	fmt.Println(len(str_rune))

	for i,v := range str_rune {
		fmt.Println(i)
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <='Z':
			wordcount++
		case v >=  '0' && v <='9':
			numbercount++
		case v ==' ':
			spacecount++
		default:
			charcount++
		}


	}

		return
	}


func initconfig()(err error) {

	return errors.New("init config failed")
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

	//var str1 string="hello world 中古人你好"

	reader :=bufio.NewReader(os.Stdin)

	result, _, err := reader.ReadLine()
	if err != nil{
		fmt.Println("read from console err.",err)
		return
	}

	//result 是[]byte 类型
	str2 := string(result)
	wordcount,spacecount,numbercount,charcount := total_count(str2)

	fmt.Println(wordcount,spacecount,numbercount,charcount)

	ret:=strings.Split(str2,"中")    //你  678是  一个人aaa中国  人123 89
	fmt.Println(ret)

	ret1 := strings.TrimSpace(ret[0])
	ret1  =strings.Replace(ret1," ","", -1) //替换字符串中的空格
	ret2 :=strings.TrimSpace(ret[1])
	fmt.Println(ret1,ret2)



	// 内置函数 new 和 make, append


	i :=new(int)
	fmt.Println(i)

	*i = 90
	fmt.Println(*i)

	k :=new([]int)
	fmt.Println(k)
	*k = []int{1,2,3,5,6}
	fmt.Println(k,*k)


	var j []int
	j = append(j,10,20,30)
	j = append(j,j...)

	fmt.Println(j)

	/*
	0xc000014110
	90
	&[]
	&[1 2 3 5 6] [1 2 3 5 6]
	[10 20 30 10 20 30]
	 */

   // recover() 函数



   // []byte 可以和 string 相互转换 ，[]rune 可以和string 相互转换
}
