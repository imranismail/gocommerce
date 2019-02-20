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

func NewUserRepo(db *sql.DB, source string) *UserRepo {
	return &UserRepo{db, source}
}

func (this *UserRepo) Insert(user *model.User) error {
	stmt := fmt.Sprintf(`INSERT INTO %s (email, hashed_password) VALUES ($1, $2) RETURNING id`, this.source)
	return this.db.QueryRow(stmt, user.Email, user.HashedPassword).Scan(&user.ID)
}

func (this *UserRepo) Find(id int) (*model.User, error) {
	user := model.User{}
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE ID = $1`, this.source)
	err := this.db.QueryRow(stmt, id).Scan(&user.ID, &user.Email, &user.HashedPassword)
	return &user, err
}

func (this *UserRepo) All() ([]*model.User, error) {
	stmt := fmt.Sprintf(`SELECT * FROM %s`, this.source)
	rows, err := this.db.Query(stmt)

	defer rows.Close()

	users := make([]*model.User, 0)

	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.ID, &user.Email, &user.HashedPassword)

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
