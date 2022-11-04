package usecase

import (
	"database/sql"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	rep "serv/repository"
)

type ProductHandler struct {
	store rep.ProductStore
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		store: *rep.NewProductStore(db),
	}
}

func (api *ProductHandler) GetProducts() ([]*model.Product, error) {

	products, err := api.store.GetProductsFromStore()
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}

	return products, nil
}
