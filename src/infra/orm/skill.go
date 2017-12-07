package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Skill struct {
	gorm.Model
	Name string `gorm:"size:20;not null"`
}

// Find find a skill matching the give
func (s *Skill) Find(id int) (err error) {
	err = db.First(s, id).Error
	return
}

func (s *Skill) FindBy(column string, value interface{}) error {
	switch column {
	case "Name":
		if v, ok := value.(string); ok {
			if err := db.Where("name = ?", v).First(s).Error; err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("Nameにはstring型の値を渡して下さい。")
		}
	default:
		return errors.New("カラム名が違います。")
	}
}

func (s *Skill) Insert() (err error) {
	err = db.Create(s).Error
	return
}
