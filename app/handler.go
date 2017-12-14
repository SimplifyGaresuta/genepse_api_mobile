package main

import (
	"encoding/json"
	"fmt"
	"genepse_api/src/domain/detail"
	"genepse_api/src/domain/feed"
	"genepse_api/src/domain/location"
	"genepse_api/src/domain/registration"
	"genepse_api/src/infra/orm"
	"genepse_api/src/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
)

type Response interface{}

func returnJSON(w http.ResponseWriter, res Response) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// TODO 本番はjson.Marshalを使用する
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}
	w.Write(r)
	return nil
}

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	if len(query["limit"]) < 1 {
		log.Println("limitを指定して下さい。")
		// TODO 異常系json
		return
	}
	limit, err := strconv.Atoi(query["limit"][0])
	if err != nil {
		log.Println("クエリパラメータが不正です", err)
		// TODO 異常系のjson
		return
	}
	if len(query["offset"]) < 1 {
		log.Println("offsetを指定して下さい。")
		// TODO 異常系json
		return
	}
	offset, err := strconv.Atoi(query["offset"][0])
	if err != nil {
		log.Println("クエリパラメータが不正です", err)
		// TODO 異常系のjson
		return
	}
	response, err := feed.GetResponse(limit, offset)
	if err != nil {
		log.Println("フィード取得時にエラー", err)
		// TODO 異常系のjson返す
		return
	}
	returnJSON(w, response)
}

func userDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Println("idが不正です。")
		// TODO 異常系json
		return
	}
	user, err := detail.GetUser(id)
	if err != nil {
		log.Println("プロフィール取得時にエラー", err)
		// TODO 異常系json
		return
	}
	returnJSON(w, user)
}

func userUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Println("idが不正です。")
		// TODO 異常系json
		return
	}
	if err := detail.UpdateUser(uint(id), r.Body); err != nil {
		log.Println("プロフィール更新時にエラー", err)
		// TODO 異常系json
		return
	}
}

func productCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	file, _, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
	}
	if err := r.ParseForm(); err != nil {
		log.Println("リクエストbodyが不正です。", err)
		// TODO 異常系json
		return
	}
	userID, err := strconv.Atoi(strings.Join(r.Form["user_id"], ""))
	if err != nil {
		// TODO 異常系json
		return
	}
	operator := &detail.ProductOperator{
		UserID: userID,
		Title:  strings.Join(r.Form["title"], ""),
		URL:    strings.Join(r.Form["url"], ""),
		Ctx:    r.Context(),
		File:   file,
	}
	response, err := detail.CreateProduct(operator)
	if err != nil {
		log.Println("作品登録時にエラー", err)
		// TODO 異常系json
		return
	}
	returnJSON(w, response)
}
func productUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	file, _, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Println("idが不正です。")
		// TODO 異常系json
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("リクエストbodyが不正です。", err)
		// TODO 異常系json
		return
	}
	operator := &detail.ProductOperator{
		ID:    id,
		Title: strings.Join(r.Form["title"], ""),
		URL:   strings.Join(r.Form["url"], ""),
		Ctx:   r.Context(),
		File:  file,
	}
	if err := detail.UpdateProduct(operator); err != nil {
		log.Println("作品更新時にエラー", err)
		// TODO 異常系json
		return
	}
}

func locationUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	id := ps.ByName("id")
	if err := location.UpdateLocation(id, r.Body); err != nil {
		log.Println("位置情報更新時にエラー", err)
		// TODO 異常系json
		return
	}
}

func nearUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query := r.URL.Query()
	if len(query["user_id"]) < 1 {
		log.Println("user_idを指定して下さい。")
		// TODO 異常系json
		return
	}
	userID := query["user_id"][0]
	res, err := location.GetNearUsers(userID, 100)
	if err != nil {
		log.Println("近くのユーザー検索時にエラー。", err)
		// TODO 異常系json
		return
	}
	returnJSON(w, res)
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
	res := registration.Login{
		LoginURL: loginURL,
	}
	fmt.Println("login_URL", loginURL)
	returnJSON(w, res)
}

// TODO gomniauth使用はmiddlewareに任せる
// TODO 登録周りはdomainに任せる
func callback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	providerName := ps.ByName("provider")
	user, err := middleware.GetUser(providerName, r.URL.RawQuery)
	if err != nil {
		log.Println(err)
		// TODO 異常系json
		return
	}
	accountID := user.IDForProvider(providerName)

	var userID uint
	if registration.Registered(providerName, accountID) {
		var provider orm.Provider
		switch providerName {
		case "facebook":
			provider = &orm.FacebookAccount{}
		}
		userID, err = registration.UserID(provider, accountID)
		if err != nil {
			log.Println(err)
			// TODO 異常系json
			return
		}
	} else { // 登録
		userID, err = registration.Register(user.Name(), user.AvatarURL(), accountID, providerName)
		if err != nil {
			log.Println(err)
			// TODO 異常系json
			return
		}
	}
	res := registration.Callback{
		UserID: userID,
	}
	returnJSON(w, res)
}

func healthCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("I am healthy!")
}
