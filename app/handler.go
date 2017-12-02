package main

import (
	"encoding/json"
	"fmt"
	"genepse_api/src/infra/orm"
	"genepse_api/src/middleware"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

func userUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func userDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// TODO gomniauth使用はmiddlewareに任せる
func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	providerName := ps.ByName("provider")
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
		return
	}

	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s: %s", provider, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := json.Marshal(middleware.Login{
		LoginURL: loginURL,
	})
	if err != nil {
		log.Println(err)
	}
	w.Write(res)
}

func callback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	providerName := ps.ByName("provider")
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
		return
	}

	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s: %s", provider, err), http.StatusInternalServerError)
		return
	}

	// プロバイダーからユーザー情報を取得
	user, err := provider.GetUser(creds)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when trying to get user from %s: %s", provider, err), http.StatusInternalServerError)
		return
	}
	accountID := user.IDForProvider(providerName)
	var userID uint
	if !registered(providerName, accountID) {
		// TODO 抽象化
		switch providerName {
		case "facebook":
			f := &orm.FacebookAccount{
				AccountId: accountID,
			}
			f.Insert()
		default:
			return
		}
		// TODO 画像をcloudStorageに入れて、AvatarUrlにそのurl入れる
		u := &orm.User{
			Name:      user.Name(),
			AvatarUrl: user.AvatarURL(),
		}
		u.Insert()
		userID = u.Model.ID
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := json.Marshal(middleware.Callback{
		UserID: userID,
	})
	if err != nil {
		log.Println(err)
	}
	w.Write(res)
}

// TODO レコード毎とらずに、存在確認のみする
// TODO プロバイダー毎に処理書かずに抽象化したい
func registered(provider string, id string) bool {
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
