package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type user struct {
	name       string
	avatarURL  string
	facebookID string
}

func userUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func userDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}