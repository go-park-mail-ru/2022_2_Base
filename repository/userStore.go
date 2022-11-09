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

func (us *UserStore) AddUser(in *model.UserDB) error {
	_, err := us.db.Exec(context.Background(), `INSERT INTO users (email, username, password) VALUES ($1, $2, $3);`, in.Email, in.Username, in.Password)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) UpdateUser(oldEmail string, in *model.UserProfile) error {
	_, err := us.db.Exec(context.Background(), `UPDATE users SET username = $1, phone = $2, avatar = $3  WHERE email = $4;`, in.Username, in.Phone, in.Avatar, oldEmail)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) GetUserByUsernameFromDB(userEmail string) (*model.UserDB, error) {
	rows, err := us.db.Query(context.Background(), `SELECT * FROM users WHERE email = $1`, userEmail)
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

func (us *UserStore) GetAdressesByUserIDFromDB(userID int) ([]*model.Adress, error) {
	adresses := []*model.Adress{}
	rows, err := us.db.Query(context.Background(), `SELECT city, street, house, priority FROM adress JOIN users ON adress.userid = users.id WHERE users.id  = $1`, userID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		dat := model.Adress{}
		err := rows.Scan(&dat.City, &dat.Street, &dat.House, &dat.Priority)
		if err != nil {
			return nil, err
		}
		adresses = append(adresses, &dat)
	}

	return adresses, nil
}

func (us *UserStore) GetPaymentMethodByUserIDFromDB(userID int) ([]*model.PaymentMethod, error) {
	payments := []*model.PaymentMethod{}
	rows, err := us.db.Query(context.Background(), `SELECT type, number, expiryDate, priority FROM payment JOIN users ON payment.userid = users.id WHERE users.id  = $1`, userID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		dat := model.PaymentMethod{}
		err := rows.Scan(&dat.Type, &dat.Number, &dat.ExpiryDate, &dat.Priority)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &dat)
	}

	return payments, nil
}
