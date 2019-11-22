package customerservice

import (
	"fmt"
	"study1/exmaple-more/customerlist/customer"
)

type Customerservice struct{

	Customers []*customer.Customer
	CustomerCount int


}


func (this *Customerservice) Add(c *customer.Customer) bool {

	this.CustomerCount++

	c.Id = this.CustomerCount
	this.Customers = append(this.Customers,c)


	return true
}

func (this *Customerservice) Delete(customer_id int ) bool {



	//只有一个元素的时候，删除会失败

	if len(this.Customers) <2 {

		fmt.Println("不可以删除，元素太少了")
		return false
	}

	index,ok := this.findbyid(customer_id)
	if !ok {

		return false
	}


	// 两个切片如何相加
	this.Customers = append(this.Customers[0:index],this.Customers[index+1:len(this.Customers)]...)



	return true
}


func  (this *Customerservice)  findbyid( id int) (int , bool) {

		// 根据传入的customer id 找切片的序号

		for i:=0;i<len(this.Customers);i++{

			if this.Customers[i].Id == id {
				return i ,true
			}



	}

		return -1 ,false


}


func (this *Customerservice) List(){



	for i , v := range this.Customers{

		fmt.Printf("序号：%d  %d-%s-%d-%s-%s-%s\n",i,v.Id,v.Name,v.Age,v.Gender,v.Email,v.Phone)
	}

}