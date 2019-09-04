package cls1

import "study1/example-package/base"
import "fmt"

type mycls1 int

func (p *mycls1) Do(){

	fmt.Println(*p,"mycls1 is doing")
}


func c1() factory.Class{

	fmt.Println("cls1 is calling")
	var v1 *mycls1 = new(mycls1)

	a:=1
	*v1 = mycls1(a)
	return v1



}



func init(){

	factory.Register_cls("cls1", c1)


}


