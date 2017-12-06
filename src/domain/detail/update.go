package detail

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func UpdateUser(r io.ReadCloser) error {
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
