package cache

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

const (
	Host = "127.0.0.1"
	Port = "6379"
)

var con redis.Conn

func DialRedis() (err error) {
	fmt.Println("b")
	con, err = redis.Dial("tcp", Host+":"+Port)
	fmt.Println("c")
	if err != nil {
		return
	}
	defer con.Close()
	con.Do("SET", "hey", "oi")
	fmt.Println("d")
	s, err := redis.String(con.Do("GET", "hey"))
	fmt.Println("e")
	if err != nil {
		return
	}
	fmt.Println("ヴァリューは", s)
	return
}
