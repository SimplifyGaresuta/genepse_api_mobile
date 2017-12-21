package location

import (
	"encoding/json"
	"genepse_api/src/infra/cache"
	"io"
	"io/ioutil"
	"log"

	"github.com/garyburd/redigo/redis"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

const key = "locations"

func UpdateLocation(con *redis.Conn, userID string, r io.ReadCloser) (err error) {
	location, err := decode(r)
	if err != nil {
		return
	}
	if err = cache.GeoAdd(con, key, userID, location.Latitude, location.Longitude); err != nil {
		c := *con
		log.Println("geoadd時にエラー", c.Err())
		return
	}
	return
}

func decode(r io.ReadCloser) (*Location, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	location := &Location{}
	if err := json.Unmarshal(bytes, &location); err != nil {
		return nil, err
	}
	return location, nil
}
