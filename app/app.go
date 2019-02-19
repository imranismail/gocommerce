package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type App struct {
	Router chi.Router
}

func (app *App) Serve() {
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

func New() App {
	app := App{}
	app.Router = chi.NewRouter()
	app.InitEndpoint()
	app.InitRoutes()
	return app
}
