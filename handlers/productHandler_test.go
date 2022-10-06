package handlers

import (
	baseErrors "serv/errors"
	"serv/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	text []*model.Product
	want []*model.Product
	err  error
}{
	{[]*model.Product{
		{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: "https://s.ek.ua/jpg_zoom1/2090045.jpg"},
		{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: "https://fainaidea.com/wp-content/uploads/2016/11/324987.jpg"},
		{ID: 2, Name: "ipad", Description: "old", Price: 3000, DiscountPrice: 3000, Rating: 1, Imgsrc: "https://cdn-files.kimovil.com/default/0005/17/thumb_416447_default_big.jpeg"},
	}, []*model.Product{
		{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: "https://s.ek.ua/jpg_zoom1/2090045.jpg"},
		{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: "https://fainaidea.com/wp-content/uploads/2016/11/324987.jpg"},
		{ID: 2, Name: "ipad", Description: "old", Price: 3000, DiscountPrice: 3000, Rating: 1, Imgsrc: "https://cdn-files.kimovil.com/default/0005/17/thumb_416447_default_big.jpeg"},
	}, nil},
}

func TestGetProducts(t *testing.T) {
	for _, c := range cases {
		t.Run("tests", func(t *testing.T) {

			productHandler := NewProductHandler()
			got, err := productHandler.GetProducts()
			if err != nil {
				t.Errorf(err.Error())
			}
			assert.Equal(t, got, c.want)
		})
	}
}

func (ps *ProductStore) GetProductsFromStore500() ([]*model.Product, error) {
	return nil, baseErrors.ErrServerError500
}

func TestGetProducts500(t *testing.T) {

	t.Run("tests", func(t *testing.T) {

		productHandler := NewProductHandler()
		_, err := productHandler.store.GetProductsFromStore500()

		assert.ErrorIs(t, err, baseErrors.ErrServerError500)
	})

}
