package repository

import (
	"database/sql"
	baseErrors "serv/domain/errors"
	"serv/domain/model"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type UserStoreInterface interface {
	AddUser(in *model.UserDB) (int, error)
	UpdateUser(userID int, in *model.UserProfile) error
	ChangeUserPasswordDB(userID int, newPass string) error
	UpdateUsersAddress(adressID int, in *model.Address) error
	AddUsersAddress(userID int, in *model.Address) error
	DeleteUsersAddress(addressID int) error
	UpdateUsersPayment(paymentID int, in *model.PaymentMethod) error
	AddUsersPayment(userID int, in *model.PaymentMethod) error
	DeleteUsersPayment(paymentID int) error
	GetUserByUsernameFromDB(userEmail string) (*model.UserDB, error)
	GetUsernameByIDFromDB(userID int) (string, error)
	GetAddressesByUserIDFromDB(userID int) ([]*model.Address, error)
	GetPaymentMethodByUserIDFromDB(userID int) ([]*model.PaymentMethod, error)
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) UserStoreInterface {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) AddUser(in *model.UserDB) (int, error) {
	//result, err := us.db.Exec(`INSERT INTO users (email, username, password) VALUES ($1, $2, $3);`, in.Email, in.Username, in.Password)
	id := 0
	err := us.db.QueryRow(`INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id;`, in.Email, in.Username, in.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }
	return int(id), nil
}

func (us *UserStore) UpdateUser(userID int, in *model.UserProfile) error {
	_, err := us.db.Exec(`UPDATE users SET email = $1, username = $2, phone = $3, avatar = $4 WHERE id = $5;`, in.Email, in.Username, in.Phone, in.Avatar, userID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) ChangeUserPasswordDB(userID int, newPass string) error {
	_, err := us.db.Exec(`UPDATE users SET password = $1 WHERE id = $2;`, newPass, userID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) UpdateUsersAddress(adressID int, in *model.Address) error {
	_, err := us.db.Exec(`UPDATE address SET city = $1, street = $2, house = $3, flat = $4, priority = $5 WHERE id = $6;`, in.City, in.Street, in.House, in.Flat, in.Priority, adressID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) AddUsersAddress(userID int, in *model.Address) error {
	_, err := us.db.Exec(`INSERT INTO address (userid, city, street, house, flat, priority) VALUES ($1, $2, $3, $4, $5, $6);`, userID, in.City, in.Street, in.House, in.Flat, in.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) DeleteUsersAddress(addressID int) error {
	_, err := us.db.Exec(`UPDATE address SET deleted = true WHERE id = $1;`, addressID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) UpdateUsersPayment(paymentID int, in *model.PaymentMethod) error {
	_, err := us.db.Exec(`UPDATE payment SET paymentType = $1, number = $2, expirydate = $3, priority = $4 WHERE id = $5;`, in.PaymentType, in.Number, in.ExpiryDate, in.Priority, paymentID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) AddUsersPayment(userID int, in *model.PaymentMethod) error {
	_, err := us.db.Exec(`INSERT INTO payment (userid, paymentType, number, expirydate, priority) VALUES ($1, $2, $3, $4, $5);`, userID, in.PaymentType, in.Number, in.ExpiryDate, in.Priority)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) DeleteUsersPayment(paymentID int) error {
	_, err := us.db.Exec(`UPDATE payment SET deleted = true WHERE id = $1;`, paymentID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) GetUserByUsernameFromDB(userEmail string) (*model.UserDB, error) {
	rows, err := us.db.Query(`SELECT * FROM users WHERE email = $1`, userEmail)
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

func (us *UserStore) GetUsernameByIDFromDB(userID int) (string, error) {
	rows, err := us.db.Query(`SELECT username FROM users WHERE id = $1`, userID)
	if err == sql.ErrNoRows {
		return "", err
	}
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			return "", err
		}
	}
	return username, nil
}

func (us *UserStore) GetAddressesByUserIDFromDB(userID int) ([]*model.Address, error) {
	adresses := []*model.Address{}
	rows, err := us.db.Query(`SELECT address.id, city, street, house, flat, priority FROM address JOIN users ON address.userid = users.id WHERE users.id  = $1 AND deleted = false`, userID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		dat := model.Address{}
		err := rows.Scan(&dat.ID, &dat.City, &dat.Street, &dat.House, &dat.Flat, &dat.Priority)
		if err != nil {
			return nil, err
		}
		adresses = append(adresses, &dat)
	}

	return adresses, nil
}

func (us *UserStore) GetPaymentMethodByUserIDFromDB(userID int) ([]*model.PaymentMethod, error) {
	payments := []*model.PaymentMethod{}
	rows, err := us.db.Query(`SELECT payment.id, paymentType, number, expiryDate, priority FROM payment JOIN users ON payment.userid = users.id WHERE users.id  = $1 AND deleted = false`, userID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		dat := model.PaymentMethod{}
		err := rows.Scan(&dat.ID, &dat.PaymentType, &dat.Number, &dat.ExpiryDate, &dat.Priority)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &dat)
	}

	return payments, nil
}
