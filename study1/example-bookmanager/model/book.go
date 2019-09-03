package model

import "errors"

type Book struct{
	Name string
	Total int
	Author string
	Createtime string



}

func CreateBook(name string,total int,author string,createtime string) *Book {

	b :=&Book{
		Name:name,
		Total:total,
		Author:author,
		Createtime,createtime,



	}

	return b



}

func (b *Book) Canborrow(borrow_num int ) bool {

	return  b.Total >=  borrow_num

}



func  (b *Book) Borrow(borrow_num int) (err error){

	if b.Canborrow(borrow_num){

		b.Total-=borrow_num
	}else{

		err = errors.New("this book is not enough.")
		return
	}



}
