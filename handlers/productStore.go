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
			{ID: 0, Name: "Монитор Xiaomi Mi Desktop Monitor 27'' RMMNT27NF (BHR4975EU)", Description: "good", Price: 14999, DiscountPrice: 13999, Rating: 4, Imgsrc: "https://img.mvideo.ru/Big/30058309bb.jpg"},
			{ID: 1, Name: "Телевизор Haier 55 Smart TV MX", Description: "goood", Price: 59999, DiscountPrice: 41999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/10030234bb.jpg"},
			{ID: 2, Name: "Планшет Apple iPad 10.2 Wi-Fi+Cell 64GB Space Grey (MK473)", Description: "old", Price: 49999, DiscountPrice: 49999, Rating: 3.7, Imgsrc: "https://img.mvideo.ru/Pdb/30064043b.jpg"},
			{ID: 3, Name: "Смартфон Tecno Spark 8с 4+64GB Magnet Black", Description: "good phone", Price: 12999, DiscountPrice: 8999, Rating: 4.5, Imgsrc: "https://img.mvideo.ru/Big/30062036bb.jpg"},
			{ID: 4, Name: "Смартфон realme GT Master Edition 6+128GB Voyager Grey (RMX3363)", Description: "goood", Price: 29999, DiscountPrice: 21999, Rating: 4.3, Imgsrc: "https://img.mvideo.ru/Big/30058843bb.jpg"},
			{ID: 5, Name: "Смартфон Apple iPhone 11 128GB White", Description: "old", Price: 62999, DiscountPrice: 54999, Rating: 5, Imgsrc: "https://img.mvideo.ru/Big/30063237bb.jpg"},
		},
	}
}

func (ps *ProductStore) GetProductsFromStore() ([]*model.Product, error) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	products := ps.products

	return products, nil
}
