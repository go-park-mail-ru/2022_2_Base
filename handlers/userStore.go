package handlers

import (
	"database/sql"
	baseErrors "serv/errors"
	"serv/model"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type UserStore struct {
	DB *sql.DB
}

func (us *UserStore) AddUser(in *model.UserDB) (uint, error) {
	result, err := us.DB.Exec(`INSERT INTO users (email, username, password) VALUES ($1, $2, $3);`, in.Email, in.Username, in.Password)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	lastID, err := result.LastInsertId()
	return uint(lastID), nil
}

func (us *UserStore) GetUserByUsernameAndPasswordFromDB(userEmail string, password string) (*model.UserDB, error) {
	rows, err := us.DB.Query("SELECT * FROM users WHERE email = $1 AND password = $2", userEmail, password)
	if err == sql.ErrNoRows {
		return nil, baseErrors.ErrUnauthorized401
	}
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}
	defer rows.Close()
	user := model.UserDB{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			return nil, baseErrors.ErrServerError500
		}

	}
	return &user, nil
}
