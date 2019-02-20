package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Repo struct {
	Users  *UserRepo
	db     *sql.DB
	config *Config
}

func New(config *Config) *Repo {
	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.Database == "" {
		log.Fatal(errors.Errorf("All fields must be set (%s)", config))
	}

	this := Repo{}
	this.config = config

	return &this
}

func (this *Repo) Open() {
	dsn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, this.config.Host, this.config.Port, this.config.User, this.config.Password, this.config.Database)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		err = errors.Wrapf(err, "Couldn't open connection to postgres database (%s)", this.config)
		log.Fatal(err)
	}

	if err != nil {
		err = errors.Wrapf(err, "Couldn't ping the postgres database (%s)", this.config)
		log.Fatal(err)
	}

	this.db = db
	this.Users = NewUserRepo(db, "users")
}

func (this *Repo) Close() {
	err := this.db.Close()

	if err != nil {
		err = errors.Wrap(err, "Couldn't close connection to postgres database")
		log.Fatal(err)
	}
}
