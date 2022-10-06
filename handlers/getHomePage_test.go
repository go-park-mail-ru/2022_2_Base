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

		expected := `{"body":[{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: "https://s.ek.ua/jpg_zoom1/2090045.jpg"},
		{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: "https://fainaidea.com/wp-content/uploads/2016/11/324987.jpg"},
		{ID: 2, Name: "ipad", Description: "old", Price: 3000, DiscountPrice: 3000, Rating: 1, Imgsrc: "https://cdn-files.kimovil.com/default/0005/17/thumb_416447_default_big.jpeg"},]}`
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
