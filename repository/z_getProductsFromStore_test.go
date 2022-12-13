package repository

import (
	"database/sql/driver"
	"reflect"
	"serv/domain/model"
	"testing"

	//"serv/usecase"

	//"github.com/driftprogramming/pgxpoolmock"

	//mocks "serv/repository/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	//"github.com/pashagolub/pgxmock/v2"
)

func TestGetProductsFromStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	//var itemID int = 1

	// good query
	// rows := sqlmock.
	// 	NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
	expect := []*model.Product{
		{ID: 1, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
	}

	// for _, item := range expect {
	// 	rows = rows.AddRow(item.ID, item.Name, item.Category, item.Price, item.NominalPrice, item.Rating, item.Imgsrc)
	// }

	// mock.
	// 	ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
	// 	WithArgs(itemID).
	// 	WillReturnRows(rows)

	// 	mock.ExpectBegin()
	// 	mock.ExpectQuery(`
	// INSERT INTO users (id, login, password, is_super)
	// VALUES ($1, $2, $3, $4)
	// RETURNING created_at;`).WithArgs(
	// 		driver.Value(1),
	// 		driver.Value("user01"),
	// 		driver.Value("first-P"),
	// 		driver.Value(false),
	// 	).WillReturnError(&pq.Error{Severity: "ERROR", Code: "23505", Message: "duplicate key value violates unique constraint \"users_pkey\"", Detail: "Key (id)=(1) already exists.", Hint: "", Position: "", InternalPosition: "", InternalQuery: "", Where: "", Schema: "public", Table: "users", Column: "", DataTypeName: "", Constraint: "users_pkey", File: "nbtinsert.c", Line: "432", Routine: "_bt_check_unique"})
	// 	mock.ExpectRollback()

	// mock.ExpectQuery(`
	// 	SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products
	// 	WHERE id = $1;`).WithArgs(
	// 	driver.Value(0),
	// ).WillReturnRows(func() *sqlmock.Rows {
	// 	rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(0, "", "", 0, 0, 0, nil)
	// 	return rr
	// }())

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id").
		WithArgs(driver.Value(0)).
		WillReturnRows(func() *sqlmock.Rows {
			rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(0, "", "", 0, 0, 0, nil)
			return rr
		}())

	//mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id >").WithArgs(driver.Value(0), driver.Value(1)).WillReturnRows(func() *sqlmock.Rows {
		rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
		return rr
	}())
	//mock.ExpectRollback()

	repo := &ProductStore{
		db: db,
	}
	var lastItemID int = 0
	var count int = 1
	items, err := repo.GetProductsFromStore(lastItemID, count, "")
	//items, err := repo.GetProductsFromStore(lastItemID, count, "")
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

	// rows = sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "asd")

	// mock.
	// 	ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
	// 	WithArgs(itemID).
	// 	WillReturnRows(rows)

	// _, err = repo.GetProductFromStoreByID(itemID)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// 	return
	// }
	// if err == nil {
	// 	t.Errorf("expected error, got nil")
	// 	return
	// }

}
