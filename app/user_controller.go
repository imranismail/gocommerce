package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/model"
)

func (a *App) UserController(r chi.Router) {
	r.Get("/", a.userIndex)
	r.Post("/", a.userCreate)
	r.Get("/{id}", a.userShow)
}

func (a *App) userIndex(w http.ResponseWriter, r *http.Request) {
	u, err := a.Repo.Users.All()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (a *App) userShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := model.User{ID: id, Email: "imran.codely@gmail.com"}

	b, err := json.Marshal(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (a *App) userCreate(w http.ResponseWriter, r *http.Request) {
	u := model.User{Email: "imran.codely@gmail.com", Password: "admin123"}
	err := a.Repo.Users.Insert(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
