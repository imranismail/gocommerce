package repo

import (
	"database/sql"
	"fmt"

	"github.com/imranismail/ecommerce/internal/model"
)

type UserRepo struct {
	db     *sql.DB
	source string
}

func NewUserRepo(db *sql.DB, source string) UserRepo {
	return UserRepo{db, source}
}

func (r *UserRepo) Insert(user *model.User) error {
	stmt := fmt.Sprintf(`INSERT INTO %s (email, hashed_password) VALUES ($1, $2) RETURNING id`, r.source)
	return r.db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepo) Find(id int) (*model.User, error) {
	user := model.User{}
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE ID = $1`, r.source)
	err := r.db.QueryRow(stmt, id).Scan(&user.ID, &user.Email, &user.Password)
	return &user, err
}

func (r *UserRepo) All() ([]*model.User, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s`, r.source)
	rows, err := r.db.Query(stmt)

	defer rows.Close()

	users := make([]*model.User, 0)

	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
