package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type FacebookAccount struct {
	gorm.Model
	AccountId string `gorm:"size:100";unique`
	MypageUrl string `gorm:"size:300"`
	DeleteFlg int    `gorm:"type:tinyint;default:0;not null"`
}

func FindFacebookBy(column string, value interface{}) (*FacebookAccount, error) {
	facebook := FacebookAccount{}
	switch column {
	case "AccountId":
		if v, ok := value.(string); ok {
			db.Where("account_id = ?", v).First(&facebook)
			return &facebook, nil
		} else {
			return nil, errors.New("AccountIdにはstring型の値を渡して下さい。")
		}
	default:
		return nil, errors.New("カラム名が違います。")
	}
}
