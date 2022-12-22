package delivery

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	conf "serv/config"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	auth "serv/microservices/auth/gen_files"
	mocks "serv/mocks"
)

func TestGetCart(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	userHandler := NewUserHandler(userUsecaseMock)
	productHandler := NewProductHandler(productUsecaseMock)
	orderHandler := NewOrderHandler(userHandler, productHandler)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)
	testCart := new(model.Order)
	err = faker.FakeData(testCart)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().GetCart(testUserProfile.ID).Return(testCart, nil)
	url := "/api/v1/cart"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.GetCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 500 no user
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.GetCart(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().GetCart(testUserProfile.ID).Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.GetCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}

// func TestUpdateCart(t *testing.T) {
// 	t.Parallel()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
// 	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
// 	userHandler := NewUserHandler(userUsecaseMock)
// 	productHandler := NewProductHandler(productUsecaseMock)
// 	orderHandler := NewOrderHandler(userHandler, productHandler)

// 	testUserProfile := new(model.UserProfile)
// 	err := faker.FakeData(testUserProfile)
// 	assert.NoError(t, err)
// 	testProductCart := new(model.ProductCart)
// 	err = faker.FakeData(testProductCart)
// 	assert.NoError(t, err)
// 	testCart := new(model.Order)
// 	err = faker.FakeData(testCart)
// 	assert.NoError(t, err)
// 	testCart.Promocode = nil

// 	//ok
// 	productUsecaseMock.EXPECT().GetCart(testUserProfile.ID).Return(testCart, nil)
// 	productUsecaseMock.EXPECT().UpdateOrder(testUserProfile.ID, &testProductCart.Items).Return(nil)

// 	url := "/api/v1/cart"
// 	data, _ := json.Marshal(testProductCart)
// 	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	orderHandler.UpdateCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// //err 400 query err
// 	// data, _ = json.Marshal("sfdsd")
// 	// req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// rr = httptest.NewRecorder()
// 	// orderHandler.UpdateCart(rr, req)
// 	// assert.Equal(t, 400, rr.Code)

// 	// //err 500 no user
// 	// data, _ = json.Marshal(testProductCart)
// 	// req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// rr = httptest.NewRecorder()
// 	// orderHandler.UpdateCart(rr, req)
// 	// assert.Equal(t, 500, rr.Code)

// 	// //err 500 db
// 	// productUsecaseMock.EXPECT().GetCart(testUserProfile.ID).Return(nil, baseErrors.ErrServerError500)
// 	// data, _ = json.Marshal(testProductCart)
// 	// req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// rr = httptest.NewRecorder()
// 	// orderHandler.UpdateCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
// 	// assert.Equal(t, 500, rr.Code)

// 	// //err 500 db
// 	// productUsecaseMock.EXPECT().GetCart(testUserProfile.ID).Return(testCart, nil)
// 	// productUsecaseMock.EXPECT().UpdateOrder(testUserProfile.ID, &testProductCart.Items).Return(baseErrors.ErrServerError500)
// 	// data, _ = json.Marshal(testProductCart)
// 	// req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// rr = httptest.NewRecorder()
// 	// orderHandler.UpdateCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
// 	// assert.Equal(t, 500, rr.Code)
// }


func TestAddItemToCart(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	userHandler := NewUserHandler(userUsecaseMock)
	productHandler := NewProductHandler(productUsecaseMock)
	orderHandler := NewOrderHandler(userHandler, productHandler)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	testItem := new(model.ProductCartItem)
	err = faker.FakeData(testItem)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().AddToOrder(testUserProfile.ID, testItem.ItemID).Return(nil)
	url := "/api/v1/cart" + conf.PathAddItemToCart
	data, _ := json.Marshal(testItem)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.AddItemToCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.AddItemToCart(rr, req)
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.AddItemToCart(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().AddToOrder(testUserProfile.ID, testItem.ItemID).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.AddItemToCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}

func TestDeleteItemFromCart(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	userHandler := NewUserHandler(userUsecaseMock)
	productHandler := NewProductHandler(productUsecaseMock)
	orderHandler := NewOrderHandler(userHandler, productHandler)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	testItem := new(model.ProductCartItem)
	err = faker.FakeData(testItem)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().DeleteFromOrder(testUserProfile.ID, testItem.ItemID).Return(nil)
	url := "/api/v1/cart" + conf.PathDeleteItemFromCart
	data, _ := json.Marshal(testItem)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.DeleteItemFromCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.DeleteItemFromCart(rr, req)
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.DeleteItemFromCart(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().DeleteFromOrder(testUserProfile.ID, testItem.ItemID).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.DeleteItemFromCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 404 db
	productUsecaseMock.EXPECT().DeleteFromOrder(testUserProfile.ID, testItem.ItemID).Return(baseErrors.ErrNotFound404)
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.DeleteItemFromCart(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 404, rr.Code)
}
