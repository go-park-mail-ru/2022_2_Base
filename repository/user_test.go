package repository

import (
	"reflect"
	"serv/domain/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userEmail string = "a@a"
	var userName string = "art"
	var userPass string = "12345678"
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(userEmail, userName, userPass).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &UserStore{
		db: db,
	}
	in := model.UserDB{Email: userEmail, Username: userName, Password: userPass}
	err = repo.AddUser(&in)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	var userEmail string = "a@a"
	var userName string = "art"
	var userPhone string = "12345678910"
	mock.
		ExpectExec("UPDATE users").
		WithArgs(userEmail, userName, userPhone, "", userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &UserStore{
		db: db,
	}
	in := model.UserProfile{Email: userEmail, Username: userName, Phone: userPhone}
	err = repo.UpdateUser(userID, &in)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetUserByUsernameFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var email string = "a@a"
	expect := model.UserDB{ID: 1, Username: "art", Email: "a@a", Password: "12345678", Phone: nil, Avatar: nil}

	mock.
		ExpectQuery("SELECT").
		WithArgs(email).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "email", "username", "password", "phone", "avatar"}).AddRow(1, "a@a", "art", "12345678", nil, nil)
			return rr
		}())
	repo := &UserStore{
		db: db,
	}
	user, err := repo.GetUserByUsernameFromDB(email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, &expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
		return
	}
	//query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(email).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id"}).AddRow(1)
			return rr
		}())
	_, err = repo.GetUserByUsernameFromDB(email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetUserByIDFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var id int = 1
	expect := "art"

	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"username"}).AddRow("art")
			return rr
		}())
	repo := &UserStore{
		db: db,
	}
	username, err := repo.GetUsernameByIDFromDB(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(username, expect) {
		t.Errorf("results not match, want %v, have %v", expect, username)
		return
	}
	//query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "password"}).AddRow(1, "s")
			return rr
		}())
	_, err = repo.GetUsernameByIDFromDB(id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetAddressesByUserIDFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var id int = 1
	expect := []*model.Address{
		{ID: 1, City: "default", Street: "default", House: "default", Flat: "default", Priority: false},
	}
	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "city", "street", "house", "flat", "priority"}).AddRow(1, "default", "default", "default", "default", false)
			return rr
		}())
	repo := &UserStore{
		db: db,
	}
	adresses, err := repo.GetAddressesByUserIDFromDB(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(adresses, expect) {
		t.Errorf("results not match, want %v, have %v", expect, adresses)
		return
	}
	//query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "password"}).AddRow(1, "s")
			return rr
		}())
	_, err = repo.GetAddressesByUserIDFromDB(id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetPaymentMethodByUserIDFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var id int = 1
	expectTime := time.Unix(1, 0)
	expect := []*model.PaymentMethod{
		{ID: 1, PaymentType: "default", Number: "default", ExpiryDate: expectTime, Priority: false},
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "paymentType", "number", "expiryDate", "priority"}).AddRow(1, "default", "default", expectTime, false)
			return rr
		}())
	repo := &UserStore{
		db: db,
	}
	adresses, err := repo.GetPaymentMethodByUserIDFromDB(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(adresses, expect) {
		t.Errorf("results not match, want %v, have %v", expect, adresses)
		return
	}
	//query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "password"}).AddRow(1, "s")
			return rr
		}())
	_, err = repo.GetPaymentMethodByUserIDFromDB(id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
