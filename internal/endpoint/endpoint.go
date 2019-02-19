package endpoint

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	v1 "github.com/imranismail/ecommerce/internal/api/v1"
	"github.com/imranismail/ecommerce/internal/repo"
)

type Endpoint struct {
	router chi.Router
	repo   *repo.Repo
}

func New() (e Endpoint) {
	e.initRepo()
	e.initRoutes()
	return
}

func (e *Endpoint) Serve() {
	e.repo.Open()
	defer e.repo.Close()
	log.Fatal(http.ListenAndServe(":8080", e.router))
}

func (e *Endpoint) initRepo() {
	r := repo.New(&repo.Config{
		User:     "postgres",
		Password: "postgres",
		Database: "ecommerce",
		Host:     "localhost",
		Port:     "5432",
	})
	e.repo = &r
}

func (e *Endpoint) initRoutes() {
	e.router = chi.NewRouter()
	e.router.Use(middleware.RequestID)
	e.router.Use(middleware.RealIP)
	e.router.Use(middleware.Logger)
	e.router.Use(middleware.Recoverer)
	e.router.Use(middleware.Timeout(60 * time.Second))
	e.router.Use(middleware.SetHeader("content-type", "application/json"))
	e.router.Route("/v1", func(r chi.Router) {
		r.Route("/", v1.NewDiscoveryController(e.repo).Routes)
		r.Route("/users", v1.NewUserController(e.repo).Routes)
	})
}
