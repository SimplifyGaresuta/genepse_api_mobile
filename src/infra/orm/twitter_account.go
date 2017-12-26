package orm

type TwitterAccount struct {
	AccountId string `gorm:"primary_key; size:100; not null; default:''"`
	UserId    uint   `gorm:"type:bigint;not null"`
	MypageUrl string `gorm:"size:300;not null"`
}

func (t *TwitterAccount) Insert() (err error) {
	err = db.Create(t).Error
	return
}

func (t *TwitterAccount) Find(accountID string) (err error) {
	return db.Where("account_id = ?", accountID).First(t).Error
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

func (t *TwitterAccount) Exists(accountID string) bool {
	result := struct{ IsExists bool }{}
	db.Raw("SELECT EXISTS(SELECT * FROM twitter_accounts WHERE account_id=?) as is_exists;", accountID).Scan(&result)
	return result.IsExists
}

// TODO 実装する
func (t *TwitterAccount) NewAvatarURL() string {
	return ""
}

func (t *TwitterAccount) SetMyPageURL() {
	t.MypageUrl = "https://twitter.com/" + t.AccountId
	return
}

func (t *TwitterAccount) SetUserID(id uint) {
	t.UserId = id
	return
}
