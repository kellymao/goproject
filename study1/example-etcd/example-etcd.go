package main

import (

	"context"
	"fmt"

	_ "strconv"
	"time"
	"go.etcd.io/etcd/clientv3"
)


func main(){

    //连接etcd:
	cli,err:=clientv3.New(clientv3.Config{

		Endpoints: []string{"http://172.18.8.178:2379"} ,
		DialTimeout: 5 * time.Second,
	})

	if err!=nil{
		panic(err)
	}

	defer cli.Close()

	kv := clientv3.NewKV(cli)

	//设置一个超时的context
	ctx,_:= context.WithTimeout(context.Background(),5*time.Second)

	//新建一个kv

	putResp,err := kv.Put(ctx,"/job/v3","push the box")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",putResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:23 raft_term:4  <nil>}

	getResp,err := kv.Get(ctx,"/job/",clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",getResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:24 raft_term:4  [key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ] %!s(bool=false) %!s(int64=1)}
	fmt.Printf("%v",getResp.Kvs) //[key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ]

	for _, ev := range getResp.Kvs{
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}


	// 更新一个kv

	putnewResp,err := kv.Put(ctx,"/job/v3","push the new",clientv3.WithPrevKV())  //withPrevKV()是为了获取操作前已经有的key-value
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",putnewResp)
	//&{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:28 raft_term:4  key:"/job/v3" create_revision:23 mod_revision:27 version:5 value:"push the box" }
	fmt.Printf("%v\n",putnewResp.PrevKv)
	// key:"/job/v3" create_revision:23 mod_revision:31 version:9 value:"push the box"

	getnewResp,err:=kv.Get(ctx,"/job/v3")
	if err != nil{
		panic(err)
	}
	fmt.Println("修改后的值")

	fmt.Printf("%s\n",getnewResp)
	fmt.Printf("%s\n",getnewResp.Kvs) //[key:"/job/v3" create_revision:23 mod_revision:34 version:12 value:"push the new" ]

	for _, ev := range getnewResp.Kvs{
		fmt.Printf("%s:%s\n", ev.Key, ev.Value) // /job/v3:push the new
	}


	// 启动一个新线程去更新
	go func(cli  *clientv3.Client){

		time.Sleep(5 * time.Second)
		_, _ = cli.Put(context.Background(),"/job/v3","ttt")

	}(cli)

	// 监控watch一个key

	wc := cli.Watch(context.Background(), "/job/", clientv3.WithPrefix(),clientv3.WithPrevKV())
	for v := range wc {
		if v.Err() != nil {
			panic(err)
		}
		for _, e := range v.Events {
			fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
		}
	}

	/*
	type:PUT
    kv:key:"/job/v3" create_revision:23 mod_revision:37 version:15 value:"ttt"   prevKey:key:"/job/v3" create_revision:23 mod_revision:36 version:14 value:"push the new"

	*/


}
