package orm

import "github.com/jinzhu/gorm"

type Award struct {
	gorm.Model
	UserId uint   `gorm:"type:bigint;not null"`
	Name   string `gorm:"size:20;not null"`
	Year   int    `gorm:"type:smallint"`
}

func (a *Award) Insert() error {
	return db.Create(a).Error
}

type Awards []Award

func (a *Awards) FindByUser(userID int) (err error) {
	user := &User{}
	if err = user.Find(userID); err != nil {
		return
	}
	db.Model(user).Related(a)
	return
}

func (_ *Awards) BatchDelete(query string, args ...interface{}) error {
	return db.Where(query, args...).Delete(&Award{}).Error
}
