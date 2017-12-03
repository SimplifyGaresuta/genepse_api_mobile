package main

import (
	"encoding/json"
	"fmt"
	"genepse_api/src/domain/registration"
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

// TODO gomniauth使用はmiddlewareに任せる
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
	if !registration.Registered(providerName, accountID) {
		userID = registration.Register(user.Name(), user.AvatarURL(), accountID, providerName)
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
