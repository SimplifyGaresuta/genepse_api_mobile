package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title        string `gorm:"size:20;not null"`
	ReferenceUrl string `gorm:"size:300"`
	ImageUrl     string `gorm:"size:300"`
}

func (p *Product) Find(id int) (err error) {
	return db.First(p, id).Error
}

func (p *Product) FindBy(column string, value interface{}) error {
	switch column {
	case "Title":
		if v, ok := value.(uint); ok {
			if err := db.Where("user_id = ?", v).First(p).Error; err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("Titleにはstring型の値を渡して下さい。")
		}
	default:
		return errors.New("カラム名が違います。")
	}
}
