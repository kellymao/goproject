package customer

import "fmt"

type Customer struct {
	Id   int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}


func NewCustomer(name string,gender string,age int ,phone string,email string ) *Customer{

	return &Customer{
		Name:name,
		Gender: gender,
		Age: age,
		Phone:phone,
		Email:email,
	}

}

func (c *Customer) printinfo (){

	fmt.Printf("%+v \n", *c)
}
