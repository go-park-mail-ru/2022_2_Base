package delivery

import (
	"encoding/json"
	"fmt"
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
	mocks "serv/mocks"
)

func TestAddItemToFavorites(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	productHandler := NewProductHandler(productUsecaseMock)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	testItem := new(model.ProductCartItem)
	err = faker.FakeData(testItem)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().InsertItemIntoFavorites(testUserProfile.ID, testItem.ItemID).Return(nil)
	url := "/api/v1/user" + conf.PathInsertIntoFavorites
	data, _ := json.Marshal(testItem)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler.InsertItemIntoFavorites(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.InsertItemIntoFavorites(rr, req)
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.InsertItemIntoFavorites(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().InsertItemIntoFavorites(testUserProfile.ID, testItem.ItemID).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.InsertItemIntoFavorites(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}

func TestDeleteItemFromFavorites(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	productHandler := NewProductHandler(productUsecaseMock)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	testItem := new(model.ProductCartItem)
	err = faker.FakeData(testItem)
	assert.NoError(t, err)

	//ok
	productUsecaseMock.EXPECT().DeleteItemFromFavorites(testUserProfile.ID, testItem.ItemID).Return(nil)
	url := "/api/v1/user" + conf.PathDeleteFromFavorites
	data, _ := json.Marshal(testItem)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler.DeleteItemFromFavorites(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 400 query err
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.DeleteItemFromFavorites(rr, req)
	assert.Equal(t, 400, rr.Code)

	//err 500 no user
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.DeleteItemFromFavorites(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 db
	productUsecaseMock.EXPECT().DeleteItemFromFavorites(testUserProfile.ID, testItem.ItemID).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testItem)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.DeleteItemFromFavorites(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}

func TestGetFavorites(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)

	testProducts := new([2]*model.Product)
	err := faker.FakeData(testProducts)
	assert.NoError(t, err)
	testProductsSlice := testProducts[:]
	testUserProfile := new(model.UserProfile)
	err = faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	mockLastItemID := 0
	mockCount := 10
	mockSort := ""

	//ok
	productUsecaseMock.EXPECT().GetFavorites(testUserProfile.ID, mockLastItemID, mockCount, mockSort).Return(testProductsSlice, nil)
	url := "/api/v1/user" + conf.PathFavorites + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler := NewProductHandler(productUsecaseMock)
	productHandler.GetFavorites(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)
	//expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	//assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")

	//err 500 db
	productUsecaseMock.EXPECT().GetFavorites(testUserProfile.ID, mockLastItemID, mockCount, mockSort).Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetFavorites(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 500 no user
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetFavorites(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 400
	url = conf.PathMain + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=z" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetFavorites(rr, req)
	assert.Equal(t, 400, rr.Code)

}
