package repository

import (
	"serv/domain/model"
)

var cases = []struct {
	text []*model.Product
	want []*model.Product
	err  error
}{
	{[]*model.Product{
		{ID: 0, Name: "Монитор Xiaomi Mi 27", Description: "good", Price: 14999, DiscountPrice: 13999, Rating: 4, Imgsrc: "https://img.mvideo.ru/Big/30058309bb.jpg"},
		{ID: 1, Name: "Телевизор Haier 55", Description: "goood", Price: 59999, DiscountPrice: 41999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/10030234bb.jpg"},
		{ID: 2, Name: "Apple iPad 10.2", Description: "old", Price: 49999, DiscountPrice: 49999, Rating: 3.7, Imgsrc: "https://img.mvideo.ru/Pdb/30064043b.jpg"},
		{ID: 3, Name: "Tecno Spark 8с", Description: "good phone", Price: 12999, DiscountPrice: 8999, Rating: 4.5, Imgsrc: "https://img.mvideo.ru/Big/30062036bb.jpg"},
		{ID: 4, Name: "realme GT Master", Description: "goood", Price: 29999, DiscountPrice: 21999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/30058843bb.jpg"},
		{ID: 5, Name: "Apple iPhone 11", Description: "old", Price: 62999, DiscountPrice: 54999, Rating: 5, Imgsrc: "https://img.mvideo.ru/Big/30063237bb.jpg"},
	}, []*model.Product{
		{ID: 0, Name: "Монитор Xiaomi Mi 27", Description: "good", Price: 14999, DiscountPrice: 13999, Rating: 4, Imgsrc: "https://img.mvideo.ru/Big/30058309bb.jpg"},
		{ID: 1, Name: "Телевизор Haier 55", Description: "goood", Price: 59999, DiscountPrice: 41999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/10030234bb.jpg"},
		{ID: 2, Name: "Apple iPad 10.2", Description: "old", Price: 49999, DiscountPrice: 49999, Rating: 3.7, Imgsrc: "https://img.mvideo.ru/Pdb/30064043b.jpg"},
		{ID: 3, Name: "Tecno Spark 8с", Description: "good phone", Price: 12999, DiscountPrice: 8999, Rating: 4.5, Imgsrc: "https://img.mvideo.ru/Big/30062036bb.jpg"},
		{ID: 4, Name: "realme GT Master", Description: "goood", Price: 29999, DiscountPrice: 21999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/30058843bb.jpg"},
		{ID: 5, Name: "Apple iPhone 11", Description: "old", Price: 62999, DiscountPrice: 54999, Rating: 5, Imgsrc: "https://img.mvideo.ru/Big/30063237bb.jpg"},
	}, nil},
}

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
