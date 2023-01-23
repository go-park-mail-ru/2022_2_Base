package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	conf "serv/config"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	mail "serv/microservices/mail/gen_files"
	orders "serv/microservices/orders/gen_files"
	rep "serv/repository"
	"time"
)

type ProductUsecaseInterface interface {
	GetProducts(lastitemid int, count int, sort string, userID int) ([]*model.Product, error)
	GetProductsWithCategory(cat string, lastitemid int, count int, sort string, userID int) ([]*model.Product, error)
	GetProductsWithBiggestDiscount(lastitemid int, count int, userID int) ([]*model.Product, error)
	GetProductByID(id int, userID int) (*model.Product, error)
	GetProductsBySearch(search string, userID int) ([]*model.Product, error)
	GetBestProductInCategory(category string, userID int) (*model.Product, error)
	GetSuggestions(search string) ([]string, error)
	GetCart(userID int) (*model.Order, error)
	UpdateOrder(userID int, items *[]int) error
	AddToOrder(userID int, itemID int) error
	DeleteFromOrder(userID int, itemID int) error
	MakeOrder(in *model.MakeOrder) (int, error)
	ChangeOrderStatus(userID int, in *model.ChangeOrderStatus) error
	GetOrders(userID int) (*orders.OrdersResponse, error)
	GetComments(productID int) ([]*model.CommentDB, error)
	CreateComment(in *model.CreateComment) error
	GetRecommendationProducts(itemID int, userID int) ([]*model.Product, error)
	SetPromocode(userID int, promocode string) error
	RecalculatePrices(userID int, promocode string) error
	GetFavorites(userID int, lastitemid int, count int, sort string) ([]*model.Product, error)
	InsertItemIntoFavorites(userID int, itemID int) error
	DeleteItemFromFavorites(userID int, itemID int) error
	RecalculateRatingsForInitscriptProducts(count int) error
}

type ProductUsecase struct {
	ordersManager orders.OrdersWorkerClient
	mailManager   mail.MailServiceClient
	store         rep.ProductStoreInterface
}

func NewProductUsecase(ps rep.ProductStoreInterface, ordersManager orders.OrdersWorkerClient, mailManager mail.MailServiceClient) ProductUsecaseInterface {
	return &ProductUsecase{
		ordersManager: ordersManager,
		mailManager:   mailManager,
		store:         ps,
	}
}

func (api *ProductUsecase) GetProducts(lastitemid int, count int, sort string, userID int) ([]*model.Product, error) {
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
		product.Properties = []*model.Property{}
		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
	}
	return products, nil
}

func (api *ProductUsecase) GetProductsWithCategory(cat string, lastitemid int, count int, sort string, userID int) ([]*model.Product, error) {
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

		properties, err := api.store.GetProductPropertiesFromStore(product.ID, product.Category)
		if err != nil {
			return nil, err
		}
		product.Properties = properties
		if len(properties) > 4 {
			product.Properties = properties[:4]
		}

		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
	}
	return products, nil
}

func (api *ProductUsecase) GetProductsWithBiggestDiscount(lastitemid int, count int, userID int) ([]*model.Product, error) {
	products, err := api.store.GetProductsWithBiggestDiscountFromStore(lastitemid, count)
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

		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
	}
	return products, nil
}

func (api *ProductUsecase) GetProductByID(id int, userID int) (*model.Product, error) {
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

	properties, err := api.store.GetProductPropertiesFromStore(product.ID, product.Category)
	if err != nil {
		return nil, err
	}
	product.Properties = properties
	if userID != 0 {
		isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
		if err != nil {
			return nil, err
		}
		product.IsFavorite = isFav
	}

	return product, nil
}

func (api *ProductUsecase) GetBestProductInCategory(category string, userID int) (*model.Product, error) {
	products, err := api.store.GetProductsWithCategoryFromStore(category, 0, 10, "ratingdown")
	if err != nil {
		return nil, err
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randitemID := r1.Intn(int(math.Min(10, float64(len(products)))))
	product := products[randitemID]
	rating, commsCount, err := api.store.GetProductsRatingAndCommsCountFromStore(product.ID)
	if err != nil {
		return nil, err
	}

	product.Rating = math.Round(rating*100) / 100
	product.CommentsCount = &commsCount

	product.Properties = []*model.Property{}
	if userID != 0 {
		isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
		if err != nil {
			return nil, err
		}
		product.IsFavorite = isFav
	}

	return product, nil
}

func (api *ProductUsecase) GetProductsBySearch(search string, userID int) ([]*model.Product, error) {
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

		properties, err := api.store.GetProductPropertiesFromStore(product.ID, product.Category)
		if err != nil {
			return nil, err
		}
		product.Properties = properties
		if len(properties) > 4 {
			product.Properties = properties[:4]
		}

		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
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
	for _, product := range cart.Items {
		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.Item.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
	}
	return cart, nil
}

func (api *ProductUsecase) UpdateOrder(userID int, items *[]int) error {
	cart, err := api.store.GetCart(userID)
	if err != nil {
		return err
	}
	if cart == nil || cart.ID == 0 {
		err = api.store.CreateCart(userID)
		if err != nil {
			return err
		}
		cart, err = api.store.GetCart(userID)
		if err != nil {
			return err
		}
	}
	err = api.store.UpdateCart(userID, items)
	if err != nil {
		return err
	}
	if cart.Promocode != nil {
		return api.RecalculatePrices(userID, *cart.Promocode)
	}
	return nil
}

func (api *ProductUsecase) RecalculatePrices(userID int, promocode string) error {
	runeSlice := []rune(promocode)
	byteSlice := []byte(promocode)
	typeP := byteSlice[0]
	discount := int(runeSlice[1]-'0')*10 + int(runeSlice[2]-'0')

	err := api.store.UpdatePricesOrderItemsInStore(userID, "clear", 0)
	if err != nil {
		return err
	}
	switch string(typeP) {
	case "A":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "all", discount)
	case "C":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "computers", discount)
	case "M":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "monitors", discount)
	case "P":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "phones", discount)
	case "V":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "tvs", discount)
	case "W":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "watches", discount)
	case "T":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "tablets", discount)
	case "X":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "accessories", discount)
	default:
		err = baseErrors.ErrForbidden403
	}
	return err
}

func (api *ProductUsecase) SetPromocode(userID int, promocode string) error {
	err := api.store.CheckPromocodeUsage(userID, promocode)
	if err != nil {
		return err
	}
	for _, specialPromo := range conf.Promos {
		if specialPromo == promocode {
			return api.store.SetPromocodeDB(userID, promocode)
		}
	}
	if promocode == "" {
		err = api.store.UpdatePricesOrderItemsInStore(userID, "clear", 0)
		if err != nil {
			return err
		}
		return api.store.SetPromocodeDB(userID, promocode)
	}
	if len(promocode) < 8 {
		return baseErrors.ErrForbidden403
	}
	err = api.RecalculatePrices(userID, promocode)
	if err != nil {
		return err
	}
	h := hmac.New(sha256.New, []byte("Base2022"))
	data := fmt.Sprintf("%d", userID)
	h.Write([]byte(data))
	hashedStr := hex.EncodeToString(h.Sum(nil))
	if hashedStr[:5] != promocode[3:] {
		return baseErrors.ErrUnauthorized401
	}
	if err != nil {
		return err
	}
	return api.store.SetPromocodeDB(userID, promocode)
}

func (api *ProductUsecase) AddToOrder(userID int, itemID int) error {
	cart, err := api.store.GetCart(userID)
	if err != nil {
		return err
	}
	if cart == nil || cart.ID == 0 {
		err = api.store.CreateCart(userID)
		if err != nil {
			return err
		}
		cart, err = api.store.GetCart(userID)
		if err != nil {
			return err
		}
	}
	flag := true
	for _, prod := range cart.Items {
		if prod.Item.ID == itemID {
			err = api.store.InsertItemIntoCartById(userID, itemID, cart.ID, 1, true)
			if err != nil {
				return err
			}
			flag = false
			break
		}
	}
	// item wasn't in cart
	if flag {
		err = api.store.InsertItemIntoCartById(userID, itemID, cart.ID, 1, false)
		if err != nil {
			return err
		}
	}

	if cart.Promocode != nil {
		return api.RecalculatePrices(userID, *cart.Promocode)
	}
	return nil
}

func (api *ProductUsecase) DeleteFromOrder(userID int, itemID int) error {
	return api.store.DeleteItemFromCartById(userID, itemID)
}

func (api *ProductUsecase) MakeOrder(in *model.MakeOrder) (int, error) {
	cart, err := api.store.GetCart(in.UserID)
	if err != nil {
		return 0, err
	}
	if cart == nil || cart.ID == 0 {
		err = api.store.CreateCart(in.UserID)
		if err != nil {
			return 0, err
		}
		cart, err = api.store.GetCart(in.UserID)
		if err != nil {
			return 0, err
		}
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
		return 0, err
	}
	if cart.Promocode != nil {
		err = api.RecalculatePrices(in.UserID, *cart.Promocode)
		if err != nil {
			return 0, err
		}
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
		return 0, err
	}

	err = api.store.CreateCart(in.UserID)
	if err != nil {
		return 0, err
	}

	return cart.ID, api.store.UpdateCart(in.UserID, &remainedItemsIDs)
}

func (api *ProductUsecase) ChangeOrderStatus(userID int, in *model.ChangeOrderStatus) error {

	_, err := api.ordersManager.ChangeOrderStatus(
		context.Background(),
		&orders.ChangeOrderStatusType{
			UserID:      int32(userID),
			OrderID:     int32(in.OrderID),
			OrderStatus: in.OrderStatus,
		})
	if err != nil {
		return err
	}

	return nil
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

func (api *ProductUsecase) GetComments(productID int) ([]*model.CommentDB, error) {
	comments, err := api.store.GetCommentsFromStore(productID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (api *ProductUsecase) CreateComment(in *model.CreateComment) error {
	err := api.store.CreateCommentInStore(in)
	if err != nil {
		return err
	}
	return api.store.UpdateProductRatingInStore(in.ItemID)
}

func (api *ProductUsecase) GetRecommendationProducts(itemID int, userID int) ([]*model.Product, error) {
	products, err := api.store.GetRecommendationProductsFromStore(itemID)
	if err != nil {
		return nil, err
	}
	ansproducts := []*model.Product{}
	// shuffle
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(products), func(i, j int) { products[i], products[j] = products[j], products[i] })

	for _, product := range products {
		rating, commsCount, err := api.store.GetProductsRatingAndCommsCountFromStore(product.ID)
		if err != nil {
			return nil, err
		}
		product.Rating = math.Round(rating*100) / 100
		product.CommentsCount = &commsCount
		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
		if product.ID != itemID {
			ansproducts = append(ansproducts, product)
		}
	}
	return ansproducts, nil
}

func (api *ProductUsecase) GetFavorites(userID int, lastitemid int, count int, sort string) ([]*model.Product, error) {
	products, err := api.store.GetFavoritesDB(userID, lastitemid, count, sort)
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

		properties, err := api.store.GetProductPropertiesFromStore(product.ID, product.Category)
		if err != nil {
			return nil, err
		}
		product.Properties = properties
		if len(properties) > 4 {
			product.Properties = properties[:4]
		}
		if userID != 0 {
			isFav, err := api.store.CheckIsProductInFavoritesDB(userID, product.ID)
			if err != nil {
				return nil, err
			}
			product.IsFavorite = isFav
		}
	}
	return products, nil
}

func (api *ProductUsecase) InsertItemIntoFavorites(userID int, itemID int) error {
	return api.store.InsertItemIntoFavoritesDB(userID, itemID)
}

func (api *ProductUsecase) DeleteItemFromFavorites(userID int, itemID int) error {
	return api.store.DeleteItemFromFavoritesDB(userID, itemID)
}

func (api *ProductUsecase) RecalculateRatingsForInitscriptProducts(count int) error {
	for i := 1; i <= count; i++ {
		err := api.store.UpdateProductRatingInStore(i)
		if err != nil {
			return err
		}
	}
	return nil
}
