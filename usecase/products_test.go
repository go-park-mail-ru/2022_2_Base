package usecase

import (
	"context"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	"testing"

	baseErrors "serv/domain/errors"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "serv/mocks"
)

func TestGetProducts(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)
	mockLastItemID := 0
	mockCount := 1
	mockSort := ""
	testProducts := new([3]*model.Product)
	err := faker.FakeData(testProducts)
	testProductsSlice := testProducts[:]
	for _, testProd := range testProductsSlice {
		testProd.Category = "phones"
	}
	assert.NoError(t, err)
	prodStoreMock.EXPECT().GetProductsFromStore(mockLastItemID, mockCount, mockSort).Return(testProductsSlice, nil)
	for _, testProd := range testProductsSlice {
		prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProd.ID).Return(testProd.Rating, *testProd.CommentsCount, nil)
	}
	products, err := prodUsecase.GetProducts(mockLastItemID, mockCount, mockSort)
	assert.NoError(t, err)
	assert.Equal(t, testProductsSlice, products)
	// error
	prodStoreMock.EXPECT().GetProductsFromStore(mockLastItemID, mockCount, mockSort).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetProducts(mockLastItemID, mockCount, mockSort)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//GetProductsWithCategory
	err = faker.FakeData(testProducts)
	testProductsSlice = testProducts[:]
	for _, testProd := range testProductsSlice {
		testProd.Category = "phones"
	}
	assert.NoError(t, err)
	prodStoreMock.EXPECT().GetProductsWithCategoryFromStore("phones", mockLastItemID, mockCount, mockSort).Return(testProductsSlice, nil)
	for _, testProd := range testProductsSlice {
		prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProd.ID).Return(testProd.Rating, *testProd.CommentsCount, nil)
		prodStoreMock.EXPECT().GetProductPropertiesFromStore(testProd.ID, testProd.Category).Return(testProd.Properties, nil)
	}
	products, err = prodUsecase.GetProductsWithCategory("phones", mockLastItemID, mockCount, mockSort)
	assert.NoError(t, err)
	assert.Equal(t, testProductsSlice, products)
	// error
	prodStoreMock.EXPECT().GetProductsWithCategoryFromStore("phones", mockLastItemID, mockCount, mockSort).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetProductsWithCategory("phones", mockLastItemID, mockCount, mockSort)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetProductsByIDAndBySearch(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)
	testProducts := new([3]*model.Product)
	err := faker.FakeData(testProducts)
	testProductsSlice := testProducts[:]
	err = faker.FakeData(testProducts)
	testProductsSlice = testProducts[:]
	search := testProductsSlice[0].Name
	assert.NoError(t, err)

	//by id
	prodStoreMock.EXPECT().GetProductFromStoreByID(testProductsSlice[0].ID).Return(testProductsSlice[0], nil)
	prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProductsSlice[0].ID).Return(testProductsSlice[0].Rating, *testProductsSlice[0].CommentsCount, nil)
	prodStoreMock.EXPECT().GetProductPropertiesFromStore(testProductsSlice[0].ID, testProductsSlice[0].Category).Return(testProductsSlice[0].Properties, nil)

	product, err := prodUsecase.GetProductByID(testProductsSlice[0].ID)
	assert.NoError(t, err)
	assert.Equal(t, testProductsSlice[0], product)

	// error
	prodStoreMock.EXPECT().GetProductFromStoreByID(testProductsSlice[0].ID).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetProductByID(testProductsSlice[0].ID)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//by searh
	prodStoreMock.EXPECT().GetProductsBySearchFromStore(search).Return([]*model.Product{testProductsSlice[0]}, nil)
	prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProductsSlice[0].ID).Return(testProductsSlice[0].Rating, *testProductsSlice[0].CommentsCount, nil)
	prodStoreMock.EXPECT().GetProductPropertiesFromStore(testProductsSlice[0].ID, testProductsSlice[0].Category).Return(testProductsSlice[0].Properties, nil)

	products, err := prodUsecase.GetProductsBySearch(search)
	assert.NoError(t, err)
	assert.Equal(t, []*model.Product{testProductsSlice[0]}, products)

	// error
	prodStoreMock.EXPECT().GetProductsBySearchFromStore(search).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetProductsBySearch(search)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetCart(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)

	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)
	testCart := new(model.Order)
	err := faker.FakeData(testCart)
	assert.NoError(t, err)

	//exist cart
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	cart, err := prodUsecase.GetCart(testCart.UserID)
	assert.NoError(t, err)
	assert.Equal(t, testCart, cart)

	//new cart
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(nil, nil)
	prodStoreMock.EXPECT().CreateCart(testCart.UserID).Return(nil)
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	cart, err = prodUsecase.GetCart(testCart.UserID)
	assert.NoError(t, err)
	assert.Equal(t, testCart, cart)

	//error
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetCart(testCart.UserID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestUpdateOrder(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)

	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)
	testCart := new(model.Order)
	err := faker.FakeData(testCart)
	assert.NoError(t, err)
	testItemsIDs := new([5]int)
	err = faker.FakeData(testItemsIDs)
	assert.NoError(t, err)
	testItemsIDsSlice := testItemsIDs[:]
	testCart.Promocode = nil

	//UpdateOrder
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	prodStoreMock.EXPECT().UpdateCart(testCart.UserID, &testItemsIDsSlice).Return(nil)
	//prodStoreMock.EXPECT().UpdatePricesOrderItemsInStore(testCart.UserID, "clear", 0).Return(nil)
	//prodStoreMock.EXPECT().UpdatePricesOrderItemsInStore(testCart.UserID, "phones", 20).Return(nil)
	err = prodUsecase.UpdateOrder(testCart.UserID, &testItemsIDsSlice)
	assert.NoError(t, err)

	//error 500
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(nil, baseErrors.ErrServerError500)
	err = prodUsecase.UpdateOrder(testCart.UserID, &testItemsIDsSlice)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//error 500
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	prodStoreMock.EXPECT().UpdateCart(testCart.UserID, &testItemsIDsSlice).Return(baseErrors.ErrServerError500)
	err = prodUsecase.UpdateOrder(testCart.UserID, &testItemsIDsSlice)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//AddToOrder
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	prodStoreMock.EXPECT().InsertItemIntoCartById(testCart.UserID, testItemsIDsSlice[0]).Return(nil)
	err = prodUsecase.AddToOrder(testCart.UserID, testItemsIDsSlice[0])
	assert.NoError(t, err)

	//error 500
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	prodStoreMock.EXPECT().InsertItemIntoCartById(testCart.UserID, testItemsIDsSlice[0]).Return(baseErrors.ErrServerError500)
	err = prodUsecase.AddToOrder(testCart.UserID, testItemsIDsSlice[0])
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//error 500
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(nil, baseErrors.ErrServerError500)
	err = prodUsecase.AddToOrder(testCart.UserID, testItemsIDsSlice[0])
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//DeleteFromOrder
	prodStoreMock.EXPECT().DeleteItemFromCartById(testCart.UserID, testItemsIDsSlice[0]).Return(nil)
	err = prodUsecase.DeleteFromOrder(testCart.UserID, testItemsIDsSlice[0])
	assert.NoError(t, err)

	//error
	prodStoreMock.EXPECT().DeleteItemFromCartById(testCart.UserID, testItemsIDsSlice[0]).Return(baseErrors.ErrServerError500)
	err = prodUsecase.DeleteFromOrder(testCart.UserID, testItemsIDsSlice[0])
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestMakeOrder(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)

	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)
	testCart := new(model.Order)
	err := faker.FakeData(testCart)
	assert.NoError(t, err)
	testOrder := new(model.MakeOrder)
	err = faker.FakeData(testOrder)
	assert.NoError(t, err)
	testCart.Promocode = nil

	testCart.UserID = testOrder.UserID

	//ok
	prodStoreMock.EXPECT().GetCart(testOrder.UserID).Return(testCart, nil)

	remainedItemsIDs := []int{}
	boughtItemsIDs := []int{}
	boughtItemsIDsINT32 := []int32{}
	for _, orderItem := range testCart.Items {
		flag := true
		flag2 := false
		for _, id := range testOrder.Items {
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

	prodStoreMock.EXPECT().UpdateCart(testOrder.UserID, &boughtItemsIDs).Return(nil)
	ordersManager.EXPECT().MakeOrder(context.Background(),
		&orders.MakeOrderType{
			UserID:        int32(testOrder.UserID),
			Items:         boughtItemsIDsINT32,
			AddressID:     int32(testOrder.AddressID),
			PaymentcardID: int32(testOrder.PaymentcardID),
			DeliveryDate:  testOrder.DeliveryDate.Unix(),
		}).Return(nil, nil)
	prodStoreMock.EXPECT().CreateCart(testOrder.UserID).Return(nil)
	prodStoreMock.EXPECT().UpdateCart(testOrder.UserID, &remainedItemsIDs).Return(nil)

	_, err = prodUsecase.MakeOrder(testOrder)
	assert.NoError(t, err)

	//error
	prodStoreMock.EXPECT().GetCart(testOrder.UserID).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.MakeOrder(testOrder)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetOrders(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)

	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)
	testOrders := new(orders.OrdersResponse)
	err := faker.FakeData(testOrders)
	assert.NoError(t, err)

	var userID int = int(testOrders.Orders[0].UserID)
	//ok
	ordersManager.EXPECT().GetOrders(
		context.Background(),
		&orders.UserID{
			UserID: int32(userID),
		}).Return(testOrders, nil)
	orders3, err := prodUsecase.GetOrders(userID)
	assert.NoError(t, err)
	assert.Equal(t, testOrders, orders3)

	//error
	ordersManager.EXPECT().GetOrders(
		context.Background(),
		&orders.UserID{
			UserID: int32(userID),
		}).Return(nil, baseErrors.ErrServerError500)

	_, err = prodUsecase.GetOrders(userID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetComments(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)

	testComms := new([5]*model.CommentDB)
	err := faker.FakeData(testComms)
	assert.NoError(t, err)
	testCommsSlice := testComms[:]

	var itemID int = testComms[0].ItemID
	//ok
	prodStoreMock.EXPECT().GetCommentsFromStore(itemID).Return(testCommsSlice, nil)
	comms, err := prodUsecase.GetComments(itemID)
	assert.NoError(t, err)
	assert.Equal(t, testCommsSlice, comms)

	//error
	prodStoreMock.EXPECT().GetCommentsFromStore(itemID).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetComments(itemID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestCreateComment(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	prodStoreMock := mocks.NewMockProductStoreInterface(ctrl)
	ordersManager := mocks.NewMockOrdersWorkerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager, mailManager)

	testComm := new(model.CreateComment)
	err := faker.FakeData(testComm)
	assert.NoError(t, err)

	//ok
	prodStoreMock.EXPECT().CreateCommentInStore(testComm).Return(nil)
	prodStoreMock.EXPECT().UpdateProductRatingInStore(testComm.ItemID).Return(nil)
	err = prodUsecase.CreateComment(testComm)
	assert.NoError(t, err)

	//error
	prodStoreMock.EXPECT().CreateCommentInStore(testComm).Return(baseErrors.ErrServerError500)
	err = prodUsecase.CreateComment(testComm)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}
