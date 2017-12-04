package orm

import "github.com/jinzhu/gorm"

type ProductUser struct {
	gorm.Model
	ProductId uint `gorm:"type:bigint;not null"`
	UserId    uint `gorm:"type:bigint;not null"`
	DispOrder uint `gorm:"type:bigint;not null"`
}
