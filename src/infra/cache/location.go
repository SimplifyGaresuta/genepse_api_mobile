package cache

import "github.com/garyburd/redigo/redis"

func GeoAdd(con *redis.Conn, key, member string, lat, lon float64) (err error) {
	c := *con
	_, err = c.Do("GEOADD", key, lon, lat, member)
	return
}

func GeoRadiusByMember(con *redis.Conn, key, member string, radius int) (members []string, err error) {
	c := *con
	members, err = redis.Strings(c.Do("GEORADIUSBYMEMBER", key, member, radius, "km", "ASC"))
	return
}

func GeoDist(con *redis.Conn, key, member1, member2, unit string) (dist float64, err error) {
	c := *con
	dist, err = redis.Float64(c.Do("GEODIST", key, member1, member2, unit))
	return
}
