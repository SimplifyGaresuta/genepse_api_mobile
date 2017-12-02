package main

import (
	"genepse_api/src/infra/orm"
	"genepse_api/src/middleware"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
)

func main() {
	//orm.Setup()
	err := orm.OpenMysql()
	if err != nil {
		log.Fatal(err)
	}
	defer orm.CloseMysql()
	orm.Setup()
	//	user := orm.FindUser(1)
	//	log.Println(*user)
	os.Exit(0)
	// setup gomniauth facebook.New(クライアントID, 秘密の値, コールバックパス)
	gomniauth.SetSecurityKey(globalSecret)
	gomniauth.WithProviders(
		facebook.New(facebookClient, clientSecret, "http://localhost:8080/auth/callback/facebook"),
	)
	router := httprouter.New()
	router.GET("/v1/auth/:action/:provider", middleware.LoginHandler)
	router.GET("/v1/users", userList)
	router.POST("/v1/users", userCreate)
	router.GET("/v1/users/:id", userDetail)
	router.PATCH("/v1/users/:id", userUpdate)

	// start the web server
	log.Println("Starting web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
