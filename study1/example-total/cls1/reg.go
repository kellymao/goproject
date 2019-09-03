
package cls1

import "study1/example-total/base"
import "fmt"

type Cls1 struct{}

func (c1 *Cls1) Know(){

	fmt.Println("cls1 is knowing")
}

func Newcls1()*Cls1{

	return &Cls1{}

}

func f1()base.Class{

	return Newcls1()

}

func init(){

	base.Registermap("cls1",f1)

}


