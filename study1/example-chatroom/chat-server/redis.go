package main


import "github.com/garyburd/redigo/redis"
import "time"

var pool *redis.Pool



func InitRedis(addr string,idleconn,maxconn int,idletimeout time.Duration){


	pool :=&redis.Pool{
		MaxIdle: idleconn ,
		MaxActive:maxconn,
		IdleTimeout:idletimeout,
		Dial: func()(redis.Conn,error){


			return redis.Dial("tcp",addr)

		},



	}

	return

}


func Getconn() redis.Conn{


	return pool.Get()


}


func Putconn(conn redis.Conn){

	conn.Close()
}

