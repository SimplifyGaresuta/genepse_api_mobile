// TODO アソシエーションしっかり
package orm

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name              string `gorm:"size:20;not null"`
	AvatarUrl         string `gorm:"size:300"`
	AttributeId       int    `gorm:"type:smallint;not null"`
	Overview          string `gorm:"size:500"`
	Awards            string `gorm:"size:500"`
	License           string `gorm:"size:500"`
	Gender            int    `gorm:"type:tinyint"`
	Age               int    `gorm:"type:smallint"`
	Address           string `gorm:"size:100"`
	SchoolCarrer      string `gorm:"size:500"`
	FacebookAccountId string `gorm:"size:100"`
	DeleteFlg         int    `gorm:"type:tinyint;default:0;not null"`
}

// FindUser find user matching the given id
func FindUser(id int) *User {
	user := User{}
	db.First(&user, id)
	return &user
}

func (u *User) Insert() {
	db.Create(u)
}
