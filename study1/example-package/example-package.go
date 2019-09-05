package main


// g工厂模式自动注册

import "fmt"
import "study1/example-package/base"
import _ "study1/example-package/cls1"
import _ "study1/example-package/cls2"


func main(){


	v1 := factory.Create("cls1")
	fmt.Println(v1)
	v1.Do()


		v2 := factory.Create("cls2")
		fmt.Println(v2)
		v2.Do()



}

