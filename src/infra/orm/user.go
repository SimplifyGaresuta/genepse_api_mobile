package orm

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name            string `gorm:"size:20;not null"`
	AvatarUrl       string `gorm:"size:300"`
	CoverUrl        string `gorm:"size:300"`
	AttributeId     int    `gorm:"type:smallint"`
	Overview        string `gorm:"size:500"`
	Awards          []Award
	Licenses        []License
	Products        []Product
	Gender          int    `gorm:"type:tinyint;not null"`
	Age             int    `gorm:"type:smallint"`
	Address         string `gorm:"size:100"`
	SchoolCarrer    string `gorm:"size:500"`
	ActivityBase    string `gorm:"size:100"`
	FacebookAccount FacebookAccount
	TwitterAccount  TwitterAccount
}

func (u *User) Insert() error {
	return db.Create(u).Error
}

// Find find user matching the given id
func (u *User) Find(id int) (err error) {
	return db.First(u, id).Error
}

func (u *User) FindByProvider(provider Provider) error {
	return db.Debug().Model(&provider).Related(u).Error
}

func (u *User) Update(id uint) error {
	before := User{}
	before.ID = id
	return db.Debug().Model(&before).Updates(u).Error
}

func (u *User) ProviderURL(p Provider) (providerURL string, err error) {
	err = db.Debug().Model(u).Related(p).Select("mypage_url").Error
	return p.GetMypageURL(), err
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
