package cls2


import "study1/example-package/base"
import "fmt"



type mycls2 int

func (p *mycls2) Do(){

	fmt.Println(*p, "mycls2 is doing")
}

func c2() factory.Class{

	fmt.Println("cls2 is calling")
	var v2 *mycls2 =  new(mycls2)

	b:=2
	*v2 = mycls2(b)
	return v2
}



func init(){

	factory.Register_cls("cls2", c2)


}

