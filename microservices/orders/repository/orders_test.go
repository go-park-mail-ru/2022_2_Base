package orders

import (
	"reflect"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

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
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, 1, "IPhone", "phones", 50000, 50000, 0, nil)
			return rr
		}())

	repo := &OrderStore{
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
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
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

func TestGetOrdersFromStore(t *testing.T) {
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
	expect := []*model.Order{
		{ID: 1, UserID: 1, Items: expectItems, OrderStatus: "cart", PaymentStatus: "not started", AddressID: 1, PaymentcardID: 1, CreationDate: &expectTime, DeliveryDate: &expectTime, Promocode: &promo},
	}
	mock.
		ExpectQuery("SELECT id").
		WithArgs(userID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "userID", "orderStatus", "paymentStatus", "addressID", "paymentcardID", "creationDate", "deliveryDate", "promocode"}).AddRow(1, userID, "cart", "not started", 1, 1, &expectTime, &expectTime, &promo)
			return rr
		}())

	mock.
		ExpectQuery("SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems").
		WithArgs(orderID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"count", "id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, 1, "IPhone", "phones", 50000, 50000, 0, &avatar)
			return rr
		}())

	repo := &OrderStore{
		db: db,
	}

	orders, err := repo.GetOrdersFromStore(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	assert.EqualValues(t, expect, orders)

	//query error
	mock.
		ExpectQuery("SELECT id").
		WithArgs(userID).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id"}).AddRow(1)
			return rr
		}())

	_, err = repo.GetOrdersFromStore(userID)
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
	expect := model.Address{ID: 1, City: "default", Street: "default", House: "default", Flat: "default", Priority: false}

	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "city", "street", "house", "flat", "priority"}).AddRow(1, "default", "default", "default", "default", false)
			return rr
		}())
	repo := &OrderStore{
		db: db,
	}
	adresses, err := repo.GetOrdersAddressFromStore(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(adresses, &expect) {
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
	_, err = repo.GetOrdersAddressFromStore(id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetOrdersPaymentFromStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	var id int = 1
	expectTime := time.Unix(1, 0)
	expect := model.PaymentMethod{ID: 1, PaymentType: "default", Number: "default", ExpiryDate: expectTime, Priority: false}

	mock.
		ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "paymentType", "number", "expiryDate", "priority"}).AddRow(1, "default", "default", expectTime, false)
			return rr
		}())
	repo := &OrderStore{
		db: db,
	}
	paym, err := repo.GetOrdersPaymentFromStore(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(paym, &expect) {
		t.Errorf("results not match, want %v, have %v", expect, paym)
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
	_, err = repo.GetOrdersPaymentFromStore(id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestMakeOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	//var userEmail string = "a@a"
	//var userName string = "art"
	//var userPass string = "12345678"
	in := orders.MakeOrderType{AddressID: 1, PaymentcardID: 1, DeliveryDate: 1, UserID: 1}
	delivDate := time.Unix(in.DeliveryDate, 0)
	mock.
		ExpectExec("UPDATE").
		WithArgs("created", "not started", in.AddressID, in.PaymentcardID, time.Now().Format("2006.01.02 15:04:05"), delivDate.Format("2006.01.02 15:04:05"), in.UserID, "cart").
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := &OrderStore{
		db: db,
	}
	//expectTime := time.Unix(1, 0)

	err = repo.MakeOrder(&in)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
