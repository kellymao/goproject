package main
import "fmt"

// 接口的声明

type Writer interface{

	writedata(v interface{})(n int ,err error)

}
type Reader interface{

	readdata(v interface{})

}

type File struct{}


func (f *File) writedata(v interface{})(n int ,err error){

	fmt.Println("writedata:", v)
	return

}

func (f *File) readdata(v interface{}){

	fmt.Println("readedata:", v)
	return


	fmt.Println("end:")
}

func Test(a interface{}){

	b,ok := a.(int)    //转成int类型之后才能加3

	if ok {

		fmt.Println("%d %T",b,b)
		b += 3

		fmt.Println(b)

		return

	}

	fmt.Println("convert failed")


}


func e1(){

	var write Writer = &File{}
	fmt.Println(write)

	n,err:=write.writedata("write a few data")

	fmt.Println(n,err)


	var read Reader = &File{}

	read.readdata("read a few data")

}

type Write interface{

	write()
}



type Summerywrite struct {

}

func ( s *Summerywrite) write(){}




func main(){

	var b2 int
	Test(b2)

	var c2 interface{}
	fmt.Println(c2)
	fmt.Printf("type is %T\n", c2)
	var a2 int
	c2 = a2


	fmt.Printf("type is %T\n", c2)


	e1()

	fmt.Println("=============")


	var a interface{}  = Summerywrite{}

	var a1 interface{}  = &Summerywrite{}


	var b interface{}  = "hello"  // b的动态类型为内置类型string


	var c Write = &Summerywrite{}


	rel1,ok1 := a.(Write)
	fmt.Println(rel1,ok1)              // <nil> false


	rel2,ok2 := a1.(Write)
	fmt.Println(rel2,ok2)              //  &{} true

	rel3,ok3 := b.(Write)
	fmt.Println(rel3,ok3)             // <nil> false



	var result interface{}              //先定义类型，在进行断言测试
	var is_boo bool
	result,is_boo = c.(interface{})     //c 实现了interface() 和 Write 两个接口

	fmt.Println(result,is_boo)           //  &{} true

	//rel:= b.(Write)                     // 断言失败出现恐慌
	//fmt.Println(rel)                    // panic: interface conversion: string is not main.Write: missing method write




/*
   <nil> false
   &{} true
   <nil> false
   panic: interface conversion: string is not main.Write: missing method write

 */

	var ii interface{} = 1

	var bi int
	//bi = ii              // 触发宕机了，空接口不能直接转换为int

	bi = ii.(int)       // 通过断言进行类型转换

	fmt.Println(bi)
}