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

func (us *UserStore) GetUsers() ([]*model.UserDB, error) {
	users := []*model.UserDB{}
	rows, err := us.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}
	defer rows.Close()

	for rows.Next() {
		dat := model.UserDB{}
		err := rows.Scan(&dat.ID, &dat.Email, &dat.Username, &dat.Password)
		if err != nil {
			return nil, baseErrors.ErrServerError500
		}
		users = append(users, &dat)
	}

	return users, nil
}
