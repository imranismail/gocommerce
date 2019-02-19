package endpoint

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/imranismail/ecommerce/internal/api"
	"github.com/imranismail/ecommerce/internal/repo"
)

type Endpoint struct {
	Router chi.Router
	Repo   *repo.Repo
}

func New() (e Endpoint) {
	e.initRepo()
	e.initRoutes()
	return
}

func (e *Endpoint) Serve() {
	e.Repo.Open()
	defer e.Repo.Close()
	log.Fatal(http.ListenAndServe(":8080", e.Router))
}

func (e *Endpoint) initRepo() {
	r := repo.New(repo.Config{
		User:     "postgres",
		Password: "postgres",
		Database: "ecommerce",
		Host:     "localhost",
		Port:     "5432",
	})
	e.Repo = &r
}

func (e *Endpoint) initRoutes() {
	e.Router = chi.NewRouter()
	e.Router.Use(middleware.RequestID)
	e.Router.Use(middleware.RealIP)
	e.Router.Use(middleware.Logger)
	e.Router.Use(middleware.Recoverer)
	e.Router.Use(middleware.Timeout(60 * time.Second))
	e.Router.Use(middleware.SetHeader("content-type", "application/json"))
	e.Router.Use(middleware.WithValue("repo", e.Repo))
	e.Router.Route("/", api.DiscoveryController)
	e.Router.Route("/users", api.UserController)
}
