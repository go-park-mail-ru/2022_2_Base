package repository

// import (
// 	"reflect"
// 	"serv/domain/model"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// )

// func TestGetProductsRatingAndCommsCountFromStore(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}
// 	defer db.Close()

// 	var itemID int = 1
// 	var expRating float64 = 3
// 	var expCommsCount int = 5

// 	// good query
// 	mock.ExpectQuery("SELECT").
// 		WithArgs(itemID).
// 		WillReturnRows(func() *sqlmock.Rows {
// 			rr := sqlmock.NewRows([]string{"commsCount", "rating"}).AddRow(expCommsCount, expRating)
// 			return rr
// 		}())

// 	repo := &ProductStore{
// 		db: db,
// 	}
// 	rating, commsCount, err := repo.GetProductsRatingAndCommsCountFromStore(itemID)
// 	if err != nil {
// 		t.Errorf("unexpected err: %s", err)
// 		return
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// 	if !reflect.DeepEqual(rating, expRating) {
// 		t.Errorf("results not match, want %v, have %v", expRating, rating)
// 		return
// 	}
// 	if !reflect.DeepEqual(commsCount, expCommsCount) {
// 		t.Errorf("results not match, want %v, have %v", expCommsCount, commsCount)
// 		return
// 	}

// 	// row scan error
// 	mock.
// 		ExpectQuery("SELECT").
// 		WithArgs(itemID).
// 		WillReturnRows(func() *sqlmock.Rows {
// 			rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "asd")
// 			return rr
// 		}())

// 	_, _, err = repo.GetProductsRatingAndCommsCountFromStore(itemID)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// 	if err == nil {
// 		t.Errorf("expected error, got nil")
// 		return
// 	}
// }

// func TestGetCommentsFromStore(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}
// 	defer db.Close()

// 	var itemID int = 1
// 	var userID int = 1
// 	var prodID int = 1

// 	expect := []*model.CommentDB{
// 		{UserID: 1, Pros: "default", Cons: "default", Comment: "default", Rating: 0},
// 	}

// 	// good query
// 	mock.ExpectQuery("SELECT").
// 		WithArgs(itemID).
// 		WillReturnRows(func() *sqlmock.Rows {
// 			rr := sqlmock.NewRows([]string{"userID", "pros", "cons", "comment", "rating"}).AddRow(userID, "default", "default", "default", 0)
// 			return rr
// 		}())

// 	repo := &ProductStore{
// 		db: db,
// 	}
// 	comments, err := repo.GetCommentsFromStore(prodID)
// 	if err != nil {
// 		t.Errorf("unexpected err: %s", err)
// 		return
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// 	if !reflect.DeepEqual(comments, expect) {
// 		t.Errorf("results not match, want %v, have %v", expect, comments)
// 		return
// 	}

// 	// row scan error
// 	mock.
// 		ExpectQuery("SELECT").
// 		WithArgs(itemID).
// 		WillReturnRows(func() *sqlmock.Rows {
// 			rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "asd")
// 			return rr
// 		}())

// 	_, err = repo.GetCommentsFromStore(prodID)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// 	if err == nil {
// 		t.Errorf("expected error, got nil")
// 		return
// 	}
// }

// func TestCreateCommentInStore(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}
// 	defer db.Close()

// 	var itemID int = 1
// 	var userID int = 1
// 	comm := &model.CreateComment{ItemID: itemID, UserID: userID, Pros: "default", Cons: "default", Comment: "default", Rating: 0}

// 	// good query
// 	mock.
// 		ExpectExec("INSERT INTO").
// 		WithArgs(comm.ItemID, comm.UserID, comm.Pros, comm.Cons, comm.Comment, comm.Rating).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	repo := &ProductStore{
// 		db: db,
// 	}
// 	err = repo.CreateCommentInStore(comm)
// 	if err != nil {
// 		t.Errorf("unexpected err: %s", err)
// 		return
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// }

// func TestUpdateProductRatingInStore(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("cant create mock: %s", err)
// 	}
// 	defer db.Close()

// 	var itemID int = 1
// 	var expRating float64 = 3
// 	var expCommsCount int = 5

// 	// good query
// 	mock.ExpectQuery("SELECT").
// 		WithArgs(itemID).
// 		WillReturnRows(func() *sqlmock.Rows {
// 			rr := sqlmock.NewRows([]string{"commsCount", "rating"}).AddRow(expCommsCount, expRating)
// 			return rr
// 		}())

// 	mock.
// 		ExpectExec("UPDATE").
// 		WithArgs(expRating, itemID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	repo := &ProductStore{
// 		db: db,
// 	}
// 	err = repo.UpdateProductRatingInStore(itemID)
// 	if err != nil {
// 		t.Errorf("unexpected err: %s", err)
// 		return
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 		return
// 	}
// }
