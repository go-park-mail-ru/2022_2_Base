package repository

import (
	"serv/domain/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestChangeUserPasswordDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1

	var newPass string = "12345678910"
	mock.
		ExpectExec("UPDATE users").
		WithArgs(newPass, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &UserStore{
		db: db,
	}
	err = repo.ChangeUserPasswordDB(userID, newPass)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUpdateUsersAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	var addrID int = 1
	mock.
		ExpectExec("UPDATE").
		WithArgs("default", "default", "default", "default", false, addrID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &UserStore{
		db: db,
	}
	addr := model.Address{ID: 1, City: "default", Street: "default", House: "default", Flat: "default", Priority: false}

	err = repo.UpdateUsersAddress(userID, &addr)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestAddUsersAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	mock.
		ExpectExec("INSERT").
		WithArgs(userID, "default", "default", "default", "default", false).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &UserStore{
		db: db,
	}
	addr := model.Address{ID: 1, City: "default", Street: "default", House: "default", Flat: "default", Priority: false}

	err = repo.AddUsersAddress(userID, &addr)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDeleteUsersAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var addrID int = 1
	mock.
		ExpectExec("UPDATE").
		WithArgs(addrID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &UserStore{
		db: db,
	}
	err = repo.DeleteUsersAddress(addrID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
