package handlers

import (
	"serv/model"
	"sync"
)

type ProductStore struct {
	products []*model.Product
	mu       sync.RWMutex
	nextID   uint
}

func NewProductStore() *ProductStore {

	return &ProductStore{
		mu: sync.RWMutex{},
		products: []*model.Product{
			{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: "https://s.ek.ua/jpg_zoom1/2090045.jpg"},
			{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: "https://fainaidea.com/wp-content/uploads/2016/11/324987.jpg"},
			{ID: 2, Name: "ipad", Description: "old", Price: 3000, DiscountPrice: 3000, Rating: 1, Imgsrc: "https://cdn-files.kimovil.com/default/0005/17/thumb_416447_default_big.jpeg"},
		},
	}
}

func (ps *ProductStore) GetProductsFromStore() ([]*model.Product, error) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	products := ps.products

	return products, nil
}
