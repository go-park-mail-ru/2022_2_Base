package repository

import (
	"reflect"
	"serv/domain/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	mock.
		ExpectExec("INSERT INTO orders").
		WithArgs(userID, "cart", "not started").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &ProductStore{
		db: db,
	}
	err = repo.CreateCart(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetOrderItemsFromStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var orderID int = 1
	expect := []*model.OrderItem{
		{Count: 1, Item: &model.Product{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil}},
	}

	mock.
		ExpectQuery("SELECT count").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, 1, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	repo := &ProductStore{
		db: db,
	}

	items, err := repo.GetOrderItemsFromStore(orderID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(items[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], items)
		return
	}

	//query error
	mock.
		ExpectQuery("SELECT count").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(0, "")
			return rr
		}())

	_, err = repo.GetOrderItemsFromStore(orderID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	var orderID int = 1
	var avatar string = "av1"
	var promo string = "A10zzzzz"
	expectItems := []*model.OrderItem{
		{Count: 1, Item: &model.Product{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: &avatar}},
	}
	expectTime := time.Unix(1, 0)
	expect := model.Order{ID: 1, UserID: 1, Items: expectItems, OrderStatus: "cart", PaymentStatus: "not started", AddressID: 1, PaymentcardID: 1, CreationDate: &expectTime, DeliveryDate: &expectTime, Promocode: &promo}

	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, 1, "cart", "not started", 1, 1, &expectTime, &expectTime, promo)
			return rr
		}())

	mock.
		ExpectQuery("SELECT count").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, 1, "IPhone", "phones", 50000, 50000, 0, &avatar)
			return rr
		}())

	repo := &ProductStore{
		db: db,
	}

	cart, err := repo.GetCart(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	assert.EqualValues(t, &expect, cart)

	//query error
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id"}).AddRow(1)
			return rr
		}())

	_, err = repo.GetCart(userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestInsertItemIntoCartById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	var orderID int = 1
	var itemID int = 1
	var promo string = "A10zzzzz"
	expectTime := time.Unix(1, 0)
	//updatequery
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())
	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, itemID, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	mock.
		ExpectExec("UPDATE orderItems").
		WithArgs(orderID, itemID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &ProductStore{
		db: db,
	}

	err = repo.InsertItemIntoCartById(userID, itemID)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//insertquery
	var itemID2 int = 2
	var newCount int = 1
	var itemprice float64 = 10000
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())
	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, itemID, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(itemID2).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(itemID2, "IPhone", "phones", itemprice, 50000, 0, nil)
			return rr
		}())

	mock.
		ExpectExec("INSERT INTO orderItems").
		WithArgs(itemID2, orderID, itemprice, newCount).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.InsertItemIntoCartById(userID, itemID2)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//query error
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id"}).AddRow(1)
			return rr
		}())

	err = repo.InsertItemIntoCartById(userID, itemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDeleteItemFromCartById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	var orderID int = 1
	var itemID int = 1
	var promo string = "A10zzzzz"
	expectTime := time.Unix(1, 0)
	//deletequery
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())
	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, itemID, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	mock.
		ExpectExec("DELETE").
		WithArgs(orderID, itemID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &ProductStore{
		db: db,
	}

	err = repo.DeleteItemFromCartById(userID, itemID)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//multicount-updatequery
	var itemID2 int = 1
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())
	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(2, itemID, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	mock.
		ExpectExec("UPDATE").
		WithArgs(itemID2, orderID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.DeleteItemFromCartById(userID, itemID2)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//query error
	mock.
		ExpectQuery("SELECT ID,").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id"}).AddRow(1)
			return rr
		}())

	err = repo.InsertItemIntoCartById(userID, itemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestUpdateCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var userID int = 1
	var orderID int = 1
	var itemID int = 1
	var promo string = "A10zzzzz"
	itemIDarr := []int{1}
	expectTime := time.Unix(1, 0)
	//deletequery
	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())
	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, itemID, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	mock.
		ExpectExec("DELETE").
		WithArgs(orderID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT ID").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())
	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, itemID, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	repo := &ProductStore{
		db: db,
	}
	err = repo.UpdateCart(userID, &itemIDarr)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//query error
	mock.
		ExpectQuery("SELECT ID, userID, orderStatus, paymentStatus, addressID, paymentcardID, creationDate, deliveryDate").
		WithArgs(userID, "cart").
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id"}).AddRow(1)
			return rr
		}())

	err = repo.InsertItemIntoCartById(userID, itemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
