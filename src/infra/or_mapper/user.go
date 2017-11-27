package or_mapper

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	name      string
	avatarURL string
}

func FindUser(id int) *User {
	user := User{}
	db.First(&user, id)
	return &user
}
