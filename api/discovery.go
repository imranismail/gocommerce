package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imranismail/ecommerce/middleware"
)

func DiscoveryController(r chi.Router) {
	r.Use(middleware.BasicAuth("imranismail", "password"))
	r.Get("/", listDiscoveryDetails)
}

func listDiscoveryDetails(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"version": "1.0.0", "description": "An eCommerce Service"}`))
}
