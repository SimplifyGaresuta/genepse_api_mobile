package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	UserId       uint
	Title        string `gorm:"size:20;not null"`
	ReferenceUrl string `gorm:"size:300"`
	ImageUrl     string `gorm:"size:300"`
	DispOrder    uint   `gorm:"type:bigint"`
}

func (p *Product) Insert() error {
	return db.Create(p).Error
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

func (p *Product) Update(id uint) error {
	before := Product{}
	before.ID = id
	return db.Debug().Model(&before).Updates(p).Error
}

type Products []Product

func (p *Products) FindByUser(id uint) error {
	user := &User{}
	user.Model.ID = id
	return db.Model(user).Related(p).Error
}
