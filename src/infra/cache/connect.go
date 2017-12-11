package cache

import (
	"github.com/garyburd/redigo/redis"
)

const (
	Host = "127.0.0.1"
	Port = "6379"
)

var con redis.Conn

func DialRedis() (err error) {
	con, err = redis.Dial("tcp", Host+":"+Port)
	return
}

func CloseRedis() {
	con.Close()
}
