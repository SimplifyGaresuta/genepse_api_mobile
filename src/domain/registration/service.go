package registration

import "genepse_api/src/infra/orm"

// TODO レコード毎とらずに、存在確認のみする
// TODO プロバイダー毎に処理書かずに抽象化したい
func Registered(provider string, id string) bool {
	switch provider {
	case "facebook":
		account, _ := orm.FindFacebookBy("AccountId", id)
		if account == nil {
			return false
		} else {
			return true
		}
	default:
		return false
	}
}

// Register register user
func Register(userName string, avatarURL string, accountID string, provider string) (userID uint) {
	// TODO 抽象化
	switch provider {
	case "facebook":
		f := &orm.FacebookAccount{
			AccountId: accountID,
		}
		f.Insert()
	}
	// TODO 画像をcloudStorageに入れて、AvatarUrlにそのurl入れる
	u := &orm.User{
		Name:      userName,
		AvatarUrl: avatarURL,
	}
	u.Insert()
	userID = u.Model.ID
	return
}
