package location

import (
	"encoding/json"
	"genepse_api/src/infra/cache"
	"io"
	"io/ioutil"
	"strconv"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

const key = "locations"

func UpdateLocation(userID string, r io.ReadCloser) (err error) {
	location, err := decode(r)
	if err != nil {
		return
	}
	if err = cache.GeoAdd(key, userID, location.Latitude, location.Longitude); err != nil {
		return
	}
	return
}

func GetNearUsers(userID string, distance int) (userIDs []int, err error) {
	ids, err := cache.GeoRadiusByMember(key, userID, distance)
	var i int
	for _, id := range ids {
		i, err = strconv.Atoi(id)
		if err != nil {
			return
		}
		userIDs = append(userIDs, i)
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
