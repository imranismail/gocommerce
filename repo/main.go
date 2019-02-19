package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var db *sql.DB

type Repo struct {
	Users  UserRepo
	db     *sql.DB
	config Config
}

func New(cfg Config) (r Repo, err error) {
	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
		err = errors.Errorf("All fields must be set (%s)", cfg)
		return
	}

	r.config = cfg
	r.Open()
	r.db = db
	r.Users = UserRepo{db: db}

	return
}

func (r *Repo) Open() {
	var err error
	dsn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, r.config.Host, r.config.Port, r.config.User, r.config.Password, r.config.Database)
	db, err = sql.Open("postgres", dsn)

	if err != nil {
		err = errors.Wrapf(err, "Couldn't open connection to postgres database (%s)", r.config)
		log.Fatal(err)
	}

	if err != nil {
		err = errors.Wrapf(err, "Couldn't ping the postgres database (%s)", r.config)
		log.Fatal(err)
	}
}

func (r *Repo) Close() {
	err := r.db.Close()

	if err != nil {
		err = errors.Wrap(err, "Couldn't close connection to postgres database")
		log.Fatal(err)
	}
}
