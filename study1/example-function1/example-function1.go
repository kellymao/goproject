package main

import (
	"fmt"
	"strings"
	"bytes"
)

import "os"
import "strconv"

// 阶乘 - 递归
func jiecheng(n uint64) uint64 {

	if n ==1 {

		return 1

	} else {

		return n * jiecheng(n-1)
	}
}

// 闭包, 函数和外部环境的整体.返回的是一个函数

func Adder() func(int) int {
	var x int

	return func(n int) int {

		x += n
		return x
	}


}


// 函数参数传递，值传递和引用传递的例子

type School struct{

	address string
	number int

}

type Student struct{

	name string
	school School
	schoolptr *School



}

func transfer(stu Student) Student{

	fmt.Println("在函数（传的值） 中：")
	fmt.Printf("in zhi ptr function %p \n", &stu)
	fmt.Println("in zhi function",stu)

	return stu
}


func transfer_ptr(stu *Student) *Student{

	fmt.Println("在函数（传的引用） 中：")
	fmt.Println("in ptr ptr function", stu)
	fmt.Println("in ptr zhi function",*stu)

	*stu = Student{}
	fmt.Println("in ptr zhi function after modify",*stu)
	return stu
}

//字符串切片链式处理，返回原切片 （ 切片默认就是传引用）

func process_string(str []string , funtool []func(s string)string){

	/* 此处没有返回原切片， 因为切片默认是传引用。 外部值已经被修改 */

	for i, v := range str {


		for _, f := range funtool{

			v = f(v)
		}


		str[i] = v

	}





}

func trim_str(str string) string{

	return strings.TrimPrefix(str,"go")

}

// 函数类型实现接口,类型转换funchander 后才能赋值给接口变量

type invoke interface{

	call(v interface{})
}

type invoker struct{
	name string
	age int }


func (t *invoker) call(v interface{}){
	fmt.Printf("invoker %+v is calling . \n",t)

}

type funchander func(int) string


func (t funchander) call(v interface{}){

	fmt.Printf("funchander %+v is calling\n", t )
}


// 可变参数

func process_arg(arg ...interface{}){

	var arg_list []interface{}

	//arg_list = make([]interface{},10)
	arg_list = append(arg_list,arg...)   // 将arg可变参数切片完整传递给下一个函数

	fmt.Printf("%v \n ",arg_list)

	var b bytes.Buffer

	for _, v := range arg{

		b.WriteString("type is ")
		switch v.(type){
			case int:
				b.WriteString("int")
			case string:
				b.WriteString("string")
			//case struct:
			//	b.WriteString("struct")
			default:

				b.WriteString("other")




	}
		b.WriteString(". value is ")
		b.WriteString(fmt.Sprintf("%v",v))  //格式化字符串
		b.WriteString("\n")


	}

	fmt.Println(b.String())

}


// defer 关闭文件

func process_file(file string) int64{

	f,err:=os.Open(file)
	if err != nil{

		fmt.Println("打开文件错误",err)
		return 0
	}

	defer f.Close()

	info,err:=f.Stat()
	if err!= nil{

		fmt.Println(err)
		return 0
	}

	return info.Size()
}


// 自定义 error


type ParseError struct{

	name string
	str string

}

func (p *ParseError) Error() string {


	return fmt.Sprintf("%s %s",p.name,p.str)


}


func process_parserr(v interface{}) (bool , error)  {


	err := &ParseError{name:"zhangsan",str:"parse error"}

	switch v.(type){

	case string:
		return true , nil

	case invoke:
		return false, err

	default:
		return false, err

	}






}
// panic defer recover

/*
defer +recover 匿名函数自执行
捕获panic 触发的异常
panic 本身是一个函数，参数为interface{}

func  panic(v interface{}){

}

 */


type panictext struct{

	text string
}


type f_call func(string) int

func panic_call(f f_call, str string ){

	defer func(){

		if err:=recover(); err !=nil{

			fmt.Println("捕获到err",err)
		}

	}()


	f(str)
}




func main(){

	 fmt.Println(jiecheng(10))

	 f := Adder()
	 fmt.Println(f(10),f(10))


	 stu := Student{
	 	name:"zhangsan" ,
	 	school:School{address:"beijing", number:10 } ,
	 	schoolptr: &School{address:"shanghai", number:15}}

	 fmt.Println("在把 stu 传给函数之前：")
	 fmt.Printf("out zhi ptr function %p \n", &stu)
	 fmt.Println("out zhi zhi function", stu)

	 rel := transfer(stu)



	fmt.Println("在把 stu 传给函数之后：")
	fmt.Printf("return  zhi ptr function %p \n ", &stu)
	fmt.Println("return  zhi zhi function", stu)

	fmt.Println("函数返回之后，打印return 值的内容：")
	fmt.Printf("return  zhi ptr function %p \n ", &rel)
	fmt.Println("return  zhi zhi function", rel)

	rel_ptr:=transfer_ptr(&stu)

	fmt.Println("在把 stu 传给函数(引用）之后：")
	fmt.Printf("return  ptr ptr function %p %p \n", &stu ,rel_ptr)
	fmt.Println("return  ptr zhi function", stu ,*rel_ptr)

	fmt.Println("结论： 无论传递值，还是返回值，只要传的值都是复制数据。 传引用的话是只是复制的地址，还是指向了原来的变量")


	str_list := []string{
		"go nni",
		"go wwi",
		"go tta" }

	funtool := []func(str string) string {    /*此处理解为[] 类型 匿名函数用来做类型， 如[] int */
		trim_str,
		strings.TrimSpace,
		strings.ToUpper}

	process_string(str_list,funtool)

	fmt.Println(str_list)  /* 切片在函数内已经被修改 */


	t:=&invoker{name:"zhangsan",age:10}
	t.call(nil)

	var ti invoke

	ti = t
	ti.call(1)

	ti= funchander(func(c int)string{return fmt.Sprintf("in func %d",c)})  //类型转换funchander 后才能赋值给接口变量 
	ti.call("abc")
	//fmt.Println(ti(3))



	process_arg(13,"中国","abc",&invoker{name:"zhangsan",age:10})

	fmt.Println("size is(字节) ",process_file("../src/study1/example-leixing/example-leixing.go"))


	for  _, parse_str:= range []interface{}{"中国",invoker{name:"zhangsan",age:10}} {
		if o, err := process_parserr(parse_str); err != nil {

			fmt.Println(err)
		} else {

			fmt.Println(o)
		}
	}


	panic_call(f_call(func(n string)int {
		fmt.Println("too hard")

		panic(panictext{"a new test"})
		if v, err := strconv.Atoi(n); err == nil {
			return v
		} else {
			return 0
		}
	}), "199")
}