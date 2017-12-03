// TODO アソシエーションしっかり
package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// TODO アソシエーションしっかり
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

func (u *User) Insert() error {
	d := db.Create(u)
	err := d.Error
	return err
}

// FindUser find user matching the given id
func (u *User) Find(id int) {
	db.First(u, id)
	return
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

type Users []User

func (u *Users) LimitOffset(limit int, offset int) {
	db.Limit(limit).Offset(offset).Find(u)
	return
}

/*
func WhereUser(query string, args ...interface{}) {
	users := []User{}
	db.Where(query, args[0]).Find(&users)
}*/
