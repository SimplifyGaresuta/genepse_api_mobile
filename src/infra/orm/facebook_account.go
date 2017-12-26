package orm

import (
	"fmt"
)

type FacebookAccount struct {
	AccountId string `gorm:"primary_key; size:100; not null; default:''"`
	UserId    uint   `gorm:"type:bigint;not null"`
	MypageUrl string `gorm:"size:300;not null"`
}

func (f *FacebookAccount) Insert() error {
	return db.Create(f).Error
}

func (f *FacebookAccount) Find(accountID string) error {
	return db.Where("account_id = ?", accountID).First(f).Error
}

func (f *FacebookAccount) GetAccountID() string {
	return f.AccountId
}

func (f *FacebookAccount) GetMypageURL() string {
	return f.MypageUrl
}

func (f *FacebookAccount) ProviderName() string {
	return "facebook"
}

func (f *FacebookAccount) Exists(accountID string) bool {
	result := struct{ IsExists bool }{}
	db.Raw("SELECT EXISTS(SELECT * FROM facebook_accounts WHERE account_id=?) as is_exists;", accountID).Scan(&result)
	return result.IsExists
}

func (f *FacebookAccount) NewAvatarURL() string {
	return fmt.Sprintf("https://graph.facebook.com/%s/picture?width=9999", f.AccountId)
}

func (f *FacebookAccount) SetMyPageURL() {
	f.MypageUrl = "https://www.facebook.com/" + f.AccountId
	return
}

func (f *FacebookAccount) SetUserID(id uint) {
	f.UserId = id
	return
}
