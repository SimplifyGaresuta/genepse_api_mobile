package feed

import "genepse_api/src/infra/orm"

// Response is フィードで返すjsonオブジェクト
type Response struct {
	HasNext bool  `json:"has_next"`
	Users   Users `json:"users"`
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
