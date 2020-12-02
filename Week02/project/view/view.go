package view

import (
	"net/http"
	"project/controller"
	"project/err"
)

func UserView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		name := r.FormValue("name")
		if len(name) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("`name` argument missing."))
			return
		}
		user, e := controller.GetUser(name)
		if e != nil && e == err.NotExist {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
		} else {
			w.Write(controller.EncodeUser(user))
		}

	case "POST":
		name, password := r.FormValue("name"), r.FormValue("password")
		if e := controller.SaveUser(name, password); e != nil && e == err.Duplicate {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("conflict"))
		} else {
			w.Write([]byte("success."))
		}
	}
}