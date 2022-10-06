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
			{ID: 0, Name: "phone", Description: "good phone", Price: 10000, DiscountPrice: 8000, Rating: 5, Imgsrc: "https://media.4rgos.it/i/Argos/9520055_R_Z001A?w=1500&h=880&qlt=70&fmt=webp"},
			{ID: 1, Name: "notebook", Description: "goood", Price: 70000, DiscountPrice: 55000, Rating: 4.3, Imgsrc: "https://www.notebookcheck-ru.com/uploads/tx_nbc2/MicrosoftSurfaceLaptop3-15__1_.JPG"},
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
