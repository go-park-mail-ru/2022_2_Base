package handlers

import (
	baseErrors "serv/errors"
	"serv/model"
)

type ProductHandler struct {
	store ProductStore
}

func (api *ProductHandler) GetProducts() ([]*model.Product, error) {

	products, err := api.store.GetProductsFromStore()
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}

	return products, nil
}
