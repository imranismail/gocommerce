package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/internal/repo"
)

type Discovery struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

type DiscoveryController struct {
	Repo *repo.Repo
}

func NewDiscoveryController(r *repo.Repo) *DiscoveryController {
	return &DiscoveryController{r}
}

func (c DiscoveryController) Routes(r chi.Router) {
	r.Get("/", c.list)
}

func (c DiscoveryController) list(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Discovery{Version: "1.0.0", Description: "An eCommerce Service"})

	if err != nil {
		log.Fatal(err)
	}

	w.Write(b)
}
