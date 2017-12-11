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

func GeoAdd(key, name string, lat, lon float64) (err error) {
	//GEOADD towers 139.745464 35.658582 "Tokyo Tower"
	_, err = con.Do("GEOADD", key, lon, lat, name)
	return
}
