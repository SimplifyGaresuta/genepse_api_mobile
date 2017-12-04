package orm

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title        string `gorm:"size:20;not null"`
	ReferenceUrl string `gorm:"size:300"`
	ImageUrl     string `gorm:"size:300"`
	DeleteFlg    int    `gorm:"type:tinyint;default:0;not null"`
}
