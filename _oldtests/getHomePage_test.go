package delivery

import (
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"serv/model"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHomePage2(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("GET", conf.PathMain, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		productHandler := NewProductHandler()
		productHandler.GetHomePage(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

		var expected = `{"body":[{"id":0,"name":"Монитор Xiaomi Mi 27","description":"good","price":14999,"lowprice":13999,"rating":4,"imgsrc":"https://img.mvideo.ru/Big/30058309bb.jpg"},{"id":1,"name":"Телевизор Haier 55","description":"goood","price":59999,"lowprice":41999,"rating":4.3,"imgsrc":"https://img.mvideo.ru/Big/10030234bb.jpg"},{"id":2,"name":"Apple iPad 10.2","description":"old","price":49999,"lowprice":49999,"rating":3.7,"imgsrc":"https://img.mvideo.ru/Pdb/30064043b.jpg"},{"id":3,"name":"Tecno Spark 8с","description":"good phone","price":12999,"lowprice":8999,"rating":4.5,"imgsrc":"https://img.mvideo.ru/Big/30062036bb.jpg"},{"id":4,"name":"realme GT Master","description":"goood","price":29999,"lowprice":21999,"rating":4.3,"imgsrc":"https://img.mvideo.ru/Big/30058843bb.jpg"},{"id":5,"name":"Apple iPhone 11","description":"old","price":62999,"lowprice":54999,"rating":5,"imgsrc":"https://img.mvideo.ru/Big/30063237bb.jpg"}]}`
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
		productHandler.GetHomePage(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
