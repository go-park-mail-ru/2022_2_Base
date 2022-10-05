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
		{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: ""},
		{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: ""},
		{ID: 2, Name: "key", Description: "fake", Price: 23, DiscountPrice: 1, Rating: 1, Imgsrc: ""},
	}, []*model.Product{
		{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: ""},
		{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: ""},
		{ID: 2, Name: "key", Description: "fake", Price: 23, DiscountPrice: 1, Rating: 1, Imgsrc: ""},
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
