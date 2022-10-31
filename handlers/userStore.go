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

func (us *UserStore) UpdateUser(oldEmail string, in *model.UserProfile) (int64, error) {
	result, err := us.DB.Exec(`UPDATE users SET username = $1 WHERE email = $2;`, in.Username, oldEmail)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	if count == 0 {
		return 0, baseErrors.ErrNotFound404
	}
	return count, nil
}

func (us *UserStore) GetUserByUsernameFromDB(userEmail string) (*model.UserDB, error) {
	rows, err := us.DB.Query("SELECT * FROM users WHERE email = $1", userEmail)
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
