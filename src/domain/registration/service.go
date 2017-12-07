package registration

import (
	"genepse_api/src/infra/orm"
)

type Login struct {
	LoginURL string `json:"login_url"`
}

type Callback struct {
	UserID uint `json:"user_id"`
}

// TODO レコードごととらずに、存在確認のみする
// TODO プロバイダー毎に処理書かずに抽象化したい
func Registered(provider string, accountID string) bool {
	switch provider {
	case "facebook":
		exists, _ := orm.ExistsFacebookBy("AccountId", accountID)
		return exists
	default:
		// TODO エラーハンドリングする
		return false
	}
}

// Register はuserを登録します
func Register(userName string, avatarURL string, accountID string, provider string) (userID uint, err error) {
	var facebookID uint
	// TODO 抽象化
	switch provider {
	case "facebook":
		f := &orm.FacebookAccount{
			AccountId: accountID,
		}
		err = f.Insert()
		if err != nil {
			return
		}
		facebookID = f.Model.ID
	}
	// TODO 画像をcloudStorageに入れて、AvatarUrlにそのurl入れる
	u := &orm.User{
		Name:              userName,
		AvatarUrl:         avatarURL,
		FacebookAccountId: facebookID,
	}
	err = u.Insert()
	if err != nil {
		return
	}
	userID = u.Model.ID
	return
}

func UserID(provider orm.Provider, accountID string) (userID uint, err error) {
	err = provider.FindBy("AccountId", accountID)
	if err != nil {
		return
	}
	user := orm.User{}
	err = user.FindByProvider(provider, provider.GetID())
	if err != nil {
		return
	}
	userID = user.Model.ID
	return
}
