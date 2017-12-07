package detail

import (
	"encoding/json"
	"genepse_api/src/domain"
	"genepse_api/src/infra/orm"
	"io"
	"io/ioutil"
	"log"
)

func UpdateUser(id uint, r io.ReadCloser) error {
	user, err := decode(r)
	if err != nil {
		return err
	}
	log.Printf("ユーザーは%#v", user)
	rawUser, err := mappingUser(user)
	if err != nil {
		return err
	}
	if err := rawUser.Update(id); err != nil {
		return err
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

func mappingUser(user *User) (*orm.User, error) {
	// TODO 埋め込んでるとこちゃんとやる
	// TODO skillsとproductsも更新
	// TODO snsを更新した時にFacebookAccountIdも更新させる
	rawUser := &orm.User{
		Name:         user.Name,
		AvatarUrl:    user.AvatarURL,
		Overview:     user.Overview,
		Awards:       "ジロッカソン優勝, SPAJAM優勝",
		License:      "TOEIC 880点",
		Age:          user.Age,
		Address:      user.Address,
		SchoolCarrer: user.SchoolCareer,
	}
	if user.Attribute != "" {
		rawUser.AttributeId = domain.GetAttributeID(user.Attribute)
	}
	if user.Gender != "" {
		rawUser.Gender = domain.GetGenderID(user.Gender)
	}

	return rawUser, nil
}
func update(user *orm.User) error {
	return nil
}
