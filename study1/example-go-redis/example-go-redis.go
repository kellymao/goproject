package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)


var redisdb *redis.Client

// 创建redis 客户端连接
func init_redis() error {

	redisdb = redis.NewClient(

		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			PoolSize: 5,
		},

		)


	_,err:=redisdb.Ping().Result()

	if err!=nil{
		return err

	}

	return nil




}

//  String 操作

func op_string(){

	// 0 表示不过期

	err := redisdb.Set("why","not",0).Err()
	if err!=nil{
		fmt.Println("set 失败")

	}

	// 设置name 一秒钟过期

	err=redisdb.Set("name","zhangsan",time.Second).Err()
	if err!=nil{
		fmt.Println("set 失败")


	}


	val ,err:= redisdb.Get("name").Result()

	if err!=nil{
		fmt.Println("get 失败")


	}
	fmt.Println("第一次取name ",val)

	err=redisdb.Set("age",17,time.Second).Err()
	if err!=nil{
		fmt.Println("set age 失败")


	}

	redisdb.Incr("age")
	redisdb.Incr("age")
	redisdb.Incr("age")  // 自增

	redisdb.Decr("age")

	age,err:=redisdb.Get("age").Result()
	if err!=nil{
		fmt.Println("get 失败")


	}
	fmt.Println(age)

	time.Sleep(time.Second * 2)

	val ,err= redisdb.Get("name").Result()

	if err!=nil{
		fmt.Println("get name失败",err.Error())  // 过期了 报错 redis: nil


	}
	fmt.Println("第二次取name ",val)

/*
   第一次取name  zhangsan
   19
   get name失败 redis: nil
   第二次取name
*/

}

// List 操作

func op_list(){

	// 在名称为 fruit 的list头添加一个值为apple的 元素
	val,err:=redisdb.LPush("fruit","apple").Result()
	if err!=nil{

		fmt.Println("LPush 失败")
	}

	fmt.Println(val)


	// 在名称为 fruit 的list尾添加一个值为huolongguo的元素
	err = redisdb.RPush("fruit","huolongguo").Err()

	if err!=nil{

		fmt.Println("RPush 失败")
	}

	// 在名称为 fruit 的list头添加一个值为banana的 元素
	redisdb.LPush("fruit","banana")

	// 返回名称为 fruit 的list的长度
	list_len,err:=redisdb.LLen("fruit").Result()

	if err!=nil{

		fmt.Println("LLen 失败")
	}

	fmt.Println("列表fruit的长度：",list_len)


	// 获取列表的值，返回一个 string的切片
	valstr, err := redisdb.LRange("fruit", 0, -1).Result()


	if err!=nil{

		fmt.Println("lrange 失败")
	}
	fmt.Println(valstr)

	// 返回并删除名称为 fruit 的list中的尾元素

	popval,err:=redisdb.RPop("fruit").Result()

	if err!=nil{

		fmt.Println("rpop 失败")
	}
	fmt.Println(popval)
	// 获取列表的值，返回一个 string的切片
	valstr, err = redisdb.LRange("fruit", 0, -1).Result()


	if err!=nil{

		fmt.Println("lrange 失败")
	}
	fmt.Println(valstr)


/*

   1
   列表fruit的长度： 3
   [banana apple huolongguo]
   huolongguo
   [banana apple]
 */
}

// set 操作

func set_list(){


	// 向 fruitset 中添加元素
	err:=redisdb.SAdd("fruitset","apple").Err()
	if err!=nil{

		fmt.Println("sadd 失败")
	}

	redisdb.SAdd("fruitset","banana")
	redisdb.SAdd("fruitset","huolongguo")
	redisdb.SAdd("fruitset","xihongshi")


	// 获取指定集合的所有元素
	fruitset,_:=redisdb.SMembers("fruitset").Result()
	fmt.Println("水果的集合:",fruitset)

	// 判断元素是否在集合中
	is ,err:= redisdb.SIsMember("fruitset","huolongguo").Result()
	fmt.Println("火龙果是水果吗？",is)

	redisdb.SAdd("shucaiset","xihongshi")
	redisdb.SAdd("shucaiset","baicai")


	// 求交集
	inter,err:=redisdb.SInter("fruitset","shucaiset").Result()

	fmt.Println(inter)


	/*
	水果的集合: [apple huolongguo xihongshi banana]
	火龙果是水果吗？ true
	[xihongshi]


	 */


}

// hash 操作

func set_hash(){


	// 向名称为 person 的 hash 中添加元素 name
	err:=redisdb.HSet("person","name","zhangsan").Err()
	if err!=nil{
		fmt.Println("hset error")
	}

	// 批量地向名称为 person2 的 hash 中添加元素 name 和 age

	err=redisdb.HMSet("person2",map[string]interface{}{"name":"zhuxi","age":10,"sex":"nan"}).Err()

	if err!=nil{
		fmt.Println("hmset error")
	}

	// 获取hash 字段的个数
	Helen,err:=redisdb.HLen("person2").Result()

	if err!=nil{
		fmt.Println("hlen error")
	}
	fmt.Println(Helen)



	val,_ :=redisdb.HGet("person","name").Result()
	fmt.Println(val)

	// 批量获取名为 person2 的 hash 中的指定字段的值.

	valor, _:=redisdb.HMGet("person2","name","age","sex").Result()
	fmt.Println(valor)

	// 获取hash 的所有值，结果是一个map
	valall,_:=redisdb.HGetAll("person2").Result()
	fmt.Println(valall)

	// 删除名为 person2 的 age 字段
	redisdb.HDel("person2","age")

	valall,_=redisdb.HGetAll("person2").Result()
	fmt.Println(valall)

	// redis 设置hash过期时间
	redisdb.Expire("person2",time.Second)

	time.Sleep(time.Second * 2)

	valall,_=redisdb.HGetAll("person2").Result()
	fmt.Println(valall)

	/*

	3
	zhangsan
	[zhuxi 10 nan]
	map[name:zhuxi age:10 sex:nan]
	map[name:zhuxi sex:nan]
	map[]
	 */




}


// 连接池

/*
初始化的时候加入参数 	PoolSize: 5,
 */

func connectPool() {
	wg := sync.WaitGroup{}
	wg.Add(15)

	for i := 0; i < 105; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				_ = redisdb.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				_, _ = redisdb.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf( "PoolStats, TotalConns: %+v \n", redisdb.PoolStats() )
		}()
	}

	wg.Wait()

	/*

	上面的例子启动了 10 个 routine 来不断向 redis 读写数据, 然后我们通过 client.PoolStats() 获取连接池的信息.
	运行这个例子, 输出如下:

	PoolStats, TotalConns: &{Hits:2938 Misses:5 Timeouts:0 TotalConns:5 IdleConns:2 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2955 Misses:5 Timeouts:0 TotalConns:5 IdleConns:2 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2959 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2960 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2962 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2964 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2966 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2966 Misses:5 Timeouts:0 TotalConns:5 IdleConns:2 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2968 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2970 Misses:5 Timeouts:0 TotalConns:5 IdleConns:1 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2970 Misses:5 Timeouts:0 TotalConns:5 IdleConns:2 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2972 Misses:5 Timeouts:0 TotalConns:5 IdleConns:2 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2972 Misses:5 Timeouts:0 TotalConns:5 IdleConns:3 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2974 Misses:5 Timeouts:0 TotalConns:5 IdleConns:4 StaleConns:0}
	PoolStats, TotalConns: &{Hits:2996 Misses:5 Timeouts:0 TotalConns:5 IdleConns:5 StaleConns:0}


	通过输出可以看到, 此时最大的连接池数量确实是 5 了(不加连接池的话是40), 并且一开始时, 因为 coroutine 的数量大于 5,
	会造成 redis 连接不足的情况(反映在 IdleConns 上就是前几次的输出 IdleConns 一直是 1 ),
	当某个 coroutine 结束后, 会释放此 redis 连接, 因此 IdleConns 会增加.


	在linux 每隔0.1 秒查看redis 连接数： 只有5个链接


	watch -n 0.1  -d "netstat -antp | grep redis"

	 */
}


// 订阅

func main(){

	err:=init_redis()

	if err!=nil{
		return
	}

	//op_string()
	//op_list()
	//set_list()

	//set_hash()

	connectPool()

}