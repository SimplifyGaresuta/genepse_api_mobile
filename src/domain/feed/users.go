package feed

import "genepse_api/src/infra/orm"

type Users []User

// Uses is 画面で表示するユーザー情報
type User struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	AvatarURL string   `json:"avatar_url"`
	Skills    []string `json:"skills"`
	Overview  string   `json:"overview"`
}

// nextExist return 与えられたidの次にレコードがあるか
func nextExist(id int) bool {
	user := orm.User{}
	user.Find(id)
	return user.Model.ID != 0
}

// 与えられたユーザーのスキル名を全て返す
func skillsOfUser(userID uint) (skillNames []string, err error) {
	skillUsers := &orm.SkillUsers{}
	if err = skillUsers.Where("user_id = ?", userID); err != nil {
		return
	}
	skill := &orm.Skill{}
	for _, s := range *skillUsers {
		if err := skill.Find(int(s.SkillId)); err != nil {
			break
		}
		skillNames = append(skillNames, skill.Name)
	}
	return
}
