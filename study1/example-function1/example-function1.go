package main

import "fmt"


// 阶乘 - 递归
func jiecheng(n uint64) uint64 {

	if n ==1 {

		return 1

	} else {

		return n * jiecheng(n-1)
	}
}

// 闭包, 函数和外部环境的整体

func Adder() func(int) int {
	var x int

	return func(n int) int {

		x += n
		return x
	}


}

func main(){

	 fmt.Println(jiecheng(10))

	 f := Adder()
	 fmt.Println(f(10),f(10))



}