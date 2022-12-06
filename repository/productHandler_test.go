package repository

import (
	"fmt"
	"log"
	"serv/domain/model"
	"testing"

	"reflect"
	//"serv/usecase"

	//"github.com/driftprogramming/pgxpoolmock"

	//mocks "serv/repository/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	//"github.com/pashagolub/pgxmock/v2"
)

func TestGetProductFromStoreByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var itemID int = 1

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"})
	expect := []*model.Product{
		{ID: itemID, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Name, item.Category, item.Price, item.NominalPrice, item.Rating, item.Imgsrc)
	}

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(itemID).
		WillReturnRows(rows)

	repo := &ProductStore{
		db: db,
	}
	item, err := repo.GetProductFromStoreByID(itemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// // query error
	// mock.
	// 	ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
	// 	WithArgs(itemID).
	// 	WillReturnError(fmt.Errorf("db_error"))

	// _, err = repo.GetProductFromStoreByID(itemID)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// 	return
	// }
	// if err == nil {
	// 	t.Errorf("expected error, got nil")
	// 	return
	// }

	// // row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "asd")

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(itemID).
		WillReturnRows(rows)

	_, err = repo.GetProductFromStoreByID(itemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

}

func TestGetProductFromStoreByIDQueryErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var itemID int = 1
	repo := &ProductStore{
		db: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).
		AddRow(itemID, "IPhone", "phones", 50000, 50000, 0, nil)

	//query error
	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(itemID).
		WillReturnRows(rows).
		//WillReturnError(baseErrors.ErrServerError500)
		WillReturnError(fmt.Errorf("db_error"))

	log.Println("zzzz")

	_, err = repo.GetProductFromStoreByID(itemID)
	if err != nil {
		log.Println(err)
	}

	log.Println("xxxx")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
