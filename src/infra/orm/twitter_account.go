package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TwitterAccount struct {
	gorm.Model
	AccountId string `gorm:"size:100;unique"`
	MypageUrl string `gorm:"size:300;not null"`
}

func (t *TwitterAccount) Insert() (err error) {
	err = db.Create(t).Error
	return
}

func (t *TwitterAccount) Find(id int) (err error) {
	return db.First(t, id).Error
}

func (t *TwitterAccount) FindBy(column string, value interface{}) error {
	switch column {
	case "AccountId":
		if v, ok := value.(string); ok {
			db.Where("account_id = ?", v).First(t)
			return nil
		} else {
			return errors.New("AccountIdにはstring型の値を渡して下さい。")
		}
	default:
		return errors.New("カラム名が違います。")
	}
}

func (t *TwitterAccount) GetID() uint {
	return t.Model.ID
}

func (t *TwitterAccount) GetAccountID() string {
	return t.AccountId
}

func (t *TwitterAccount) GetMypageURL() string {
	return t.MypageUrl
}

func (t *TwitterAccount) ProviderName() string {
	return "twitter"
}
