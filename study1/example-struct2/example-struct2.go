package main

import "fmt"

import "encoding/json"

// 结构体的初始化

func initst(){

	type People struct{
		name string
		child *People
	}




	var person *People = &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},

	}

	fmt.Println(person)

}
// 匿名结构体的初始化

func Printtype(p struct{
	id int
	name string
	sex rune
}){

	fmt.Printf("打印的类型%T\n",p)
	fmt.Println("打印值",p)

}
func initnoname(){

	var p struct{
		id int
		name string
	}

	p = struct{
		id int
		name string
	}{
		id:10,
		name:"hello world",
	}

	fmt.Println(p)

	pp := struct{
		id int
		name string
		sex rune
	}{
		id:1,
		name:"mao",
		sex:'男',

	}

	fmt.Printf("%+v\n",pp)

	Printtype(pp)

}

// 模拟构造函数
// 模拟构造函数重载 ,两个不同的new函数

type Cat struct{
	Name string
	Color string
}

func Newcat1(name string) *Cat{

	cat := &Cat{}
	cat.Name = name
	return cat

}

func Newcat2(color string) *Cat{

	return &Cat{
		Color:color,
	}
}

func chongzai(){

	fmt.Println(Newcat1("heimao"))
	fmt.Println(Newcat2("red"))
}

// 模拟构造函数 继承。通过结构体包涵

type Car struct{
	name string
	color string
}
type BlankCar struct{
	Car
	madein string
}


func Newcar(name string,color string) *Car{
	// 基类的构造函数

	return &Car{
		name:name,
		color:color,
	}

}

func Newblankcar(name string,color string,madein string) *BlankCar{
	// 子类的构造函数

	bc:=&BlankCar{}
	bc.name=name
	bc.color=color
	bc.madein=madein

	return bc

}

func jicheng(){

	bc := Newblankcar("beinchi","blank","china")
	fmt.Println(bc)
}


// 结构体方法

type Bag struct{

	items []string
}


func (p *Bag) install(item string){

	p.items = append(p.items,item)


}

func to_bag(){

	a := &Bag{}

	a.install("apple")

	fmt.Println(a)

}


// 方法和函数的统一调用，调用者不关新掉的是函数还是方法
/*
本节的例子将让一个结构体的方法（class.Do）的参数和一个普通函数（funcDo）的参数完全一致，也就是方法与函数的签名一致。
然后使用与它们签名一致的函数变量（delegate）分别赋值方法与函数，接着调用它们，观察实际效果。


 */
type Class struct{

}

func (p *Class) do( i int){

	fmt.Println("struct class do is being called")
}


func funcdoing(i int){
	fmt.Println("func do is beging called")
}

func tongyidong(){



	class := &Class{}
	var alldoing func(int) =  class.do
	alldoing(1)

	alldoing = funcdoing

	alldoing(2)

}

// 事件系统的基本原理


var Eventbyname map[string][]func(v interface{})  = make(map[string][]func(v interface{}))


func RegisterEvent(event_name string, method func(interface{}) ) {


	if _,ok := Eventbyname[event_name]; ok {

		v := Eventbyname[event_name]
		v = append(v,method)

		Eventbyname[event_name] = v
	}else{

		Eventbyname[event_name] = []func(v interface{}){method}

	}
}


func CallEvent(event_name string,value interface{}){

	func_list := Eventbyname[event_name]
	for _, f :=range func_list {

		f(value)


	}


}


func  global_event(v interface{}){

	fmt.Println("this student is guojixuexiao shang ke ")
}

type Actor struct{
	name string
	sex rune

}

func (p *Actor) study(yuyan interface{}){

	fmt.Println(p.name,"is studying ", yuyan )
}

func (p *Actor) sing(gequming  interface{}){
	fmt.Println(p.name , "is singing ", gequming)
}




func example_event(){



	stu1 := &Actor{name:"zhangsan",}
	stu2 := &Actor{name:"wangwu",}


	RegisterEvent("study" , global_event)
	RegisterEvent("study", stu1.study)



	RegisterEvent("sing" , global_event)
	RegisterEvent("sing", stu2.sing)



	CallEvent("study","english")

	CallEvent("sing",123567)


}


// 结构体内嵌 和 匿名字段内嵌


type innerS struct{

	A ,B int

}

type outerS struct{

	innerS  // 匿名内嵌结构体 innerS
	i int
	j int
	int     // 匿名字段int

}



type outerS1 struct{
	I int
	J int
	C int     // 匿名字段int
	*innerS  // 匿名内嵌结构体 innerS

}
// 初始化内嵌结构体


func neiqian_struct(){

	var t *outerS = &outerS{

		i:1,
		j:2,
		int:3 ,
		innerS: innerS{
			A: 5,
			B:6,
		},

	}

	fmt.Println(t)

	var t1 *outerS = &outerS{}
	t1.i = 1
	t1.j = 2
	t1.int = 3
	t1.A = 5
	t1.B = 8
	fmt.Println(t1)

	var t2 *outerS1 = &outerS1{}
	t2.I = 1
	t2.J = 2
	t2.C = 3
	t2.innerS = &innerS{
		A:5,
		B:6,
	}
	fmt.Println(t2)

}

// 内嵌结构体会继承 内部结构体的方法，外部结构体可以直接调用


// 结构体序列化和反序列化

func xuliehua(){


	var t2 	 *outerS1 = &outerS1{


		innerS: &innerS{
			A: 5,
			B:6,
		},
		I:1,
		J:2,
		C:3 ,

	}


	fmt.Println(t2)
	rel,err:=json.Marshal(t2)
	if err != nil {

		fmt.Println(err)

	} else{
		fmt.Println(string(rel))
	}
}


// 定义手机屏幕
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}
// 定义电池
type Battery struct {
	Capacity int // 容量
}
// 生成json数据
func genJsonData() []byte {
	// 完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		// 电池参数
		Battery: Battery{
			2910,
		},
		// 是否有指纹识别
		HasTouchID: true,
	}

	fmt.Println(raw)
	// 将数据序列化为json
	jsonData, _ := json.Marshal(raw)
	return jsonData
}


func main(){

	initst()
	initnoname()
	chongzai()

	jicheng()
	to_bag()

	tongyidong()
	example_event()

	neiqian_struct()

	b:=1
	var a  *int = new(int)
	a = &b
	fmt.Println(a)

	xuliehua()

	jsonData := genJsonData()
	fmt.Println(string(jsonData))


}