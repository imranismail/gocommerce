package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/internal/model"
	"github.com/imranismail/ecommerce/internal/repo"
)

type UserController struct {
	repo *repo.Repo
}

func NewUserController(r *repo.Repo) *UserController {
	return &UserController{r}
}

func (c UserController) Routes(r chi.Router) {
	r.Get("/", c.list)
	r.Get("/{id}", c.get)
	r.Post("/", c.create)
}

func (c UserController) list(w http.ResponseWriter, r *http.Request) {
	u, err := c.repo.Users.All()

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

func (c UserController) get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := c.repo.Users.Find(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	b, err := json.Marshal(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (c UserController) create(w http.ResponseWriter, r *http.Request) {
	u := model.User{Email: "imran.codely@gmail.com", Password: "admin123"}
	err := c.repo.Users.Insert(&u)

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
