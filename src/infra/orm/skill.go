package orm

import "github.com/jinzhu/gorm"

type Skill struct {
	gorm.Model
	Name      string `gorm:"size:20;not null"`
	DeleteFlg int    `gorm:"type:tinyint;default:0;not null"`
}
