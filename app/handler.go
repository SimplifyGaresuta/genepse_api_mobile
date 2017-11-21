package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func userUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("userList"))
}

func userDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
