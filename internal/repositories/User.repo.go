package repositories

import (
	"errors"
	"julianmindria/backendCoffee/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r RepoUser) CreateUser(data *models.User) (string, error) {
	queryUser := `INSERT INTO users(
		username,
		email,
		pass,
		phone,
		roles,
		user_image
		) 
		VALUES(
			:username, 
			:email, 
			:pass, 
			:phone,
			:roles,
			:user_image
		)`

	_, err := r.NamedExec(queryUser, data)
	if err != nil {
		return "", err
	}
	return "1 data user created", nil
}

func (r RepoUser) UpdateUser(data *models.User) (string, error) {
	queryUser := `UPDATE public.users
	SET username=:username,email=:email,pass=:pass,phone=:phone
	WHERE id_user=:id_user
	`
	_, err := r.NamedExec(queryUser, data)
	if err != nil {
		return "", err
	}
	return "1 data Updated", nil

}

func (r RepoUser) DeleteUser(data *models.User) (string, error) {
	queryUser := `DELETE FROM public.users WHERE id_user=:id_user`
	_, err := r.NamedExec(queryUser, data)
	if err != nil {
		return "", err
	}
	return "1 data Deleted", nil

}

func (r RepoUser) Get_Users(data *models.User, search string) ([]models.User, error) {
	users_data := []models.User{}
	if search == "" {
		r.Select(&users_data, `select u.username, u.roles, u.created_at FROM public.users u`)
	} else {
		r.Select(&users_data, `select u.username, u.roles, u.created_at FROM public.users u`, search)
	}
	if len(users_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return users_data, nil
}

func (r RepoUser) GetAuthUser(user string) (*models.User, error) {
	var result models.User
	q := `SELECT id_user, username, roles, pass FROM public.users WHERE username = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}
