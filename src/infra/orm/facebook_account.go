package orm

import (
	"github.com/jinzhu/gorm"
)

type FacebookAccount struct {
	gorm.Model
	AccountId string `gorm:"size:100"`
	MypageUrl string `gorm:"size:300"`
	DeleteFlg int    `gorm:"type:tinyint;default:0;not null"`
}
