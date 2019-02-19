package app

import (
	"time"

	"github.com/go-chi/chi/middleware"
)

func (app *App) InitEndpoint() {
	app.Router.Use(middleware.RequestID)
	app.Router.Use(middleware.RealIP)
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)
	app.Router.Use(middleware.Timeout(60 * time.Second))
}
