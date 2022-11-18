package usecase

import (
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

func (api *ProductUsecase) GetProducts(lastitemid int, count int, sort string) ([]*model.Product, error) {
	products, err := api.store.GetProductsFromStore(lastitemid, count, sort)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (api *ProductUsecase) GetProductsWithCategory(cat string, lastitemid int, count int, sort string) ([]*model.Product, error) {
	products, err := api.store.GetProductsWithCategoryFromStore(cat, lastitemid, count, sort)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (api *ProductUsecase) GetCart(userID int) (*model.Order, error) {
	cart, err := api.store.GetCart(userID)
	if err != nil {
		return nil, err
	}
	if cart == nil || cart.ID == 0 {
		err = api.store.CreateCart(userID)
		if err != nil {
			return nil, err
		}
		cart, err = api.store.GetCart(userID)
		if err != nil {
			return nil, err
		}
	}
	return cart, nil
}

func (api *ProductUsecase) UpdateOrder(userID int, items *[]int) error {
	return api.store.UpdateCart(userID, items)
}

func (api *ProductUsecase) AddToOrder(userID int, itemID int) error {
	return api.store.InsertItemIntoCartById(userID, itemID)
}

func (api *ProductUsecase) DeleteFromOrder(userID int, itemID int) error {
	return api.store.DeleteItemFromCartById(userID, itemID)
}

func (api *ProductUsecase) MakeOrder(in *model.MakeOrder) error {
	cart, err := api.store.GetCart(in.UserID)
	if err != nil {
		return err
	}
	remainedItemsIDs := []int{}
	for _, orderItem := range cart.Items {
		flag := true
		for _, id := range in.Items {
			if orderItem.Item.ID == id {
				flag = false
			}
		}
		if flag {
			for i := 0; i < orderItem.Count; i++ {
				remainedItemsIDs = append(remainedItemsIDs, orderItem.Item.ID)
			}

		}
	}

	err = api.store.MakeOrder(in)
	if err != nil {
		return err
	}

	err = api.store.CreateCart(in.UserID)
	if err != nil {
		return err
	}

	return api.store.UpdateCart(in.UserID, &remainedItemsIDs)
}
