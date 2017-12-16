package cache

import "github.com/garyburd/redigo/redis"

func GeoAdd(key, member string, lat, lon float64) (err error) {
	_, err = con.Do("GEOADD", key, lon, lat, member)
	return
}

func GeoRadiusByMember(key, member string, radius int) (members []string, err error) {
	members, err = redis.Strings(con.Do("GEORADIUSBYMEMBER", key, member, radius, "km", "ASC"))
	return
}

func GeoDist(key, member1, member2, unit string) (dist float64, err error) {
	dist, err = redis.Float64(con.Do("GEODIST", key, member1, member2, unit))
	return
}
