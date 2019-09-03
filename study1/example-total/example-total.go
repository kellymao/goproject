package main

import "fmt"

import "encoding/json"
import "study1/example-total/base"
import _ "study1/example-total/cls2"
import _ "study1/example-total/cls1"

type Wheel struct{

	Size int
}

type Engine struct{

	Power int
	Type string
}


type Car struct{
	Wheel
	Engine
}


func NewCar(Size int,Power int,Type string) *Car{


	return &Car{
		Wheel:Wheel{
			Size:Size,

		},

		Engine:Engine{
			Power:Power,
			Type:Type,
		},


	}


}



type Flying struct{

}

func (f *Flying) Fly(){

	fmt.Println(*f,"is flying")
}


type Walking struct{}


func (w *Walking) Walk(){

	fmt.Println(*w,"is walking")

}

type Bird struct{
	Flying
	Walking

}


func NewBird() *Bird{

	return &Bird{}
}


func t_main(){

	var bird *Bird = NewBird()
	bird.Fly()
	bird.Walk()

}



func j_main(){

	var car *Car = NewCar(10,80,"1.4T")

	data,_ := json.Marshal(car)
	fmt.Println(string(data))

	var car1 *Car = new(Car)

	json.Unmarshal(data,car1)
	fmt.Println(car1)
}
func main(){


	var car *Car = NewCar(10,80,"1.4T")
	fmt.Println(car)

	t_main()
	j_main()

	base.Callmap("cls1")
}