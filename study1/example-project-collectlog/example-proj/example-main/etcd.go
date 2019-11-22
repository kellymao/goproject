package main

import (

	"context"
	"fmt"

	_ "strconv"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"encoding/json"

)

import "github.com/astaxie/beego/logs"

type Handler struct{

		Handles map[string]*Collectconf

}

var handler Handler = Handler{
	make(map[string]*Collectconf),
}



func initetcd(){

	//连接etcd:
	cli,err:=clientv3.New(clientv3.Config{

		Endpoints: []string{appconf.etcd_addr} ,
		DialTimeout: 5 * time.Second,
	})

	if err!=nil{
		logs.Error(err)
		return
	}

	defer cli.Close()

	kv := clientv3.NewKV(cli)

	//设置一个超时的context
	ctx,_:= context.WithTimeout(context.Background(),5*time.Second)


	getResp,err := kv.Get(ctx,"/proj/localhost/collect/",clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil{
		logs.Error(err)
		return
	}
	//fmt.Printf("%s\n",getResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:24 raft_term:4  [key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ] %!s(bool=false) %!s(int64=1)}
	//fmt.Printf("%v \n",getResp.Kvs) //[key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ]

	for _, ev := range getResp.Kvs{
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)


		newtask(ev)



	}




	go watch()









}

func newtask(ev *mvccpb.KeyValue){
	var collect *Collectconf = &Collectconf{
		Exitchan:make(chan string),
	}

	err:=json.Unmarshal(ev.Value,&collect)
	if err!=nil{
		logs.Error("unmarshal collect error:",err)
		return

	}

	handler.Handles[string(ev.Key)] = collect

	// collects = append(collects,collect)

	go tailfile(collect)
}


func watch(){
	//连接etcd:
	cli,err:=clientv3.New(clientv3.Config{

		Endpoints: []string{appconf.etcd_addr} ,
		DialTimeout: 5 * time.Second,
	})

	if err!=nil{
		logs.Error(err)
		return
	}

	defer cli.Close()



	// 更新一个kv

	// 监控watch一个key

	wc := cli.Watch(context.Background(), "/proj/localhost/collect/", clientv3.WithPrefix(),clientv3.WithPrevKV())
	for v := range wc {
		if v.Err() != nil {
			logs.Error(v.Err())
			return
		}
		for _, e := range v.Events {
			//fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)

			if e.Type == mvccpb.DELETE {
					v,ok := handler.Handles[string(e.Kv.Key)]

					if ok{
						v.Exitchan <- "exit"

						delete(handler.Handles,string(e.Kv.Key))

					}else{
						logs.Error("删除的key不存在",string(e.Kv.Key))
					}



			}

			if e.Type == mvccpb.PUT{
				v,ok := handler.Handles[string(e.Kv.Key)]

				if ok {

					logs.Error("修改一个key的配置",string(e.Kv.Key))
					fmt.Printf("%s \n",v)
					v.Exitchan <- "exit"
					fmt.Printf("%s \n",v)

					newtask(e.Kv)


				}else{
					logs.Error("新加一个key的配置",string(e.Kv.Key))

					newtask(e.Kv)

				}

			}


		}
	}

	/*

	fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
			type:PUT
		    kv:key:"/job/v3" create_revision:23 mod_revision:37 version:15 value:"ttt"   prevKey:key:"/job/v3" create_revision:23 mod_revision:36 version:14 value:"push the new"

	*/
}

