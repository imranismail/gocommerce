package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/middleware"
	"github.com/imranismail/ecommerce/model"
)

func UserController(router chi.Router) {
	router.Use(middleware.BasicAuth("imranismail", "password"))
	router.Get("/", listUsers)
	router.Get("/{id}", showUser)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users := []model.User{
		model.User{ID: 1, Email: "imran.codely@gmail.com"},
	}

	json, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func showUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := model.User{ID: id, Email: "imran.codely@gmail.com"}

	json, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}
