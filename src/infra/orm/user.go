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
	AttributeId       uint   `gorm:"type:smallint;default:1;not null"`
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

func (u *User) Insert() (err error) {
	err = db.Create(u).Error
	return
}

// FindUser find user matching the given id
func (u *User) Find(id int) (err error) {
	err = db.First(u, id).Error
	return
}

func (u *User) FindBy(column string, value interface{}) error {
	switch column {
	case "FacebookAccountId":
		if v, ok := value.(uint); ok {
			db.Where("facebook_account_id = ?", v).First(u)
			return nil
		} else {
			return errors.New("FacebookAccountIdにはuint型の値を渡して下さい。")
		}
	default:
		return errors.New("カラム名が違います。")
	}
}

func (u *User) FindByProvider(provider Provider, id uint) error {
	providerName := provider.ProviderName()
	return db.Where(providerName+"_account_id = ?", id).First(u).Error
}

type Users []User

func (u *Users) LimitOffset(limit int, offset int) (err error) {
	return db.Limit(limit).Offset(offset).Find(u).Error
}
