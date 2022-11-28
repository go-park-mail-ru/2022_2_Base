package usecase

import (
	"math"
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
	for _, product := range products {
		rating, commsCount, err := api.store.GetProductsRatingAndCommsCountFromStore(product.ID)
		if err != nil {
			return nil, err
		}
		product.Rating = math.Round(rating*100) / 100
		product.CommentsCount = &commsCount
	}
	return products, nil
}

func (api *ProductUsecase) GetProductsWithCategory(cat string, lastitemid int, count int, sort string) ([]*model.Product, error) {
	products, err := api.store.GetProductsWithCategoryFromStore(cat, lastitemid, count, sort)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		rating, commsCount, err := api.store.GetProductsRatingAndCommsCountFromStore(product.ID)
		if err != nil {
			return nil, err
		}
		product.Rating = math.Round(rating*100) / 100
		product.CommentsCount = &commsCount
	}
	return products, nil
}

func (api *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	product, err := api.store.GetProductFromStoreByID(id)
	if err != nil {
		return nil, err
	}
	rating, commsCount, err := api.store.GetProductsRatingAndCommsCountFromStore(product.ID)
	if err != nil {
		return nil, err
	}
	product.Rating = math.Round(rating*100) / 100
	product.CommentsCount = &commsCount
	return product, nil
}

func (api *ProductUsecase) GetProductsBySearch(search string) ([]*model.Product, error) {
	products, err := api.store.GetProductsBySearchFromStore(search)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		rating, commsCount, err := api.store.GetProductsRatingAndCommsCountFromStore(product.ID)
		if err != nil {
			return nil, err
		}
		product.Rating = math.Round(rating*100) / 100
		product.CommentsCount = &commsCount
	}
	return products, nil
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
			}
		}
	}

	err = api.store.MakeOrder(in, &boughtItemsIDs)
	if err != nil {
		return err
	}

	err = api.store.CreateCart(in.UserID)
	if err != nil {
		return err
	}

	return api.store.UpdateCart(in.UserID, &remainedItemsIDs)
}

func (api *ProductUsecase) GetOrders(userID int) ([]*model.Order, error) {
	orders, err := api.store.GetOrdersFromStore(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (api *ProductUsecase) GetOrdersAddress(addressID int) (model.Address, error) {
	address, err := api.store.GetOrdersAddressFromStore(addressID)
	if err != nil {
		return model.Address{}, err
	}
	return *address, nil
}

func (api *ProductUsecase) GetOrdersPayment(paymentID int) (model.PaymentMethod, error) {
	payment, err := api.store.GetOrdersPaymentFromStore(paymentID)
	if err != nil {
		return model.PaymentMethod{}, err
	}
	return *payment, nil
}

func (api *ProductUsecase) GetComments(productID int) ([]*model.CommentDB, error) {
	comments, err := api.store.GetCommentsFromStore(productID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (api *ProductUsecase) CreateComment(in *model.CreateComment) error {
	return api.store.CreateCommentInStore(in)
}
