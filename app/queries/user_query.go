package queries

import (
	"apigo/app/models"

	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserById(id int) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE id = $1`

	err := q.Get(&user, query, id)
	return user, err
}

func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE email = $1`

	err := q.Get(&user, query, email)
	return user, err
}

func (q *UserQueries) CreateUser(user *models.User) error {
	query := `INSERT INTO users VALUES($1, $2, $3, $4, $5)`

	_, err := q.Exec(
		query,
		user.ID, user.CreatedAt, user.UpdatedAt, user.Email, user.PassHash,
	)

	return err
}

func (q *UserQueries) HasUserByEmail(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	err := q.Get(&exists, query, email)

	if err != nil {
		return false, err
	} else {
		return exists, nil
	}
}
