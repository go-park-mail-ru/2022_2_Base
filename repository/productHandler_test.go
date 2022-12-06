package repository

import (
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
	//id, name, category, price, nominalprice, rating, imgsrc
	rows := sqlmock.
		NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"})
		//NewRows([]string{"id", "title", "updated", "description"})
	expect := []*model.Product{
		{ID: itemID, Name: "IPhone", Category: "phones", Price: 50000, NominalPrice: 50000, Rating: 0, Imgsrc: nil, CommentsCount: nil},
		//{itemID, "IPhone", "phones", 50000, 50000, 0, sql.NullString{}, nil},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Name, item.Category, item.Price, item.NominalPrice, item.Rating, item.Imgsrc)
	}

	mock.
		ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").
		WithArgs(itemID).
		WillReturnRows(rows)

	repo := &ProductStoreDB{
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
	// 	ExpectQuery("SELECT id, title, updated, description FROM items WHERE").
	// 	WithArgs(elemID).
	// 	WillReturnError(fmt.Errorf("db_error"))

	// _, err = repo.SelectByID(elemID)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// 	return
	// }
	// if err == nil {
	// 	t.Errorf("expected error, got nil")
	// 	return
	// }

	// // row scan error
	// rows = sqlmock.NewRows([]string{"id", "title"}).
	// 	AddRow(1, "title")

	// mock.
	// 	ExpectQuery("SELECT id, title, updated, description FROM items WHERE").
	// 	WithArgs(elemID).
	// 	WillReturnRows(rows)

	// _, err = repo.SelectByID(elemID)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// 	return
	// }
	// if err == nil {
	// 	t.Errorf("expected error, got nil")
	// 	return
	// }

}

// func TestShouldGetPosts(t *testing.T) {
// 	mock, err := pgxmock.NewConn()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer mock.Close(context.Background())

// 	// create app with mocked db, request and response to test
// 	app := &api{mock}
// 	req, err := http.NewRequest("GET", "http://localhost/posts", nil)
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected while creating request", err)
// 	}
// 	w := httptest.NewRecorder()

// 	// before we actually execute our api function, we need to expect required DB actions
// 	rows := mock.NewRows([]string{"id", "title", "body"}).
// 		AddRow(1, "post 1", "hello").
// 		AddRow(2, "post 2", "world")

// 	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

// 	// now we execute our request
// 	app.posts(w, req)

// 	if w.Code != 200 {
// 		t.Fatalf("expected status code to be 200, but got: %d\nBody: %v", w.Code, w.Body)
// 	}

// 	data := struct {
// 		Posts []*post
// 	}{Posts: []*post{
// 		{ID: 1, Title: "post 1", Body: "hello"},
// 		{ID: 2, Title: "post 2", Body: "world"},
// 	}}
// 	app.assertJSON(w.Body.Bytes(), data, t)

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestName(t *testing.T) {
// 	t.Parallel()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// // given
// 	// mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
// 	// columns := []string{"id", "price"}
// 	// pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
// 	// mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
// 	// orderDao := &NewMockProductStoreInterface{
// 	// 	db: mockPool,
// 	// }
// 	orderDao := mocks.NewMockProductStoreInterface(ctrl)

// 	// given
// 	//mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
// 	//columns := []string{"id", "price"}
// 	//pgxRows := pgxpoolmock.NewRows(columns).AddRow(100, 100000.9).ToPgxRows()
// 	//mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

// 	// when
// 	actualOrder, err := orderDao.GetProductFromStoreByID(1)
// 	if err != nil {
// 		assert.NotNil(t, actualOrder)
// 	}

// 	// then
// 	assert.NotNil(t, actualOrder)
// 	assert.Equal(t, 100, actualOrder.ID)
// 	assert.Equal(t, 100000.9, actualOrder.Price)
// }

// var cases = []struct {
// 	text []*model.Product
// 	want []*model.Product
// 	err  error
// }{
// 	{[]*model.Product{
// 		{ID: 0, Name: "Монитор Xiaomi Mi 27", Description: "good", Price: 14999, DiscountPrice: 13999, Rating: 4, Imgsrc: "https://img.mvideo.ru/Big/30058309bb.jpg"},
// 		{ID: 1, Name: "Телевизор Haier 55", Description: "goood", Price: 59999, DiscountPrice: 41999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/10030234bb.jpg"},
// 		{ID: 2, Name: "Apple iPad 10.2", Description: "old", Price: 49999, DiscountPrice: 49999, Rating: 3.7, Imgsrc: "https://img.mvideo.ru/Pdb/30064043b.jpg"},
// 		{ID: 3, Name: "Tecno Spark 8с", Description: "good phone", Price: 12999, DiscountPrice: 8999, Rating: 4.5, Imgsrc: "https://img.mvideo.ru/Big/30062036bb.jpg"},
// 		{ID: 4, Name: "realme GT Master", Description: "goood", Price: 29999, DiscountPrice: 21999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/30058843bb.jpg"},
// 		{ID: 5, Name: "Apple iPhone 11", Description: "old", Price: 62999, DiscountPrice: 54999, Rating: 5, Imgsrc: "https://img.mvideo.ru/Big/30063237bb.jpg"},
// 	}, []*model.Product{
// 		{ID: 0, Name: "Монитор Xiaomi Mi 27", Description: "good", Price: 14999, DiscountPrice: 13999, Rating: 4, Imgsrc: "https://img.mvideo.ru/Big/30058309bb.jpg"},
// 		{ID: 1, Name: "Телевизор Haier 55", Description: "goood", Price: 59999, DiscountPrice: 41999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/10030234bb.jpg"},
// 		{ID: 2, Name: "Apple iPad 10.2", Description: "old", Price: 49999, DiscountPrice: 49999, Rating: 3.7, Imgsrc: "https://img.mvideo.ru/Pdb/30064043b.jpg"},
// 		{ID: 3, Name: "Tecno Spark 8с", Description: "good phone", Price: 12999, DiscountPrice: 8999, Rating: 4.5, Imgsrc: "https://img.mvideo.ru/Big/30062036bb.jpg"},
// 		{ID: 4, Name: "realme GT Master", Description: "goood", Price: 29999, DiscountPrice: 21999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/30058843bb.jpg"},
// 		{ID: 5, Name: "Apple iPhone 11", Description: "old", Price: 62999, DiscountPrice: 54999, Rating: 5, Imgsrc: "https://img.mvideo.ru/Big/30063237bb.jpg"},
// 	}, nil},
// }

// func TestGetProducts(t *testing.T) {
// 	for _, c := range cases {
// 		t.Run("tests", func(t *testing.T) {

// 			productHandler := usecase.NewProductHandler()
// 			got, err := productHandler.GetProducts()
// 			if err != nil {
// 				t.Errorf(err.Error())
// 			}
// 			assert.Equal(t, got, c.want)
// 		})
// 	}
// }

// func (ps *ProductStore) GetProductsFromStore500() ([]*model.Product, error) {
// 	return nil, baseErrors.ErrServerError500
// }

// func TestGetProducts500(t *testing.T) {

// 	t.Run("tests", func(t *testing.T) {

// 		productHandler := usecase.NewProductHandler()
// 		_, err := productHandler.store.GetProductsFromStore500()

// 		assert.ErrorIs(t, err, baseErrors.ErrServerError500)
// 	})

// }

// func TestShouldUpdateStats(t *testing.T) {
// 	mock, err := pgxmock.NewConn()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer mock.Close(context.Background())

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE products").WillReturnResult(pgxmock.NewResult("UPDATE", 1))
// 	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(pgxmock.NewResult("INSERT", 1))
// 	mock.ExpectCommit()

// 	// now we execute our method
// 	if err = recordStats(mock, 2, 3); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
