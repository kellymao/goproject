
package cls2
import "study1/example-total/base"
import "fmt"

type Cls2 struct{}

func (c2 *Cls2) Know(){

	fmt.Println("cls2 is knowing")
}

func Newcls2()*Cls2{

	return &Cls2{}

}

func f2()base.Class{

	return Newcls2()

}

func init(){

	base.Registermap("cls2",f2)

}
