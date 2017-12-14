package cache

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

var addr string
var con redis.Conn

func DialRedis() (err error) {
	if isDevelop() {
		addr = "127.0.0.1:6379"
	} else {
		addr = "10.146.0.2:6379"
	}
	con, err = redis.Dial("tcp", addr)
	return
}

func CloseRedis() {
	con.Close()
}

func GetErr() error {
	return con.Err()
}

func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}
