package usecase

import (
	"serv/domain/model"
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
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager)
	mockLastItemID := 0
	mockCount := 1
	mockSort := ""
	testProducts := new([3]*model.Product)
	err := faker.FakeData(testProducts)
	testProductsSlice := testProducts[:]
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
	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager)
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

	prodUsecase := NewProductUsecase(prodStoreMock, ordersManager)
	testCart := new(model.Order)
	err := faker.FakeData(testCart)
	//testProductsSlice := testProducts[:]
	//err = faker.FakeData(testProducts)
	//testProductsSlice = testProducts[:]
	//search := testProductsSlice[0].Name
	assert.NoError(t, err)

	//var userID int = 1

	//exist cart
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	//prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProductsSlice[0].ID).Return(testProductsSlice[0].Rating, *testProductsSlice[0].CommentsCount, nil)
	cart, err := prodUsecase.GetCart(testCart.UserID)
	assert.NoError(t, err)
	assert.Equal(t, testCart, cart)

	//new cart
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(nil, nil)
	prodStoreMock.EXPECT().CreateCart(testCart.UserID).Return(nil)
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(testCart, nil)
	//prodStoreMock.EXPECT().GetProductsRatingAndCommsCountFromStore(testProductsSlice[0].ID).Return(testProductsSlice[0].Rating, *testProductsSlice[0].CommentsCount, nil)
	cart, err = prodUsecase.GetCart(testCart.UserID)
	assert.NoError(t, err)
	assert.Equal(t, testCart, cart)

	//error
	prodStoreMock.EXPECT().GetCart(testCart.UserID).Return(nil, baseErrors.ErrServerError500)
	_, err = prodUsecase.GetCart(testCart.UserID)
	assert.Equal(t, baseErrors.ErrServerError500, err)

}
