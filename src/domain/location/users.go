package location

import (
	"genepse_api/src/infra/cache"
	"genepse_api/src/infra/orm"
	"strconv"
)

type Response struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
}

func GetNearUsers(userID string, distance int) (response *Response, err error) {
	ids, err := getUserIDs(userID, distance)
	if err != nil {
		return
	}

	users, err := getUsers(ids, userID)

	response = &Response{
		Users: users,
	}
	return
}

func getUserIDs(userID string, distance int) (userIDs []string, err error) {
	userIDs, err = cache.GeoRadiusByMember(key, userID, distance)
	return
}

func getUsers(ids []string, userID string) (users []User, err error) {
	var i int
	for _, id := range ids {
		if id == userID {
			continue
		}
		i, err = strconv.Atoi(id)
		if err != nil {
			return
		}

		// TODO 必要カラムだけselect
		rawUser := &orm.User{}
		// 位置情報だけ残ってるけどユーザーが削除されている可能性があるから
		if err = rawUser.Find(i); err != nil {
			continue
		}
		// TODO 表示項目決まり次第ここにマッピングする
		user := User{
			Name: rawUser.Name,
		}
		users = append(users, user)
	}
	return
}
