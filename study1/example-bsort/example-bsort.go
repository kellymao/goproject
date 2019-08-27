package main

import "fmt"

// 冒泡排序算法, 把最小的放第一个

func bsort_min2max(a []int){


	for i:=0; i<len(a);i++{

		for j:=i+1;j<len(a);j++{

			if a[i] > a[j] {
				a[i],a[j] = a[j],a[i]


			}
		}


	}

}


func bsort_max2min(a []int){


	for i:=0; i<len(a);i++{

		for j:=i+1;j<len(a);j++{

			if a[i] < a[j] {
				a[i],a[j] = a[j],a[i]


			}
		}


	}

}

// 快速排序法
/*
对于一个数组，随机把一个元素定下来，左边的比他小，右边的比他大 .分成两数组，在递归处理

数组元素==1 ，则退出

 */

func fastsort(a []int, left , right int ){

	if left>= right{
		return
	}
	base := a[left]
	k := left

	for i:= left+1 ; i<right; i++{
		if a[i] < base {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	a[k] = base
	fastsort(a,left,k-1)
	fastsort(a,k+1,right)


}

func main(){

	a:= [...]int{10,8,50,3,2,89,908,76}
	bsort_min2max(a[:])
	fmt.Println(a)   // 修改了 a 本身的值 [2 3 8 10 50 76 89 908]

	bsort_max2min(a[:])
	fmt.Println(a)

	fastsort(a[:],0 ,len(a))
	fmt.Println(a)

}