// TODO 全体的にリファクタリングする
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
	// TODO どれか一つ実行するようにする---------------------------------------
	rawUser, err := mappingUser(id, user)
	if err != nil {
		return err
	}
	if err := rawUser.Update(id); err != nil {
		return err
	}
	if err := updateAward(id, user.Awards); err != nil {
		return err
	}
	if err := updateLicense(id, user.Licenses); err != nil {
		return err
	}
	// アソシエーションしたら消す
	if err := updateSkills(id, user.Skills); err != nil {
		return err
	}
	// ---------------------------------------------------------
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

func mappingUser(id uint, user *User) (rawUser *orm.User, err error) {
	// TODO snsを更新した時にFacebookAccountIdも更新させる
	rawUser = &orm.User{
		Name:         user.Name,
		AvatarUrl:    user.AvatarURL,
		Overview:     user.Overview,
		Age:          user.Age,
		Address:      user.Address,
		SchoolCarrer: user.SchoolCareer,
		ActivityBase: user.ActivityBase,
	}
	if user.Attribute != "" {
		rawUser.AttributeId = domain.GetAttributeID(user.Attribute)
	}
	if user.Gender != "" {
		rawUser.Gender = domain.GetGenderID(user.Gender)
	}
	return
}

// TODO 空配列だった場合全部消す
func updateAward(userID uint, awardNames []string) (err error) {
	if len(awardNames) >= 1 {
		awards := orm.Awards{}
		if err = awards.BatchDelete("user_id = ?", userID); err != nil {
			return
		}
		// TODO 消したい場合は空文字を入れる実装だが、変える
		if awardNames[0] == "" {
			return
		}

		for _, awardName := range awardNames {
			award := &orm.Award{UserId: userID, Name: awardName}
			if err = award.Insert(); err != nil {
				return
			}
		}
	}
	return
}

// TODO 空配列だった場合全部消す
func updateLicense(userID uint, licenseNames []string) (err error) {
	if len(licenseNames) >= 1 {
		licenses := orm.Licenses{}
		if err = licenses.BatchDelete("user_id = ?", userID); err != nil {
			return
		}
		// TODO 消したい場合は空文字を入れる実装だが、変える
		if licenseNames[0] == "" {
			return
		}

		for _, name := range licenseNames {
			license := &orm.License{UserId: userID, Name: name}
			if err = license.Insert(); err != nil {
				return
			}
		}
	}
	return
}

// TODO 空配列だった場合全部消す
func updateSkills(userID uint, skillNames []string) (err error) {
	if len(skillNames) >= 1 {
		skillUsers := orm.SkillUsers{}
		if err = skillUsers.BatchDelete("user_id = ?", userID); err != nil {
			return
		}
		// TODO 消したい場合は空文字を入れる実装だが、変える
		if skillNames[0] == "" {
			return
		}

		dispOrder := 1
		for _, skillName := range skillNames {
			skill := orm.Skill{}
			if err = skill.FindBy("Name", skillName); err != nil {
				return
			}
			skillUser := orm.SkillUser{SkillId: skill.Model.ID, UserId: userID, DispOrder: uint(dispOrder)}
			if err = skillUser.Insert(); err != nil {
				return
			}
			dispOrder++
		}
	}
	return
}
