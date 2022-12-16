package repository

import (
	"reflect"
	"serv/domain/model"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var casesGetProductsBySearchFromStore = []struct {
	search string
}{
	{"asgh"},
	{"iphone"},
}

func TestGetProductsBySearchFromStore(t *testing.T) {
	for _, c := range casesGetProductsBySearchFromStore {
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

			var search string = c.search
			searchWords := strings.Split(search, " ")
			searchWordsUnite := strings.Join(searchWords, "")
			searchLetters := strings.Split(searchWordsUnite, "")
			searchString := strings.ToLower(`%` + strings.Join(searchLetters, "%") + `%`)

			mock.ExpectQuery("SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE").WithArgs(searchString).WillReturnRows(func() *sqlmock.Rows {
				rr := sqlmock.NewRows([]string{"id", "name", "category", "price", "nominalprice", "rating", "imgsrc"}).AddRow(1, "IPhone", "phones", 50000, 50000, 0, nil)
				return rr
			}())

			repo := &ProductStore{
				db: db,
			}

			items, err := repo.GetProductsBySearchFromStore(search)
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
				WithArgs(searchString).
				WillReturnRows(func() *sqlmock.Rows {
					rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(0, "")
					return rr
				}())

			_, err = repo.GetProductsBySearchFromStore(search)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
				return
			}
			if err == nil {
				t.Errorf("expected error, got nil")
				return
			}
		})
	}
}

func TestGetSuggestionsFromStore(t *testing.T) {
	for _, c := range casesGetProductsBySearchFromStore {
		t.Run("tests", func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("cant create mock: %s", err)
			}
			defer db.Close()

			expect := []string{"IPhone"}

			var search string = c.search
			searchWords := strings.Split(search, " ")
			searchString := strings.ToLower(`%` + strings.Join(searchWords, " ") + `%`)

			mock.ExpectQuery("SELECT name FROM products WHERE").WithArgs(searchString).WillReturnRows(func() *sqlmock.Rows {
				rr := sqlmock.NewRows([]string{"name"}).AddRow("IPhone")
				return rr
			}())

			repo := &ProductStore{
				db: db,
			}

			items, err := repo.GetSuggestionsFromStore(search)
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
				ExpectQuery("SELECT name FROM products WHERE").
				WithArgs(searchString).
				WillReturnRows(func() *sqlmock.Rows {
					rr := sqlmock.NewRows([]string{"id", "name"}).AddRow(0, "")
					return rr
				}())

			_, err = repo.GetSuggestionsFromStore(search)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
				return
			}
			if err == nil {
				t.Errorf("expected error, got nil")
				return
			}
		})
	}
}
