package orm

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"size:20;not null"`
	AvatarUrl string `gorm:"size:300"`
}

func FindUser(id int) *User {
	user := User{}
	db.First(&user, id)
	return &user
}
