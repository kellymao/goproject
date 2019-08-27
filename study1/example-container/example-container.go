package main

import (

	"fmt"
	"sync"
	_ "sort"

	"math/rand"

)
import "container/list"


// 数组的两种遍历方式



// 数组是值传递
func test2(a [6]int){

	a[0] = 3

}

func test2_ptr(a *[6]int){

	(*a)[0]=3
}

// 二位数组
func testarray(){

	var s [2][5]int = [2][5]int{
		{1,2,3,5,6}, {7,6,8,9,8}}

	for row,v:= range s{
		for col,k:= range v{
			fmt.Println(row,col,k)

		}


	}
}


// 切片拷贝copy 和 赋值的去区别

func slice_copy(){


	const maxnum = 10
	src_slice := make([]int,maxnum)

	src_slice_2 := src_slice

	dst_slice := make([]int,maxnum)

	for i:=0; i < maxnum; i++{
		src_slice[i] = i
	}

	copy(dst_slice,src_slice)

	src_slice_2[0] =999

	fmt.Println("src_slice result is ", src_slice)
	fmt.Println("src_slice_2 result is ", src_slice_2)
	fmt.Println("dst_slice result is ", dst_slice)


	copy(dst_slice[5:8],src_slice[5:8])

	fmt.Println("src_slice result is ", src_slice)
	fmt.Println("src_slice_2 result is ", src_slice_2)
	fmt.Println("dst_slice result is ", dst_slice)



}

// map 的初始化和遍历
func process_map(){

	// 取出map 的所有键
	var keys []string
	map_var :=make(map[string]string)
	map_var["English"]="england"
	map_var["Chinese"]="china"

	for k:= range map_var{

		keys=append(keys,k)
	}

	fmt.Println(keys)

}

// map  的删除

func del_map(){

	var map_str map[string]int = make(map[string]int,100)

	map_str["English"] = 1
	map_str["Chinese"] =2

	fmt.Println(map_str)
	delete(map_str,"English")
	fmt.Println(map_str)
	map_str = make(map[string]int)
	fmt.Println(map_str)


}

//sync.map 操作


func sync_map() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

// 列表的操作  http://c.biancheng.net/view/35.html


func process_list() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	//遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}


// 切片的查找和排序
func process_sorts(){

	//var a []int = []int{1,9,6,78,987,67}
	//sort.Ints(a)
	//fmt.Println(a)

	//index:=sort.SearchInts(sort.Ints(a),67)
	//fmt.Println(index)
}

// 二维map的查找和更新
func map_erwei(){

	var map_str map[string]map[string]string = make(map[string]map[string]string,10)

	if v,ok:= map_str["zhangsan"] ; ok {
		v["user"] = "zhs"
		v["pwd"] = "asd"

	}else{
		map_str["zhangsan"] = make(map[string]string)
		map_str["zhangsan"]["user"] = "zs"
		map_str["zhangsan"]["pwd"] = "123"
	}

	fmt.Println(map_str)

	// 优化版
	if _,ok := map_str["zhangsan"] ; !ok{
		map_str["zhangsan"] = make(map[string]string)
	}

	map_str["zhangsan"]["user"] = "zhs"
	map_str["zhangsan"]["pwd"] = "asd"

	fmt.Println(map_str)



}

func testmap(){
	var lock sync.Mutex
	var a map[int]int

	a = make(map[int]int,10)
	a[0] = 1
	a[5] = 8

	for i:=0; i<2; i++{

		go func(b map[int]int){
			lock.Lock()
			b[0] = rand.Intn(100)
			lock.Unlock()
		}(a)


	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

}


func testRWlockmap(){

	var lock sync.RWMutex
	var a map[int]int

	a = make(map[int]int,10)


	for i:=0; i<10; i++{

		go func(b map[int]int){
			lock.Lock()             // 写锁不允许其他线程读和写
			//b[i] = rand.Intn(100)
			fmt.Println(i)
			lock.Unlock()
		}(a)


	}


	for i:=0; i<10;i++{
		go func(b map[int]int){
			lock.RLock()         // 读锁允许其他线程读 ，不允许其他线程写
			fmt.Println(b)
			lock.RUnlock()
		}(a)

	}





}

func main() {

	var a [6]int

	test2(a)
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[0])

	}
	test2_ptr(&a)
	fmt.Println(a)

	testarray()

	slice_copy()

	process_map()

	del_map()

	process_list()
	process_sorts()
	map_erwei()
	testmap()
	testRWlockmap()

}