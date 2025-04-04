package queries

import (
	"apigo/app/models"

	"github.com/jmoiron/sqlx"
)

// = = = = = = = = =
// Regular Queries
// = = = = = = = = =
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

// = = = = = = = = = = =
// Transaction Queries
// = = = = = = = = = = =
type UserTxQueries struct {
	*sqlx.Tx
}

func (tx *UserTxQueries) CreateUser(user *models.User, profile *models.UserProfile) error {
	userQuery := `
		INSERT INTO users 
		(id, created_at, updated_at, email, pass_hash) VALUES
		(:id, :created_at, :updated_at, :email, :pass_hash)
	`

	profileQuery := `
		INSERT INTO user_profiles
		(user_id, nickname, bio) VALUES
		(:user_id, :nickname, :bio)
	`

	_, err := tx.NamedExec(userQuery, user)
	if err != nil {
		return err
	}

	_, err = tx.NamedExec(profileQuery, profile)
	if err != nil {
		return err
	}

	return nil
}

func (tx *UserTxQueries) HasUserByNicknameOrEmail(nickname string, email string) (bool, error) {
	nicknameQuery := `SELECT EXISTS(SELECT 1 FROM user_profiles WHERE nickname = $1)`
	emailQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var match bool

	if err := tx.Get(&match, nicknameQuery, nickname); err != nil {
		return false, err
	} else if match {
		return true, nil
	}

	if err := tx.Get(&match, emailQuery, email); err != nil {
		return false, err
	} else if match {
		return true, nil
	}

	return false, nil
}

// func (up *UserProfileQueries) GetProfileById(id string, includeUser bool) (*models.UserProfile, *models.User, error) {
// 	return nil, nil, nil
// }
