// TODO アソシエーションしっかり
package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name              string `gorm:"size:20;not null"`
	AvatarUrl         string `gorm:"size:300"`
	AttributeId       uint   `gorm:"type:smallint;not null"`
	Overview          string `gorm:"size:500"`
	Awards            string `gorm:"size:500"`
	License           string `gorm:"size:500"`
	Gender            int    `gorm:"type:tinyint"`
	Age               int    `gorm:"type:smallint"`
	Address           string `gorm:"size:100"`
	SchoolCarrer      string `gorm:"size:500"`
	FacebookAccountId uint   `gorm:"type:bigint"`
	DeleteFlg         int    `gorm:"type:tinyint;default:0;not null"`
}

// FindUser find user matching the given id
func FindUser(id int) *User {
	user := User{}
	db.First(&user, id)
	return &user
}

func (u *User) Insert() error {
	d := db.Create(u)
	err := d.Error
	return err
}

func FindUserBy(column string, value interface{}) (*User, error) {
	user := User{}
	switch column {
	case "FacebookAccountId":
		if v, ok := value.(uint); ok {
			db.Where("facebook_account_id = ?", v).First(&user)
			return &user, nil
		} else {
			return nil, errors.New("FacebookAccountIdにはuint型の値を渡して下さい。")
		}
	default:
		return nil, errors.New("カラム名が違います。")
	}
}
