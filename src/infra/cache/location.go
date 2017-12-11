package cache

import "github.com/garyburd/redigo/redis"

func GeoAdd(key, name string, lat, lon float64) (err error) {
	_, err = con.Do("GEOADD", key, lon, lat, name)
	return
}

func GeoRadiusByMember(key, name string, distance int) (names []string, err error) {
	names, err = redis.Strings(con.Do("GEORADIUSBYMEMBER", key, name, distance, "km", "ASC"))
	return
}
