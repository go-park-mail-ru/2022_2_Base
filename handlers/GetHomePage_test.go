package handlers

import (
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"serv/model"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHomePage(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("GET", conf.PathMain, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		productHandler := NewProductHandler()
		handler := http.HandlerFunc(productHandler.GetHomePage)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

		expected := `{"body":[{"id":0,"name":"phone","description":"good phone","price":10000,"lowprice":8000,"rating":5,"imgsrc":""},{"id":1,"name":"notebook","description":"goood","price":70000,"lowprice":55000,"rating":4.3,"imgsrc":""},{"id":2,"name":"key","description":"fake","price":23,"lowprice":1,"rating":1,"imgsrc":""}]}`
		assert.Equal(t, rr.Body.String(), expected+"\n")
	})
}

func MockProductStore() *ProductStore {
	return &ProductStore{
		mu:       sync.RWMutex{},
		products: []*model.Product{},
	}
}

func TestGetHomePageError404(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("GET", conf.PathMain, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		productHandler := ProductHandler{store: *MockProductStore()}

		handler := http.HandlerFunc(productHandler.GetHomePage)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)

	})
}
