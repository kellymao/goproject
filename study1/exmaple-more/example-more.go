package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var mu sync.RWMutex
var wg sync.WaitGroup


var loadonce sync.Once
var m map[string]string

func add(){

	defer wg.Done()
	for i:=0;i<10000;i++{
		mu.Lock()
		x+=1
		mu.Unlock()
	}

}

func write(){

	defer wg.Done()
	for i:=0;i<10000;i++{
		mu.Lock()
		x+=1
		mu.Unlock()

	}

}

func read(){

	defer wg.Done()
	for i:=0;i<10000;i++{
		mu.RLock()
		time.Sleep(time.Microsecond)
		mu.RUnlock()

	}


}

func loadm(){

	m = map[string]string{
		"p1":"zhangsan",
		"p2":"mazi",
	}
	fmt.Println(m)

}


func initproc(name string) string {

	defer wg.Done()
	loadonce.Do(loadm)

	fmt.Println(m[name])
	return m[name]



}


func timeconvert(){


	now := time.Now()
	now.Year()
	now.Month()

	fmt.Println(now.Format("2006-01-02 15:04:05"))

	t1, err := time.Parse("2006/01/02 15:04:05", "1998/09/12 23:36:10")
	if err!=nil{
		fmt.Println("convert failed")
	}

	fmt.Println(t1)
	fmt.Println(t1.Format("2006-01-02 15:04:05"))

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "1998/09/12 23:36:10", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)

	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))


	fmt.Println("unix 时间戳",now.Unix())  // unix 时间戳 1574150524
	fmt.Println("unix 纳秒时间戳",now.UnixNano()) // unix 纳秒时间戳 1574150524460694875


	a := time.Unix(1574150524,0)
	fmt.Println(a) // 2019-11-19 16:02:04 +0800 CST

	now = time.Now()
	fmt.Println("当前的时间",now)
	fmt.Println("十分钟之后",now.Add(time.Minute * 10))

	time.Sleep(time.Second * 10 )

	now2:= time.Now()
	fmt.Println("时间差",now2.Sub(now))
	fmt.Println("时间差",time.Since(now))



}

//func main(){
	/*
	wg.Add(3)
	go initproc("p1")
	go initproc("p2")
	go initproc("p1")

	wg.Wait()

	fmt.Println(x)


	 */

	//timeconvert()

//}
