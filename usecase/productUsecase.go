package usecase

import (
	"context"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	rep "serv/repository"
)

type ProductUsecase struct {
	ordersManager orders.OrdersWorkerClient
	store         rep.ProductStore
}

func NewProductUsecase(ps *rep.ProductStore, ordersManager *orders.OrdersWorkerClient) *ProductUsecase {
	return &ProductUsecase{
		ordersManager: *ordersManager,
		store:         *ps,
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

func (api *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	return api.store.GetProductFromStoreByID(id)
}

func (api *ProductUsecase) GetProductsBySearch(search string) ([]*model.Product, error) {
	return api.store.GetProductsBySearchFromStore(search)
}

func (api *ProductUsecase) GetSuggestions(search string) ([]string, error) {
	return api.store.GetSuggestionsFromStore(search)
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
	boughtItemsIDs := []int{}
	boughtItemsIDsINT32 := []int32{}
	for _, orderItem := range cart.Items {
		flag := true
		flag2 := false
		for _, id := range in.Items {
			if orderItem.Item.ID == id {
				flag = false
				flag2 = true
			}
		}
		if flag {
			for i := 0; i < orderItem.Count; i++ {
				remainedItemsIDs = append(remainedItemsIDs, orderItem.Item.ID)
			}
		}
		if flag2 {
			for i := 0; i < orderItem.Count; i++ {
				boughtItemsIDs = append(boughtItemsIDs, orderItem.Item.ID)
				boughtItemsIDsINT32 = append(boughtItemsIDsINT32, int32(orderItem.Item.ID))
			}
		}
	}

	err = api.store.UpdateCart(in.UserID, &boughtItemsIDs)
	if err != nil {
		return err
	}

	_, err = api.ordersManager.MakeOrder(
		context.Background(),
		&orders.MakeOrderType{
			UserID:        int32(in.UserID),
			Items:         boughtItemsIDsINT32,
			AddressID:     int32(in.AddressID),
			PaymentcardID: int32(in.PaymentcardID),
			DeliveryDate:  in.DeliveryDate.Unix(),
		})
	if err != nil {
		return err
	}

	err = api.store.CreateCart(in.UserID)
	if err != nil {
		return err
	}

	return api.store.UpdateCart(in.UserID, &remainedItemsIDs)
}

func (api *ProductUsecase) GetOrders(userID int) (*orders.OrdersResponse, error) {
	ordersResponse, err := api.ordersManager.GetOrders(
		context.Background(),
		&orders.UserID{
			UserID: int32(userID),
		})
	if err != nil {
		return nil, err
	}
	return ordersResponse, nil
}

func (api *ProductUsecase) GetComments(productID int) ([]*model.Comment, error) {
	comments, err := api.store.GetCommentsFromStore(productID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (api *ProductUsecase) CreateComment(in *model.Comment) error {
	return api.store.CreateCommentInStore(in)
}
