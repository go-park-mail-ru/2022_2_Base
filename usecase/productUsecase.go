package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/rand"
	conf "serv/config"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	rep "serv/repository"
	"time"
)

type ProductUsecaseInterface interface {
	GetProducts(lastitemid int, count int, sort string) ([]*model.Product, error)
	GetProductsWithCategory(cat string, lastitemid int, count int, sort string) ([]*model.Product, error)
	GetProductByID(id int) (*model.Product, error)
	GetProductsBySearch(search string) ([]*model.Product, error)
	GetSuggestions(search string) ([]string, error)
	GetCart(userID int) (*model.Order, error)
	UpdateOrder(userID int, items *[]int) error
	AddToOrder(userID int, itemID int) error
	DeleteFromOrder(userID int, itemID int) error
	MakeOrder(in *model.MakeOrder) error
	GetOrders(userID int) (*orders.OrdersResponse, error)
	GetComments(productID int) ([]*model.CommentDB, error)
	CreateComment(in *model.CreateComment) error
	GetRecommendationProducts(itemID int) ([]*model.Product, error)
	SetPromocode(userID int, promocode string) error
	RecalculatePrices(userID int, promocode string) error
}

type ProductUsecase struct {
	ordersManager orders.OrdersWorkerClient
	store         rep.ProductStoreInterface
}

func NewProductUsecase(ps rep.ProductStoreInterface, ordersManager orders.OrdersWorkerClient) ProductUsecaseInterface {
	return &ProductUsecase{
		ordersManager: ordersManager,
		store:         ps,
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

// func ParsePromocode(promocode string) error {
// 	runeSlice := []rune(promocode)
// }
func (api *ProductUsecase) RecalculatePrices(userID int, promocode string) error {
	runeSlice := []rune(promocode)
	byteSlice := []byte(promocode)
	typeP := byteSlice[0]
	//var err error
	discount := int(runeSlice[1]-'0')*10 + int(runeSlice[2]-'0')
	log.Println(string(typeP), discount)

	err := api.store.UpdatePricesOrderItemsInStore(userID, "clear", 0)
	if err != nil {
		return err
	}
	log.Println("]]]]]]]]]]]")
	//strconv.QuoteRune(typeP)
	switch string(typeP) {
	case "A":
		err = api.store.UpdatePricesOrderItemsInStore(userID, "all", discount)
		log.Println("123535")
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
		err = nil

	}
	log.Println("sdsddsd")
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
	}
	if len(promocode) < 8 {
		return baseErrors.ErrForbidden403
	}
	err = api.RecalculatePrices(userID, promocode)
	if err != nil {
		return err
	}
	// salt := []byte("Base2022")
	// hashedPass := hashPass(salt, req.NewPassword)
	// b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	h := hmac.New(sha256.New, []byte("Base2022"))
	data := fmt.Sprintf("%d", userID)
	h.Write([]byte(data))
	hashedStr := hex.EncodeToString(h.Sum(nil))
	log.Println(hashedStr[:5])
	if hashedStr[:5] != promocode[3:] {
		return baseErrors.ErrUnauthorized401
	}
	//err = UpdatePricesOrderItemsInStore(userID, promocode)
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
	err = api.store.InsertItemIntoCartById(userID, itemID)
	if err != nil {
		return err
	}
	//log.Println("yyy")
	if cart.Promocode != nil {
		//log.Println("pr", *cart.Promocode)
		return api.RecalculatePrices(userID, *cart.Promocode)
		// if err != nil {
		// 	return err
		// }
	}
	//log.Println("www")
	return nil
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
	if cart.Promocode != nil {
		err = api.RecalculatePrices(in.UserID, *cart.Promocode)
		if err != nil {
			return err
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

func (api *ProductUsecase) GetRecommendationProducts(itemID int) ([]*model.Product, error) {
	products, err := api.store.GetRecommendationProductsFromStore(itemID)
	if err != nil {
		return nil, err
	}
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
	}
	return products, nil
}
