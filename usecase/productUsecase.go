package usecase

import (
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	rep "serv/repository"
)

type ProductUsecase struct {
	store rep.ProductStore
}

func NewProductUsecase(ps *rep.ProductStore) *ProductUsecase {
	return &ProductUsecase{
		store: *ps,
	}
}

func (api *ProductUsecase) GetProducts() ([]*model.Product, error) {

	products, err := api.store.GetProductsFromStore()
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}

	return products, nil
}
