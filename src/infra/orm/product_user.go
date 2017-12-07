package orm

import "github.com/jinzhu/gorm"

type ProductUser struct {
	gorm.Model
	ProductId uint `gorm:"type:bigint;not null"`
	UserId    uint `gorm:"type:bigint;not null"`
	DispOrder uint `gorm:"type:bigint;not null"`
}

type ProductUsers []ProductUser

func (p *ProductUsers) Where(query string, args ...interface{}) (err error) {
	return db.Where(query, args...).Find(p).Error
}

func (_ *ProductUsers) BatchDelete(query string, args ...interface{}) error {
	return db.Where(query, args...).Delete(&ProductUser{}).Error
}
