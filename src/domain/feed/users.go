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
func GetResponse(limit int, offset int) *Response {
	rawUsers := orm.Users{}
	rawUsers.LimitOffset(limit, offset)
	lastID := rawUsers[len(rawUsers)-1].Model.ID
	users := Users{}
	for _, u := range rawUsers {
		user := User{
			ID:        u.Model.ID,
			Name:      u.Name,
			AvatarURL: u.AvatarUrl,
			Skills:    []string{},
			Overview:  u.Overview,
		}
		users = append(users, user)
	}
	return &Response{
		HasNext: nextExist(int(lastID)),
		Users:   users,
	}
}

func nextExist(id int) bool {
	user := orm.User{}
	user.Find(id)
	return user.Model.ID != 0
}
