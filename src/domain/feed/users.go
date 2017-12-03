package feed

import "genepse_api/src/infra/orm"

type Response struct {
	HasNext bool  `json:"has_next"`
	Users   Users `json:"users"`
}

type Users []User

type User struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	AvatarURL string   `json:"avatar_url"`
	Skills    []string `json:"skills"`
	Overview  string   `json:"overview"`
}

// GetResponse return response
func GetResponse(limit int, offset int) (response *Response, err error) {
	rawUsers := orm.Users{}
	if err = rawUsers.LimitOffset(limit, offset); err != nil {
		return
	}

	users := Users{}
	for _, u := range rawUsers {
		skills, err := skillsOfUser(u.Model.ID)
		if err != nil {
			break
		}
		user := User{
			ID:        u.Model.ID,
			Name:      u.Name,
			AvatarURL: u.AvatarUrl,
			Skills:    skills,
			Overview:  u.Overview,
		}
		users = append(users, user)
	}

	lastID := rawUsers[len(rawUsers)-1].Model.ID
	response = &Response{
		HasNext: nextExist(int(lastID)),
		Users:   users,
	}
	return
}

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
