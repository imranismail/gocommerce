package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/internal/model"
	"github.com/imranismail/ecommerce/internal/repo"
)

func UserController(r chi.Router) {
	r.Get("/", userIndex)
	r.Post("/", userCreate)
	r.Get("/{id}", userShow)
}

func userIndex(w http.ResponseWriter, r *http.Request) {
	u, err := repo.Of(r.Context()).Users.All()

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

func userShow(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := repo.Of(r.Context()).Users.Find(id)

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

func userCreate(w http.ResponseWriter, r *http.Request) {
	u := model.User{Email: "imran.codely@gmail.com", Password: "admin123"}
	err := repo.Of(r.Context()).Users.Insert(&u)

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
