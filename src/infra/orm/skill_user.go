package orm

import "github.com/jinzhu/gorm"

// TODO アソシエーションしっかり
type SkillUser struct {
	gorm.Model
	SkillId   uint `gorm:"type:bigint;not null"`
	UserId    uint `gorm:"type:bigint;not null"`
	DispOrder uint `gorm:"type:bigint;not null"`
	DeleteFlg int  `gorm:"type:tinyint;default:0;not null"`
}
