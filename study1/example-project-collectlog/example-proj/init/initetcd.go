package main



import (

"context"
"fmt"

_ "strconv"
"time"
"go.etcd.io/etcd/clientv3"
 "encoding/json"
)


func initadd(){


	type Valuestr struct{

		Topic string
		Logfile string


	}

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


	var v1 Valuestr = Valuestr{

		Topic:"mysql",
		Logfile:"/var/log/mysqld.log",

	}

	var v2 Valuestr = Valuestr{

		Topic:"dmesg",
		Logfile:"/var/log/dmesg",

	}

	data1,err :=json.Marshal(v1)
	if err!=nil{
		panic(err)
	}

	fmt.Println(data1)
	data2,_:=json.Marshal(v2)

	putResp,err := kv.Put(ctx,"/proj/localhost/collect/mysql",string(data1))
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",putResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:23 raft_term:4  <nil>}


	kv.Put(ctx,"/proj/localhost/collect/dmesg",string(data2))


	getResp,err := kv.Get(ctx,"/proj/localhost/collect/",clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",getResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:24 raft_term:4  [key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ] %!s(bool=false) %!s(int64=1)}
	fmt.Printf("%v \n",getResp.Kvs) //[key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ]


	// /workspace/bin/init



}


func update(){


	type Valuestr struct{

		Topic string
		Logfile string


	}

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


	var v2 Valuestr = Valuestr{

		Topic:"dmesg",
		Logfile:"/var/log/message1",

	}

	// 更新key dmesg的收集位置
	data2,_:=json.Marshal(v2)

	putResp,err := kv.Put(ctx,"/proj/localhost/collect/dmesg",string(data2))
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",putResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:23 raft_term:4  <nil>}






	getResp,err := kv.Get(ctx,"/proj/localhost/collect/",clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",getResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:24 raft_term:4  [key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ] %!s(bool=false) %!s(int64=1)}
	fmt.Printf("%v \n",getResp.Kvs) //[key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ]


	// /workspace/bin/init


}

func delete(){

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
	_,_ = kv.Delete(ctx,"/proj/localhost/collect/dmesg")

}

func add(){

	type Valuestr struct{

		Topic string
		Logfile string


	}

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




	var v3 Valuestr = Valuestr{

		Topic:"newtt",
		Logfile:"/var/log/mysqld.log",

	}



	// 新添加一个key测试
	data3,_:=json.Marshal(v3)
	kv.Put(ctx,"/proj/localhost/collect/newtt",string(data3))


	getResp,err := kv.Get(ctx,"/proj/localhost/collect/",clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",getResp) // &{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:24 raft_term:4  [key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ] %!s(bool=false) %!s(int64=1)}
	fmt.Printf("%v \n",getResp.Kvs) //[key:"/job/v3" create_revision:23 mod_revision:24 version:2 value:"push the box" ]


	// /workspace/bin/init


}

func main(){

	//delete()
	update()
}

