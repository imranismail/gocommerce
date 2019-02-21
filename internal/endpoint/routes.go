package endpoint

import (
	"github.com/go-chi/chi"
	v1 "github.com/imranismail/ecommerce/internal/api/v1"
)

func (this *Endpoint) drawRoutes() {
	this.router.Route("/v1", func(router chi.Router) {
		router.Mount("/", v1.DiscoveryController())
		router.Mount("/users", v1.UserController(this.repo.User))
	})
}
