package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/imranismail/ecommerce/repo"
)

type App struct {
	Router chi.Router
	Repo   *repo.Repo
}

func New() App {
	a := App{}
	a.initializeRepo()
	a.initializeRouter()
	a.drawRoutes()
	return a
}

func (a *App) Serve() {
	a.Repo.Open()
	log.Fatal(http.ListenAndServe(":8080", a.Router))
	a.Repo.Close()
}

func (a *App) initializeRepo() {
	r, err := repo.New(repo.Config{
		User:     "postgres",
		Password: "postgres",
		Database: "ecommerce",
		Host:     "localhost",
		Port:     "5432",
	})

	if err != nil {
		log.Fatal(err)
	}

	a.Repo = &r
}

func (a *App) initializeRouter() {
	a.Router = chi.NewRouter()
	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.RealIP)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.Timeout(60 * time.Second))
}
