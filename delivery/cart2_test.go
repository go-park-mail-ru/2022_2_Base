package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"time"

	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	baseErrors "serv/domain/errors"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	mocks "serv/mocks"
)

func TestGetOrders(t *testing.T) {
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
	testOrders := new(orders.OrdersResponse)
	err = faker.FakeData(testOrders)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().GetOrders(testUserProfile.ID).Return(testOrders, nil)
	url := "/api/v1/" + "cart/orders"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.GetOrders(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 500 microservice
	productUsecaseMock.EXPECT().GetOrders(testUserProfile.ID).Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.GetOrders(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}

func TestMakeOrder(t *testing.T) {
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
	testOrder := new(model.MakeOrder)
	err = faker.FakeData(testOrder)
	assert.NoError(t, err)
	testOrder.UserID = testUserProfile.ID
	testOrder.DeliveryDate = time.Unix(1, 10)
	testOrderID := 1
	testMail := model.Mail{Type: "orderstatus", Username: testUserProfile.Username, Useremail: testUserProfile.Email, OrderID: testOrderID, OrderStatus: "created"}

	//ok
	productUsecaseMock.EXPECT().MakeOrder(testOrder).Return(testOrderID, nil)
	wg := sync.WaitGroup{}
	wg.Add(1)
	userUsecaseMock.EXPECT().SendMail(testMail).Do(func(arg1 interface{}) {
		defer wg.Done()
	})
	url := "/api/v1/" + "cart/makeorder"
	data, _ := json.Marshal(testOrder)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.MakeOrder(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	wg.Wait()
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.MakeOrder(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testOrder)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.MakeOrder(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().MakeOrder(testOrder).Return(0, baseErrors.ErrServerError500)
	data, _ = json.Marshal(testOrder)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.MakeOrder(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 401 db
	testOrder.UserID = testUserProfile.ID + 1
	data, _ = json.Marshal(testOrder)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.CreateComment(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 401, rr.Code)

	// //err 500 microservice
	// productUsecaseMock.EXPECT().MakeOrder(testOrder).Return(testOrderID, nil)
	// userUsecaseMock.EXPECT().SendMail(testMail).Return(baseErrors.ErrServerError500)
	// data, _ = json.Marshal(testOrder)
	// req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// rr = httptest.NewRecorder()
	// orderHandler.MakeOrder(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	// assert.Equal(t, 500, rr.Code)
}

func TestGetComments(t *testing.T) {
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
	testCommentsDB := new([5]*model.CommentDB)
	err = faker.FakeData(testCommentsDB)
	assert.NoError(t, err)
	testCommentsDBSlice := testCommentsDB[:]
	testComments := new([5]*model.Comment)
	err = faker.FakeData(testComments)
	assert.NoError(t, err)
	testCommentsSlice := testComments[:]
	mockItemID := 1

	//ok
	productUsecaseMock.EXPECT().GetComments(mockItemID).Return(testCommentsDBSlice, nil)
	userUsecaseMock.EXPECT().SetUsernamesForComments(testCommentsDBSlice).Return(testCommentsSlice, nil)
	url := "/api/v1/" + "products/comments/" + fmt.Sprint(mockItemID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.GetComments(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().GetComments(mockItemID).Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.GetComments(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().GetComments(mockItemID).Return(testCommentsDBSlice, nil)
	userUsecaseMock.EXPECT().SetUsernamesForComments(testCommentsDBSlice).Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.GetComments(rr, req)
	assert.Equal(t, 500, rr.Code)
}

func TestCreateComment(t *testing.T) {
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
	testComment := new(model.CreateComment)
	err = faker.FakeData(testComment)
	assert.NoError(t, err)
	testComment.UserID = testUserProfile.ID

	//ok
	productUsecaseMock.EXPECT().CreateComment(testComment).Return(nil)
	url := "/api/v1/user" + "/makecomment"
	data, _ := json.Marshal(testComment)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.CreateComment(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.CreateComment(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testComment)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.CreateComment(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().CreateComment(testComment).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testComment)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.CreateComment(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 401 db
	testComment.UserID = testUserProfile.ID + 1
	data, _ = json.Marshal(testComment)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.CreateComment(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 401, rr.Code)
}

func TestSetPromocode(t *testing.T) {
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
	testPromocode := new(model.Promocode)
	err = faker.FakeData(testPromocode)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().SetPromocode(testUserProfile.ID, testPromocode.Promocode).Return(nil)
	url := "/api/v1" + "/cart/setpromocode"
	data, _ := json.Marshal(testPromocode)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testPromocode)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().SetPromocode(testUserProfile.ID, testPromocode.Promocode).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testPromocode)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 409 db
	productUsecaseMock.EXPECT().SetPromocode(testUserProfile.ID, testPromocode.Promocode).Return(baseErrors.ErrConflict409)
	data, _ = json.Marshal(testPromocode)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 409, rr.Code)

	//err 401 db
	productUsecaseMock.EXPECT().SetPromocode(testUserProfile.ID, testPromocode.Promocode).Return(baseErrors.ErrUnauthorized401)
	data, _ = json.Marshal(testPromocode)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 403, rr.Code)

	//err 403 db
	productUsecaseMock.EXPECT().SetPromocode(testUserProfile.ID, testPromocode.Promocode).Return(baseErrors.ErrForbidden403)
	data, _ = json.Marshal(testPromocode)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	orderHandler.SetPromocode(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 403, rr.Code)
}
