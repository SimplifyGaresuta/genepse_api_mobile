package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type FacebookAccount struct {
	gorm.Model
	AccountId string `gorm:"size:100;unique"`
	MypageUrl string `gorm:"size:300;not null"`
}

func (f *FacebookAccount) Insert() (err error) {
	err = db.Create(f).Error
	return
}

func (f *FacebookAccount) Find(id int) (err error) {
	return db.First(f, id).Error
}

func (f *FacebookAccount) FindBy(column string, value interface{}) error {
	switch column {
	case "AccountId":
		if v, ok := value.(string); ok {
			db.Where("account_id = ?", v).First(f)
			return nil
		} else {
			return errors.New("AccountIdにはstring型の値を渡して下さい。")
		}
	default:
		return errors.New("カラム名が違います。")
	}
}

func (f *FacebookAccount) GetID() uint {
	return f.Model.ID
}

func (f *FacebookAccount) ProviderName() string {
	return "facebook"
}

// TODO メソッドにする
func ExistsFacebookBy(column string, value interface{}) (bool, error) {
	facebook := FacebookAccount{}
	switch column {
	case "AccountId":
		if v, ok := value.(string); ok {
			db.Where("account_id = ?", v).First(&facebook)
			if facebook.Model.ID == 0 {
				return false, nil
			} else {
				return true, nil
			}
		} else {
			return false, errors.New("AccountIdにはstring型の値を渡して下さい。")
		}
	default:
		return false, errors.New("カラム名が違います。")
	}
}
