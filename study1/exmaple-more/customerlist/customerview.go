package main

import (
	"fmt"
	"os"
	"study1/exmaple-more/customerlist/customerservice"
	"study1/exmaple-more/customerlist/customer"
)

type Customerview struct{

	customerservice_obj  *customerservice.Customerservice


}



func (c Customerview) listmenu(){

	fmt.Printf("%s\n",`


	欢迎来到客户管理系统：

		1: 打印客户详细信息

		2. 新加客户

		3. 删除客户

		5. 退出
`)

	var intvar int

	for {

	fmt.Println("请输入序号 ")
	_, err := fmt.Fscanf(os.Stdin, "%d\n" ,&intvar)

	if err!=nil{

		fmt.Println("输入错误！",err)
		os.Exit(1)
	}

	fmt.Println("你选择的是 ",intvar)

	switch intvar {

	case 1:
		c.listcustomer()

	case 2:
		c.add()
	case 3:
		c.del()
	case 5:
		c.exit()
	default:
		fmt.Println("没有匹配上")

	}

	}






}


func (c Customerview) listcustomer(){


	c.customerservice_obj.List()


}

func (c Customerview) add(){


	var name string
	var age int
	var gender string
	var email string
	var phone string


	fmt.Println("请输入名字 ")
	fmt.Fscanf(os.Stdin,"%s\n", &name)
	fmt.Println("请输入年龄  ")
	fmt.Fscanf(os.Stdin,"%d\n", &age)
	fmt.Println("请输入性别  ")
	fmt.Fscanf(os.Stdin,"%s\n", &gender)
	fmt.Println("请输入邮箱 ")
	fmt.Fscanf(os.Stdin,"%s\n", &email)
	fmt.Println("请输入电话  ")
	fmt.Fscanf(os.Stdin,"%s\n", &phone)


	c1 := customer.NewCustomer(name,gender,age,phone,email)

	fmt.Printf("要添加的数据： %+v \n",c1)

	if ok := c.customerservice_obj.Add(c1); ok {

		fmt.Println("添加成功！")
		return
	}

	fmt.Println("添加失败！")


}


func (c Customerview) del(){

	var id int

	fmt.Println("请输入删除的客户id  ")
	fmt.Fscanf(os.Stdin,"%d\n", &id)
	if ok := c.customerservice_obj.Delete(id); ok {

		fmt.Println("删除成功！")
		return
	}

	fmt.Println("删除失败！")


}


func (c Customerview) exit(){

	fmt.Println("欢迎下次光临！")
	os.Exit(0)

}


// 客户管理系统

func main(){

	viewobj := Customerview{

		&customerservice.Customerservice{

			make([]*customer.Customer,0),
			0,
		},

	}

	viewobj.listmenu()

}