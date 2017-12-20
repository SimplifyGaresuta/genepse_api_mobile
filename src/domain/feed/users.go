package feed

import (
	"genepse_api/src/infra/orm"
)

type Users []User

// User is 画面で表示するユーザー情報
type User struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	AvatarURL    string   `json:"avatar_url"`
	Attribute    string   `json:"attribute"`
	Skills       []string `json:"skills"`
	Overview     string   `json:"overview"`
	ActivityBase string   `json:"activity_base"`
}

// nextExist return 与えられたidの次にレコードがあるか
func nextExist(id int) bool {
	user := orm.User{}
	user.Find(id + 1)
	return user.Model.ID != 0
}

// skillsTerms is スキル検索時の条件
type skillsTerms struct {
	UserID uint
	Limit  uint
}

// 与えられたユーザーのスキル名を返す
func skillsOfUser(terms skillsTerms) (skillNames []string, err error) {
	skillUsers := &orm.SkillUsers{}
	if terms.Limit == 0 {
		if err = skillUsers.Where("user_id = ?", terms.UserID); err != nil {
			return
		}
	} else {
		if err = skillUsers.WhereLimit("user_id = ? and disp_order <= ?", 3, terms.UserID, terms.Limit); err != nil {
			return
		}
	}
	for _, s := range *skillUsers {
		skill := &orm.Skill{}
		if err := skill.Find(int(s.SkillId)); err != nil {
			break
		}
		skillNames = append(skillNames, skill.Name)
	}
	return
}
