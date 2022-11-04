package usecase

import (
	"database/sql"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	rep "serv/repository"
)

type ProductUsecase struct {
	store rep.ProductStore
}

func NewProductUsecase(db *sql.DB) *ProductUsecase {
	return &ProductUsecase{
		store: *rep.NewProductStore(db),
	}
}

func (api *ProductUsecase) GetProducts() ([]*model.Product, error) {

	products, err := api.store.GetProductsFromStore()
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}

	return products, nil
}
