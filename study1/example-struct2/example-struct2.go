package main

import "fmt"

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

func main(){

	initst()
	initnoname()

}