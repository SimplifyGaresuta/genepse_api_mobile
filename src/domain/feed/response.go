package feed

import (
	"errors"
	"genepse_api/src/domain"
	"genepse_api/src/infra/orm"
	"log"
)

// Response is フィードで返すjsonオブジェクト
type Response struct {
	HasNext bool  `json:"has_next"`
	Users   Users `json:"users"`
}

// スキルの表示数
const (
	numberOfSkills uint = 3
	query               = `
	select distinct u.id, u.name, u.avatar_url, u.attribute_id, u.overview
  from users as u left join skill_users as s on u.id=s.user_id
  where u.attribute_id is not null or u.overview is not null or s.user_id is not null
	limit ? offset ?;
`
)

// GetResponse return response
func GetResponse(limit, offset int) (response *Response, err error) {
	rawUsers := orm.Users{}
	// TODO 必要なカラムだけselectする
	if err = rawUsers.RawQuery(query, limit, offset); err != nil {
		log.Println("heyheyhey")
		return
	}
	if len(rawUsers) < 1 {
		err = errors.New("指定条件のユーザーは存在しません。")
		return
	}

	users := Users{}
	for _, u := range rawUsers {
		skills, err := skillsOfUser(skillsTerms{UserID: u.Model.ID, Limit: numberOfSkills})
		if err != nil {
			break
		}
		user := User{
			ID:        u.Model.ID,
			Name:      u.Name,
			AvatarURL: u.AvatarUrl,
			Attribute: domain.GetAttribute(u.AttributeId),
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
