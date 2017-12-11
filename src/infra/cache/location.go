package cache

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func Tes() (err error) {
	con.Do("SET", "nakao", "ryoryoryo")
	s, err := redis.String(con.Do("GET", "nakao"))
	if err != nil {
		return
	}
	fmt.Println("ヴァリューは", s)
	return
}
