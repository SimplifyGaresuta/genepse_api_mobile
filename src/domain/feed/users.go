package feed

import (
	"genepse_api/src/infra/orm"
	"log"
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
	q := `select distinct u.id, u.name, u.avatar_url, u.attribute_id, u.overview, u.activity_base
from users as u left join skill_users as s on u.id=s.user_id
where u.id > ? and u.attribute_id != 0 and u.overview != "" and s.user_id is not null
limit 1;`
	user := orm.User{}
	if err := user.RawQuery(q, id); err != nil {
		log.Println(err)
		return false
	}
	log.Printf("ユーザーは%#v", user)
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
