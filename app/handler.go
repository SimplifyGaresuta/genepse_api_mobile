package main

import (
	"encoding/json"
	"fmt"
	"genepse_api/src/domain/feed"
	"genepse_api/src/domain/registration"
	"genepse_api/src/infra/orm"
	"genepse_api/src/middleware"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
)

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	limit, err := strconv.Atoi(query["limit"][0])
	offset, err := strconv.Atoi(query["offset"][0])
	if err != nil {
		log.Println("クエリパラメータが不正です", err)
		// TODO 異常系のjson
		return
	}
	response, err := feed.GetResponse(limit, offset)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := json.Marshal(response)
	if err != nil {
		log.Println("フィード取得時にエラー", err)
		// TODO 異常系のjson返す
		return
	}
	w.Write(res)
}

func userDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	res, err := json.Marshal(registration.Login{
		LoginURL: loginURL,
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ログインURL", loginURL)
	w.Write(res)
}

// TODO gomniauth使用はmiddlewareに任せる
// TODO 登録周りはdomainに任せる
func callback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	providerName := ps.ByName("provider")
	user, err := middleware.GetUser(providerName, r.URL.RawQuery)
	accountID := user.IDForProvider(providerName)
	var userID uint

	if !registration.Registered(providerName, accountID) {
		userID, err = registration.Register(user.Name(), user.AvatarURL(), accountID, providerName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to register user with %s: %s", providerName, err), http.StatusInternalServerError)
			return
		}
	} else {
		// TODO 抽象化
		switch providerName {
		case "facebook":
			fb, err := orm.FindFacebookBy("AccountId", accountID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error when trying to register user with %s: %s", providerName, err), http.StatusInternalServerError)
			}
			u, err := orm.FindUserBy("FacebookAccountId", fb.Model.ID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error when trying to register user with %s: %s", providerName, err), http.StatusInternalServerError)
			}
			userID = u.Model.ID
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := json.Marshal(registration.Callback{
		UserID: userID,
	})
	if err != nil {
		log.Println(err)
	}
	w.Write(res)
}
