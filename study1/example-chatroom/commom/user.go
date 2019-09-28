package main

import "fmt"


var pool string


func init(){


	pool = "abc123"
	fmt.Println("test")

}

type Favorite struct{

	Own string  `json:"own"`
	Like string `json:"like"`


}


type Person struct {

	Name string  `json:"name"`
	Age int	     `json:"age"`
	Fav Favorite `json:"fav"`


}


