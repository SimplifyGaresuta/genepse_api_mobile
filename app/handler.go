package main

import (
	"net/http"

	"genepse_api/src/repository/entity"

	"github.com/julienschmidt/httprouter"
)

func userUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hey := entity.Hey()
	w.Write([]byte(hey))
}

func userDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func userCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
