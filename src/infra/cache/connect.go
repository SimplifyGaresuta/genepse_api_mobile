package cache

import (
	"github.com/garyburd/redigo/redis"
)

// TODO 環境変数に
const (
	Host = "127.0.0.1"
	//Host = "35.200.63.13"
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
