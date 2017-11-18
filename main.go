package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
)

func main() {
	flag.Parse() // parse the flags

	// setup gomniauth facebook.New(クライアントID, 秘密の値, コールバックパス)
	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		facebook.New("300123313807716", "276dd1a14df05c304b0ebb3cc66a4c59", "http://localhost:8080/auth/callback/facebook"),
	)

	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/user/update/", profileUpdate)

	// start the web server
	log.Println("Starting web server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
