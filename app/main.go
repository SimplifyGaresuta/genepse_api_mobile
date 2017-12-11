package main

import (
	"genepse_api/src/infra/cache"
	"genepse_api/src/infra/orm"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
)

func main() {
	err := orm.OpenMysql()
	if err != nil {
		log.Fatal(err)
	}
	defer orm.CloseMysql()
	orm.Setup()

	err = cache.DialRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	// setup gomniauth facebook.New(クライアントID, 秘密の値, コールバックパス)
	// TODO リファクタリング
	gomniauth.SetSecurityKey(globalSecret)
	var (
		fbClient       string
		fbClientSecret string
		fbCallback     string
	)
	if isDevelop() {
		fbClient = devFacebookClient
		fbClientSecret = devClientSecret
		fbCallback = devHost + "/v1/callback/facebook"
	} else {
		fbClient = proFacebookClient
		fbClientSecret = proClientSecret
		fbCallback = proHost + "/v1/callback/facebook"
	}
	gomniauth.WithProviders(
		facebook.New(fbClient, fbClientSecret, fbCallback),
	)

	router := httprouter.New()
	router.GET("/readiness_check", healthCheck)
	router.GET("/v1/login_url/:provider", login)
	router.GET("/v1/callback/:provider", callback)
	router.GET("/v1/users", userList)
	router.GET("/v1/users/:id", userDetail)
	router.PATCH("/v1/users/:id", userUpdate)
	router.POST("/v1/products", productCreate)
	router.PUT("/v1/products/:id", productUpdate)
	router.PUT("/v1/locations/:id", locationUpdate)

	log.Println("Start api server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}
