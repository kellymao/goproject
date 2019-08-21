package main

import "fmt"

import  "math"
import "strings"
import "strconv"

import "time"
import "math/rand"


func modify(v int){

	v=10
	return
}

func modifyptr(v *int){

	*v = 10
	return


}
func modify_ptr(i *int){
	*i = 100
	fmt.Println(i)
}

func main(){

	a:=5
	var b chan int = make(chan int,1)
	var c float32 = 1.15

	fmt.Println("a=",a)

	// 管道是引用类型，输出的是地址
	fmt.Println("b=",b) // b= 0xc0000180e0  输出的地址



	modify(a)
	fmt.Println("a=",a)

	//修改变量的值要传引用
	modifyptr(&a)
	fmt.Println("a=",a)

	fmt.Println(c)

	var d int
	var e int32
	d = 3
	e = int32(d)

	// Printf 进行格式化输出
	fmt.Printf("e=%d\n", e)

	/*
	for i:=0; i<10; i++{

		a:= rand.Int()
		fmt.Println(a)
	}

	for i:=0;i <10 ; i++{
		a:= rand.Intn(100)
		fmt.Println(a)

		b:= rand.Float32()
		fmt.Println(b)
	}

	*/
	var f byte = 'u'
	fmt.Println(f)
	fmt.Printf("%s\n",f)
	fmt.Printf("%c\n",f)


	var g int
	var h bool
	fmt.Printf("%v\n",f)
	fmt.Printf("g=%v\n",g)
	fmt.Printf("h=%v\n",h)

	fmt.Printf("%p\n",&h)
	fmt.Printf("%T\n",h)

	// Sprintf 数字转字符串以及 赋值给变量

	var c1 = fmt.Sprintf("%d-why",123)
	fmt.Println(c1)
	// strings 包的方法


	// 自定义字符串反转

	var c2 string="hello world!"
	var result string
	for i:=0; i<len(c2);i++{
		//result =result + string(c2[len(c2)-i-1])
		result = result + fmt.Sprintf("%c",c2[len(c2)-i-1])
	}
	fmt.Println(result)


	// 字节数组和字符串转换
	var result2 []byte
	for i:=0;i<len(result);i++{
		tmp:=[]byte(result)
		result2 = append(result2,tmp[len(result)-i-1])

	}
	fmt.Println(string(result2))

	// 输出100-200 之间的素数
	var flag bool
	for i:=100; i<=200;i++{
		flag = true
		for j:=2;j<i;j++{
			if i % j == 0{
				flag = false
				break
			}

		}
		if flag{
			fmt.Println(i,"是一个素数")
		}
	}

	// 命令行接受参数
	var c3,c5 int
	//fmt.Scanf("%d %d", &c3,&c5)

	fmt.Println(c3)
	fmt.Println(c5)

	//通过指针传递修改函数外部变量的值

	var n int

	modify_ptr(&n)
	fmt.Println(n)
	//求平方根math.Sqrt

	var flag2 bool
	for i:=100; i<=200;i++{
		flag2 = true
		for j:=2;j<int(math.Sqrt(float64(i)));j++{
			if i % j == 0{
				flag2 = false
				break
			}

		}
		if flag2{
			fmt.Println(i,"是一个素数")
		}
	}


	//strings 包

	c8:=strings.Fields("abc dfe jlj")
	fmt.Println(c8)

	//strconv
	number,err :=strconv.Atoi("123")

	if err !=nil {
		fmt.Println("convert is failed!")
	}

	c10:=strconv.Itoa(123)
	fmt.Println(number,c10)

	// 统计耗时间

	t1:= time.Now().UnixNano()

	time.Sleep(time.Millisecond * 1000 )
	t2:= time.Now().UnixNano()

	fmt.Println("t2 - t1 毫秒 is ", (t2-t1)/1000/1000)


	// switch 查找随机数: 不能直接加break ，没有这种语法

	rand_number := rand.Intn(100)
	var w int


	for{

	stop_flag := false

	fmt.Println("请输入一个1-100之间的整数： ")
	fmt.Scanf("%d\n", &w)

	switch {

	case w == rand_number:
		fmt.Println("you are right.")
		stop_flag=true

	case w <	rand_number:
		fmt.Println("smaller")

	case w > rand_number:
		fmt.Println("bigger")


	}

	if stop_flag{
		break
	}
	}

	// for range 字符串


	



}
