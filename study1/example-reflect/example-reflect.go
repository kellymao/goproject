package main

import (
	"fmt"
	"reflect"
)



func testint( b interface{}){

	val :=reflect.ValueOf(b)

	c:= val.Int()
	fmt.Println(c)

}


type Student struct{}

func teststruct( b interface{}){

val :=reflect.ValueOf(b)

iv:=val.Interface()

v,ok := iv.(Student)

fmt.Println(v,ok)

}


func setInt( a interface{}){

	val := reflect.ValueOf(a)

	// val.SetInt(100)  panic: reflect: reflect.Value.SetInt using unaddressable value
	//val := reflect.ValueOf(&a) 传地址进去，才可以修改值
	//val.Elem().SetInt(100)  通过Elem() 获取指针的值

	c:=val.Int()
	fmt.Println(c)


}


func modify(){

	var f float32 = 3.2
	val := reflect.ValueOf(&f)
	val.Elem().SetFloat(6.8)

	c:=val.Float()
	fmt.Println(c)



}


type User struct {
	Id int
	Name string
	Age int
}

func test_user(){

	u := User{Id: 1, Name:"andes", Age: 20}
	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)
	//值类型是不可修改的
	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet()) //false false
	//指针类型是可修改的
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet()) //false     true
	fmt.Printf("%v\n", vb)    //&{1 andes 20}
	name :="shine"
	vc := reflect.ValueOf(name)
	//通过 Set 函数修改变量的值
	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Printf("%v\n", vb)   //&{1 shine 20}


}

type Stu struct{


	Name string `json:"stu_name"`
	Age int
}

func (s *Stu) Printstu() {

	fmt.Println("xxxxx",s)
}


func (s *Stu) Printage(a int, b int ) int {

	fmt.Println("xxxxx ",s)
	fmt.Println("result is ", a+b)

	return a+b
}

func test_reflect(){




	var stu *Stu = &Stu{"zhangsan",10,}
	stu.Printstu()

	t_ptr := reflect.TypeOf(stu)

	t:=t_ptr.Elem()

	fmt.Println(t)
	fmt.Printf("t Name() is %v",t.Name())
	fmt.Println()
	fmt.Printf("t Kind() is %v\n",t.Kind())

	fmt.Printf("t String() is %v",t.String())
	fmt.Println()
	fmt.Printf("t NumField() is %v",t.NumField())
	fmt.Println()
	fmt.Printf("t_ptr NumMethod() is %v \n",t_ptr.NumMethod())


	fmt.Printf("t_ptr Method(0) is %T %v \n ",t_ptr.Method(0),t_ptr.Method(0))

	fmt.Printf("t_ptr Method(0).Name is %v \n ",t_ptr.Method(0).Name)  //Printstu

	fmt.Println()
	fmt.Printf("t Field(1) is %T %v",t.Field(0),t.Field(0))
	fmt.Println()
	fmt.Printf("t Field(1).Name is %v",t.Field(0).Name)
	fmt.Println()
	fmt.Printf("t Field(1).Tag is %s",t.Field(0).Tag)
	fmt.Println()


	fmt.Printf("t Field(1).Tag.Get('json') is %s \n ",t.Field(0).Tag.Get("json"))

	s,ok:=t.FieldByName("Name")
	if ok {

		fmt.Println(s.Name )
		fmt.Println(s.Tag )
	}


	v_ptr := reflect.ValueOf(stu)

	v:=v_ptr.Elem()

	fmt.Println(v)

	fmt.Println()
	fmt.Printf("v Type() is %T %v \n",v.Type(),v.Type())

	fmt.Printf("v Type().Field(0) is %T %v %v\n",v.Type().Field(0),v.Type().Field(0),v.Type().Field(0).Name)

	fmt.Printf("v Kind() is %v",v.Kind())
	fmt.Println()
	fmt.Printf("v NumField() is %v",v.NumField())
	fmt.Println()
	fmt.Printf("v_ptr NumMethod() is %v \n ",v_ptr.NumMethod())

	fmt.Printf("v_ptr Method(0) is %T %v \n",v_ptr.Method(0),v_ptr.Method(0))



	fmt.Printf("v_ptr Method(0).Type() is %T %v \n",v_ptr.Method(0).Type(),v_ptr.Method(0))

	fmt.Printf("v_ptr Method(0).Type().Name() String() is %v %v \n",v_ptr.Method(0).Type().Name(),v_ptr.Method(0).Type().String())



	var paramList0  []reflect.Value   // 没有值得情况

	retList0 :=v_ptr.Method(1).Call(paramList0)
	fmt.Println(retList0)

	paramList1 := []reflect.Value{reflect.ValueOf(10),reflect.ValueOf(20)}  // 传值的情况

	retList1  := v_ptr.Method(0).Call(paramList1)

	fmt.Println(retList1)

	fmt.Println(retList1[0].Int())  // 获取第一个返回值, 取整数值

	fmt.Println()
	fmt.Printf("v Field(1) is %T %v",v.Field(0),v.Field(0))
	fmt.Println()
	fmt.Printf("v Field(1).CanSet is %v",v.Field(0).CanSet())
	fmt.Println()
	fmt.Printf("v Field(1).Interface is %s",v.Field(0).Interface())
	fmt.Println()

	fmt.Printf("v FieldByName('Name').Interface is %T %s \n",v.FieldByName("Name"),v.FieldByName("Name").Interface())





}

func main(){


	setInt(100)

	test_user()

	test_reflect()
}