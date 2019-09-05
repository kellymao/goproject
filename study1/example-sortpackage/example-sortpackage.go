package main

import "sort"
import "fmt"

// 自定义 字符串切片排序


// 内置的 字符串切片排序
func sort_str(){
	var str sort.StringSlice

	str = sort.StringSlice{
		"apple","banana","taozi","xigua","why","cm",
	}


	sort.Sort(str)

	fmt.Println(str)  // [apple banana cm taozi why xigua]



}

func sort_str1(){

	var str []string = []string{"apple","banana","taozi","xigua","why","cm",}

	sort.Strings(str)

	fmt.Println(str)  // [apple banana cm taozi why xigua]


}

// 内置的整数切片排序

func sort_int(){

	var i sort.IntSlice  = sort.IntSlice{6,8,9,3,2,7}

	sort.Sort(i)

	fmt.Println(i) // [2 3 6 7 8 9]


}

func sort_int1(){


	var i []int = []int{6,8,9,3,2,7}

	sort.Ints(i)

	fmt.Println(i)

}


// 结构体排序

func sort_hero(){

type HeroKind int
const (
	None = iota
	Tank
	Assassin
	Mage
)


type Hero struct{
	name string
	kind HeroKind


}

var hero_se []*Hero = []*Hero{

	&Hero{"张飞",Mage},
	&Hero{"吕布",Assassin},
	&Hero{"关于",Tank},
	&Hero{"曹操",Assassin},
	&Hero{"张辽",Mage},
	&Hero{"孙权",Assassin},
	&Hero{"周瑜",Tank},




}

sort.Slice(hero_se,func(i, j int)bool{

	return  hero_se[i].kind >  hero_se[j].kind


})

for index,v := range hero_se{

	fmt.Println(index, *v )

}





}

func main(){


	sort_str()
	sort_str1()

	sort_int()
	sort_int1()

	sort_hero()
}










