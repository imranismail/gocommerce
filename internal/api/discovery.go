package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Discovery struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

func DiscoveryController(r chi.Router) {
	r.Get("/", discoveryIndex)
}

func discoveryIndex(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Discovery{Version: "1.0.0", Description: "An eCommerce Service"})

	if err != nil {
		log.Fatal(err)
	}

	w.Write(b)
}
