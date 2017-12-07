package detail

import (
	"encoding/json"
	"genepse_api/src/domain"
	"genepse_api/src/infra/orm"
	"io"
	"io/ioutil"
	"log"
)

func UpdateUser(id int, r io.ReadCloser) error {
	user, err := decode(r)
	if err != nil {
		return err
	}
	log.Printf("ユーザーは%#v", user)
	rawUser, err := doMapping(user)
	if err != nil {
		return err
	}
	if err := update(rawUser); err != nil {
		return err
	}
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

func doMapping(user *User) (*orm.User, error) {
	// TODO 埋め込んでるとこちゃんとやる
	// TODO snsを更新した時にFacebookAccountIdも更新させる
	rawUser := &orm.User{
		Name:         user.Name,
		AvatarUrl:    user.AvatarURL,
		AttributeId:  domain.GetAttributeID(user.Attribute),
		Overview:     user.Overview,
		Awards:       "ジロッカソン優勝, SPAJAM優勝",
		License:      "TOEIC 880点",
		Gender:       domain.GetGenderID(user.Gender),
		Age:          user.Age,
		Address:      user.Address,
		SchoolCarrer: user.SchoolCareer,
	}
	return rawUser, nil
}
func update(user *orm.User) error {
	return nil
}
