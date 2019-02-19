package app

import (
	"github.com/go-chi/chi/middleware"
	"github.com/imranismail/ecommerce/api"
)

func (app *App) InitRoutes() {
	app.Router.Use(middleware.SetHeader("content-type", "application/json"))
	app.Router.Route("/", api.DiscoveryController)
	app.Router.Route("/users", api.UserController)
}
