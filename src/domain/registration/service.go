package registration

import (
	"context"
	"genepse_api/src/infra/objstorage"
	"genepse_api/src/infra/orm"
	"io"
	"net/http"
)

type Login struct {
	LoginURL string `json:"login_url"`
}

type Callback struct {
	UserID uint `json:"user_id"`
}

// InputItems is 登録時に必要な項目
type RequiredItems struct {
	UserName  string
	AvatarURL string
	Provider  orm.Provider
	Ctx       context.Context
}

func Registered(provider orm.Provider) (bool, error) {
	return provider.Exists(provider.GetAccountID())
}

// Register はuserを登録します
func Register(items RequiredItems) (userID uint, err error) {
	items.AvatarURL = items.Provider.NewAvatarURL()

	r, err := downloadImage(items.AvatarURL)
	if err != nil {
		return
	}
	if items.AvatarURL, err = uploadImage(r, items.Ctx); err != nil {
		return
	}

	u := &orm.User{
		Name:      items.UserName,
		AvatarUrl: items.AvatarURL,
	}
	if err = u.Insert(); err != nil {
		return
	}
	userID = u.Model.ID

	items.Provider.SetMyPageURL()
	items.Provider.SetUserID(userID)
	err = items.Provider.Insert()
	return
}

func downloadImage(u string) (io.ReadCloser, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func uploadImage(r io.ReadCloser, ctx context.Context) (imageURL string, err error) {
	imageURL, err = objstorage.Upload(ctx, r, objstorage.ProfileDir)
	if err != nil {
		return
	}
	return
}

func UserID(provider orm.Provider) (userID uint, err error) {
	err = provider.Find(provider.GetAccountID())
	if err != nil {
		return
	}
	user := orm.User{}
	err = user.FindByProvider(provider)
	if err != nil {
		return
	}
	userID = user.Model.ID
	return
}
