package repo

import (
	"database/sql"

	"github.com/imranismail/ecommerce/model"
)

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) Insert(user *model.User) error {
	return r.db.QueryRow(`INSERT INTO users(email, hashed_password) VALUES($1, $2) RETURNING id`, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepo) All() ([]*model.User, error) {
	rows, err := r.db.Query(`
		SELECT * FROM users
	`)

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
