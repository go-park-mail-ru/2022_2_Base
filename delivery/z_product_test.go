package delivery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	conf "serv/config"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	mocks "serv/mocks"
)

func TestGetHomePage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	//sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	//userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	testProducts := new([10]*model.Product)
	err := faker.FakeData(testProducts)
	assert.NoError(t, err)
	testProductsSlice := testProducts[:]
	mockLastItemID := 0
	mockCount := 10
	mockSort := ""

	//default sort
	productUsecaseMock.EXPECT().GetProducts(mockLastItemID, mockCount, mockSort).Return(testProductsSlice, nil)
	url := conf.PathMain + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	productHandler := NewProductHandler(productUsecaseMock)
	productHandler.GetHomePage(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 500
	productUsecaseMock.EXPECT().GetProducts(mockLastItemID, mockCount, mockSort).Return(nil, baseErrors.ErrServerError500)
	url = conf.PathMain + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	productHandler.GetHomePage(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 400
	//productUsecaseMock.EXPECT().GetProducts(mockLastItemID, mockCount, mockSort).Return(nil, baseErrors.ErrServerError500)
	url = conf.PathMain + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=z" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	productHandler.GetHomePage(rr, req)
	assert.Equal(t, 400, rr.Code)

	//var expected = `{"body":[{"id":0,"name":"Монитор Xiaomi Mi 27","description":"good","price":14999,"lowprice":13999,"rating":4,"imgsrc":"https://img.mvideo.ru/Big/30058309bb.jpg"},{"id":1,"name":"Телевизор Haier 55","description":"goood","price":59999,"lowprice":41999,"rating":4.3,"imgsrc":"https://img.mvideo.ru/Big/10030234bb.jpg"},{"id":2,"name":"Apple iPad 10.2","description":"old","price":49999,"lowprice":49999,"rating":3.7,"imgsrc":"https://img.mvideo.ru/Pdb/30064043b.jpg"},{"id":3,"name":"Tecno Spark 8с","description":"good phone","price":12999,"lowprice":8999,"rating":4.5,"imgsrc":"https://img.mvideo.ru/Big/30062036bb.jpg"},{"id":4,"name":"realme GT Master","description":"goood","price":29999,"lowprice":21999,"rating":4.3,"imgsrc":"https://img.mvideo.ru/Big/30058843bb.jpg"},{"id":5,"name":"Apple iPhone 11","description":"old","price":62999,"lowprice":54999,"rating":5,"imgsrc":"https://img.mvideo.ru/Big/30063237bb.jpg"}]}`
	//assert.Equal(t, rr.Body, testProductsSlice)

}
