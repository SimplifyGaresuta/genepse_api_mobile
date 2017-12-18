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
	CoverUrl          string `gorm:"size:300"`
	AttributeId       int    `gorm:"type:smallint"`
	Overview          string `gorm:"size:500"`
	Awards            []Award
	Licenses          []License
	Products          []Product
	Gender            int    `gorm:"type:tinyint;not null"`
	Age               int    `gorm:"type:smallint"`
	Address           string `gorm:"size:100"`
	SchoolCarrer      string `gorm:"size:500"`
	ActivityBase      string `gorm:"size:100"`
	FacebookAccountId uint   `gorm:"type:bigint"`
	TwitterAccountId  uint   `gorm:"type:bigint"`
}

func (u *User) Insert() error {
	return db.Create(u).Error
}

// Find find user matching the given id
func (u *User) Find(id int) (err error) {
	return db.First(u, id).Error
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

func (u *User) Update(id uint) error {
	before := User{}
	before.ID = id
	return db.Debug().Model(&before).Updates(u).Error
}

// TODO 上手く抽象化。各providerテーブルにユーザーidもたせるかも
func (u *User) ProviderURL(p Provider) (providerURL string, err error) {
	switch p.ProviderName() {
	case "facebook":
		i := u.FacebookAccountId
		if i == 0 {
			err = errors.New("facebookが登録されていません")
			return
		}
		if err = p.Find(int(i)); err != nil {
			return
		}
		providerURL = p.GetMypageURL()
		return
	case "twitter":
		i := u.TwitterAccountId
		if i == 0 {
			err = errors.New("twitterが登録されていません")
			return
		}
		if err = p.Find(int(i)); err != nil {
			return
		}
		providerURL = p.GetMypageURL()
		return
	default:
		return
	}
}

func (u *User) RawQuery(query string, args ...interface{}) error {
	return db.Raw(query, args...).Scan(u).Error
}

type Users []User

// TODO この辺gormの仕様読んで、各クエリメソッドを部品化
func (u *Users) LimitOffset(limit int, offset int) (err error) {
	return db.Limit(limit).Offset(offset).Find(u).Error
}

/*
func (u *Users) LimitOffsetWhere(limit, offset int, query string, args ...interface{}) error {
	return db.Where(query, args...).Limit(limit).Offset(offset).Find(u).Error
}
*/

func (u *Users) RawQuery(query string, args ...interface{}) error {
	return db.Raw(query, args...).Scan(u).Error
}
