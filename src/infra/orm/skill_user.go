package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// TODO アソシエーションしっかり
type SkillUser struct {
	gorm.Model
	SkillId   uint `gorm:"type:bigint;not null"`
	UserId    uint `gorm:"type:bigint;not null"`
	DispOrder uint `gorm:"type:bigint;not null"`
	DeleteFlg int  `gorm:"type:tinyint;default:0;not null"`
}

func (s *SkillUser) FindBy(column string, value interface{}) error {
	switch column {
	case "UserId":
		if v, ok := value.(uint); ok {
			if err := db.Where("user_id = ?", v).First(s).Error; err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("UserIdにはuint型の値を渡して下さい。")
		}
	default:
		return errors.New("カラム名が違います。")
	}
}

type SkillUsers []SkillUser

func (s *SkillUsers) Where(query string, args ...interface{}) (err error) {
	err = db.Debug().Where(query, args...).Find(s).Error
	return
}

func (s *SkillUsers) WhereLimit(query string, limit int, args ...interface{}) (err error) {
	err = db.Debug().Where(query, args...).Find(s).Limit(limit).Error
	return
}
