package detail

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func UpdateUser(r io.ReadCloser) error {
	user, err := decode(r)
	log.Printf("ユーザーは%#v", user)
	if err != nil {
		return nil
	}
	return nil
}

func decode(r io.ReadCloser) (*User, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	user := &User{}
	if err := json.Unmarshal(bytes, &user); err != nil {
		return nil, err
	}
	return user, nil
}
