package delivery

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

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
