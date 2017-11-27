package infra

type User struct {
	name       string
	avatarURL  string
	facebookID string
}

func (u *User) Find(id int) *User {
	return &User{}
}
