package repository

import (
	"database/sql/driver"
	"reflect"
	"serv/domain/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	// row scan error
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

func TestGetProductsFromStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	expect := []*model.Product{
		{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
	}

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(driver.Value(0)).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(0, "", "", 0, 0, 0, nil)
			return rr
		}())
	mock.ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").WithArgs(driver.Value(0), driver.Value(1)).WillReturnRows(func() *sqlmock.Rows {
		rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
		return rr
	}())

	repo := &ProductStore{
		db: db,
	}
	var lastItemID int = 0
	var count int = 1
	items, err := repo.GetProductsFromStore(lastItemID, count, "")
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
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(driver.Value(0)).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(0, "")
			return rr
		}())

	_, err = repo.GetProductsFromStore(lastItemID, count, "")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

var casesGetProductsFromStoreSortings = []struct {
	lastItemID int
	count      int
	sort       string
	lastprice  float64
	lastrating float64
}{
	{0, 1, "priceup", 0, 0},
	{0, 1, "pricedown", 1e10, 0},
	{0, 1, "ratingdown", 0, 0},
	{1e9, 1, "ratingup", 0, 10},
}

func TestGetProductsFromStoreSortings(t *testing.T) {
	for _, c := range casesGetProductsFromStoreSortings {
		t.Run("tests", func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("cant create mock: %s", err)
			}
			defer db.Close()

			expect := []*model.Product{
				{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
			}

			var lastItemID int = c.lastItemID
			var count int = c.count
			var sort string = c.sort

			var lastparam float64 = 0
			if c.lastprice > 0 {
				lastparam = c.lastprice
			} else if c.lastrating > 0 {
				lastparam = c.lastrating
			}

			mock.
				ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
				WithArgs(driver.Value(0)).
				WillReturnRows(func() *sqlmock.Rows {
					rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(0, "", "", 0, 0, 0, nil)
					return rr
				}())
			mock.ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").WithArgs(lastparam, lastItemID, count).WillReturnRows(func() *sqlmock.Rows {
				rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
				return rr
			}())

			repo := &ProductStore{
				db: db,
			}
			items, err := repo.GetProductsFromStore(0, count, sort)
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
		})
	}
}

func TestGetProductsWithCategoryFromStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	expect := []*model.Product{
		{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
	}

	var lastItemID int = 0
	var count int = 1
	var category string = "phones"

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(driver.Value(0)).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(0, "", "", 0, 0, 0, nil)
			return rr
		}())
	mock.ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").WithArgs(category, lastItemID, count).WillReturnRows(func() *sqlmock.Rows {
		rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
		return rr
	}())

	repo := &ProductStore{
		db: db,
	}

	items, err := repo.GetProductsWithCategoryFromStore(category, lastItemID, count, "")
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
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(driver.Value(0)).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(0, "")
			return rr
		}())

	_, err = repo.GetProductsFromStore(lastItemID, count, "")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

var casesGetProductsWithCategoryFromStoreSortings = []struct {
	category   string
	lastItemID int
	count      int
	sort       string
	lastprice  float64
	lastrating float64
}{
	{"phones", 0, 1, "priceup", 0, 0},
	{"phones", 0, 1, "pricedown", 1e10, 0},
	{"phones", 0, 1, "ratingdown", 0, 0},
	{"phones", 1e9, 1, "ratingup", 0, 10},
}

func TestGetProductsWithCategoryFromStoreSortings(t *testing.T) {
	for _, c := range casesGetProductsWithCategoryFromStoreSortings {
		t.Run("tests", func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("cant create mock: %s", err)
			}
			defer db.Close()

			expect := []*model.Product{
				{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
			}

			var lastItemID int = c.lastItemID
			var count int = c.count
			var sort string = c.sort
			var category string = c.category
			var lastparam float64 = 0
			if c.lastprice > 0 {
				lastparam = c.lastprice
			} else if c.lastrating > 0 {
				lastparam = c.lastrating
			}

			mock.
				ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id").
				WithArgs(driver.Value(0)).
				WillReturnRows(func() *sqlmock.Rows {
					rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(0, "", "", 0, 0, 0, nil)
					return rr
				}())
			mock.ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").WithArgs(category, lastparam, lastItemID, count).WillReturnRows(func() *sqlmock.Rows {
				rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
				return rr
			}())

			repo := &ProductStore{
				db: db,
			}
			items, err := repo.GetProductsWithCategoryFromStore(category, 0, count, sort)
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
		})
	}
}
