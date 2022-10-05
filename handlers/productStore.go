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
			{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: ""},
			{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: ""},
			{ID: 2, Name: "key", Description: "fake", Price: 23, DiscountPrice: 1, Rating: 1, Imgsrc: ""},
		},
	}
}

func (ps *ProductStore) GetProductsFromStore() ([]*model.Product, error) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	products := ps.products

	return products, nil
}
