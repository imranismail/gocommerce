package endpoint

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	this.router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
		middleware.SetHeader("content-type", "application/json"),
		middleware.SetHeader("accept", "application/json"),
	)

	return &this
}

func (this *Endpoint) Serve() {
	this.repo.Open()
	defer this.repo.Close()
	this.drawRoutes()
	log.Fatal(http.ListenAndServe(":8080", this.router))
}
