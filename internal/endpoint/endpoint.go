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

func New() *Endpoint {
	this := Endpoint{}

	// Repo
	this.repo = repo.New(&repo.Config{
		User:     "postgres",
		Password: "postgres",
		Database: "ecommerce",
		Host:     "localhost",
		Port:     "5432",
	})

	// Router
	this.router = chi.NewRouter()
	this.router.Use(middleware.RequestID)
	this.router.Use(middleware.RealIP)
	this.router.Use(middleware.Logger)
	this.router.Use(middleware.Recoverer)
	this.router.Use(middleware.Timeout(60 * time.Second))
	this.router.Use(middleware.SetHeader("content-type", "application/json"))
	this.router.Use(middleware.SetHeader("accept", "application/json"))
	this.router.Route("/v1", func(router chi.Router) {
		router.Mount("/", v1.DiscoveryController())
		router.Mount("/users", v1.UserController(this.repo.User))
	})

	return &this
}

func (this *Endpoint) Serve() {
	this.repo.Open()
	defer this.repo.Close()
	log.Fatal(http.ListenAndServe(":8080", this.router))
}
