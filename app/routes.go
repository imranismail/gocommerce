package app

import (
	"github.com/go-chi/chi/middleware"
)

func (a *App) drawRoutes() {
	a.Router.Use(middleware.SetHeader("content-type", "application/json"))
	a.Router.Route("/", a.DiscoveryController)
	a.Router.Route("/users", a.UserController)
}
