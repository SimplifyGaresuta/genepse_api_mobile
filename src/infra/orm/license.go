package orm

import "github.com/jinzhu/gorm"

// TODO Awardと抽象化
type License struct {
	gorm.Model
	UserId uint   `gorm:"type:bigint;not null"`
	Name   string `gorm:"size:20;not null"`
	Year   int    `gorm:"type:smallint"`
}

func (l *License) Insert() error {
	return db.Create(l).Error
}

type Licenses []License

func (a *Licenses) FindByUser(userID int) (err error) {
	user := &User{}
	if err = user.Find(userID); err != nil {
		return
	}
	db.Model(user).Related(a)
	return
}

func (_ *Licenses) BatchDelete(query string, args ...interface{}) error {
	return db.Where(query, args...).Delete(&License{}).Error
}
