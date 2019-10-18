package main

import "fmt"
import "github.com/garyburd/redigo/redis"
import "time"
import "encoding/json"

import "reflect"

func process_err(err error){

	if err !=nil{

		fmt.Println(err)
		panic(err)

	}


}

func json_to_redis(conn redis.Conn){

	val :=map[string]string{"name":"zhangsan",
		"age":"10",
		}

	v,err:=json.Marshal(val)

	process_err(err)
	n,err:=conn.Do("set","key1",v)

	process_err(err)

	if n == int64(1) {
		fmt.Println("success")
	}

	fmt.Println(n,reflect.TypeOf(n))  // OK string
	rel,err:=redis.Bytes(conn.Do("get","key1"))
	process_err(err)
	fmt.Println(rel)  // [123 34 97 103 101 34 58 34 ]


	var r map[string]string = make(map[string]string)

	err= json.Unmarshal(rel,&r)

	process_err(err)

	fmt.Println(r)        // map[age:10 name:zhangsan]

	fmt.Println(r["name"])  // zhangsan





}

func list_operate(conn redis.Conn){

	n,err:=conn.Do("lpush","dblist","mysql","oracle","mongo","中国","test")
	process_err(err)

	fmt.Println(n)  //7

	rel,err:=redis.String(conn.Do("lpop","dblist"))

	process_err(err)
	fmt.Println(rel) // mongo

	relrange,err := redis.Values(conn.Do("lrange","dblist",0,10))

	process_err(err)

	fmt.Println(reflect.TypeOf(relrange))  // []interface {}

	for _,v := range relrange {

		fmt.Println(reflect.TypeOf(v)) //[]uint8
		fmt.Println(v.([]byte)) // [111 114 97 99 108 101] 通过接口的断言来转换类型

		fmt.Println(string(v.([]byte))) // mongo


	}



}

func hash_operate(conn redis.Conn){

	conn.Do("hmset","table","key1","val1","key2",10)

	rel,_:=conn.Do("hget","table","key1")

	fmt.Println(reflect.TypeOf(rel)) // []uint8

	c_rel,_:=redis.String(conn.Do("hget","table","key1"))
	fmt.Println(c_rel) // val1

	c_rel1,_:=redis.Int64(conn.Do("hget","table","key2"))
	fmt.Println(c_rel1) // 10

	rel2,_:=conn.Do("HMGET","table","key1","key2")
	fmt.Println(reflect.TypeOf(rel2))  // []interface {}


	c_rel2,_ :=redis.Strings(conn.Do("HMGET","table","key1","key2"))
	fmt.Println(c_rel2)  // [val1 10]


	rel3,_:=conn.Do("hgetall","table")
	fmt.Println(reflect.TypeOf(rel3)) // []interface {}

	c_rel3,_:=redis.Values(conn.Do("hgetall","table"))
	fmt.Println(c_rel3) // [[107 101 121 49] [118 97 108 49] [107 101 121 50] [49 48]]

	for _,v := range c_rel3{

		fmt.Println(string(v.([]byte)))
		/*
		key1
		val1
		key2
		10

		 */
	}

	c_rel5 ,_:= redis.Strings(conn.Do("hgetall","table"))
	fmt.Println(c_rel5) // [key1 val1 key2 10]


}

func main(){

	conn,err:=redis.Dial("tcp","127.0.0.1:6379")

	process_err(err)

	defer conn.Close()

	_,err = conn.Do("set","name","wd","ex",5)

	process_err(err)

	username,err:=redis.String(conn.Do("GET","name"))

	process_err(err)

	_,err = conn.Do("expire","name","35")
	process_err(err)

	_ ,err =conn.Do("HSET","hashuser","user1","mao")

	process_err(err)

	_,err = conn.Do("HMSET","hashuser","user1","maomao","user2","zhangsan","user3","wangwu")

	process_err(err)

	fmt.Println(username)

	_,err = conn.Do("expire","hashuser",30)
	process_err(err)

	hashuser,err := redis.String(conn.Do("HGET","hashuser","user1"))

	process_err(err)

	fmt.Println(hashuser)

	hashusers,err := redis.Strings(conn.Do("HMGET","hashuser","user1","user2"))

	process_err(err)

	fmt.Println( reflect.TypeOf(hashusers))  //[]string
	fmt.Println(hashusers)                   // [maomao zhangsan]

	time.Sleep(3 * time.Second)

	is_exist,err:=redis.Bool(conn.Do("exists","hashuser"))

	process_err(err)

	fmt.Println(is_exist)

	json_to_redis(conn)

	list_operate(conn)

	hash_operate(conn)

}