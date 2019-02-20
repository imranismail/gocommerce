package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/internal/model"
	"github.com/imranismail/ecommerce/internal/repo"
)

type (
	UserController struct {
		repo *repo.Repo
	}

	UserCreateReqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewUserController(r *repo.Repo) *UserController {
	return &UserController{r}
}

func (this UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()
	router.Get("/", this.list)
	router.Get("/{id}", this.get)
	router.Post("/", this.create)
	router.ServeHTTP(w, r)
}

func (this UserController) list(w http.ResponseWriter, r *http.Request) {
	u, err := this.repo.Users.All()

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

func (this UserController) get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := this.repo.Users.Find(id)

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

func (this UserController) create(w http.ResponseWriter, r *http.Request) {
	var body UserCreateReqBody
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.HashedPassword = body.Password
	user.Email = body.Email

	err = this.repo.Users.Insert(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
