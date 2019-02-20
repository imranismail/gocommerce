package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/internal/model"
	"github.com/imranismail/ecommerce/internal/repo"
	"github.com/imranismail/ecommerce/pkg/middleware"
)

type DiscoveryController struct {
	repo *repo.Repo
}

func NewDiscoveryController(r *repo.Repo) *DiscoveryController {
	return &DiscoveryController{r}
}

func (this DiscoveryController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()
	router.Use(middleware.BasicAuth("Basic realm=\"com.imranismail.ecommerce\"", "username", "password"))
	router.Get("/", this.list)
	router.ServeHTTP(w, r)
}

func (this DiscoveryController) list(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(model.Discovery{Version: "1.0.0", Description: "An eCommerce Service"})

	if err != nil {
		log.Fatal(err)
	}

	w.Write(b)
}
