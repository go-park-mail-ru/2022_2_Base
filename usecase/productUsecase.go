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

func (api *ProductUsecase) GetCart(userID int) (*model.Order, error) {
	cart, err := api.store.GetCart(userID)

	if cart == nil || cart.ID == 0 {
		err = api.store.CreateCart(userID)
		if err != nil {
			return nil, err
		}
	}
	cart, err = api.store.GetCart(userID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (api *ProductUsecase) UpdateOrder(userID int, items *[]int) error {
	return api.store.UpdateCart(userID, items)
}

func (api *ProductUsecase) MakeOrder(userID int) error {

	err := api.store.MakeOrder(userID)
	if err != nil {
		return err
	}
	return api.store.CreateCart(userID)
}
