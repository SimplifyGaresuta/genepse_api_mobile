package cache

import (
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	addr string
	//con  redis.Conn
	pool *redis.Pool
)

func SetPool() {
	if isDevelop() {
		addr = devAddr
	} else {
		addr = proAddr
	}
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
	return
}

func GetConn() redis.Conn {
	return pool.Get()
}

/*
func CloseConn() {
	con.Close()
}
*/
/*
func GetErr() error {
	return con.Err()
}
*/
func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}
