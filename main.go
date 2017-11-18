package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
)

func main() {
	// setup gomniauth facebook.New(クライアントID, 秘密の値, コールバックパス)
	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		facebook.New("300123313807716", "276dd1a14df05c304b0ebb3cc66a4c59", "http://localhost:8080/auth/callback/facebook"),
	)
	router := httprouter.New()
	router.GET("/v1/auth/:action/:provider", loginHandler)
	router.GET("/v1/users", userList)
	router.POST("/v1/users", userCreate)
	router.GET("/v1/users/:id", userDetail)
	router.PATCH("/v1/users/:id", userUpdate)

	// start the web server
	log.Println("Starting web server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
