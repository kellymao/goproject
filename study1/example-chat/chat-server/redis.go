package main
import "time"

import "github.com/garyburd/redigo/redis"


var pool *redis.Pool


func init_redis(idleconn,maxconn int ,idletimeout time.Duration,addr string) {

	pool = &redis.Pool{

		MaxIdle:     idleconn,
		MaxActive:   maxconn,
		IdleTimeout: idletimeout,
		Dial: func() (redis.Conn, error) {

			return redis.Dial("tcp", addr)

		},
	}
}

func get_conn() redis.Conn {

	return pool.Get()

}

