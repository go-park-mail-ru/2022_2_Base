package repository

import (
	"context"
	"database/sql"
	baseErrors "serv/domain/errors"
	"serv/domain/model"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) AddUser(in *model.UserDB) (uint, error) {
	//result, err := us.db.Exec(context.Background(), `INSERT INTO users (email, username, password) VALUES ($1, $2, $3);`, in.Email, in.Username, in.Password)
	_, err := us.db.Exec(context.Background(), `INSERT INTO users (email, username, password) VALUES ($1, $2, $3);`, in.Email, in.Username, in.Password)
	if err != nil {
		return 0, err
	}
	//affected, err := result.RowsAffected()
	affected, err := 1, nil
	if err != nil {
		return 0, err
	}
	return uint(affected), nil
}

func (us *UserStore) UpdateUser(oldEmail string, in *model.UserProfile) (int64, error) {
	//result, err := us.db.Exec(context.Background(), `UPDATE users SET username = $1, phone = $2, avatar = $3  WHERE email = $4;`, in.Username, in.Phone, in.Avatar, oldEmail)
	_, err := us.db.Exec(context.Background(), `UPDATE users SET username = $1, phone = $2, avatar = $3  WHERE email = $4;`, in.Username, in.Phone, in.Avatar, oldEmail)
	if err != nil {
		return 0, err
	}
	//count, err := result.RowsAffected()
	count, err := int64(1), nil
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, baseErrors.ErrNotFound404
	}
	return count, nil
}

func (us *UserStore) GetUserByUsernameFromDB(userEmail string) (*model.UserDB, error) {
	rows, err := us.db.Query(context.Background(), "SELECT * FROM users WHERE email = $1", userEmail)
	if err == sql.ErrNoRows {
		return nil, baseErrors.ErrUnauthorized401
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user := model.UserDB{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Phone, &user.Avatar)
		if err != nil {
			return nil, err
		}

	}
	return &user, nil
}
