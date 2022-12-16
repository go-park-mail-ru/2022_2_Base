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
	//"github.com/mailru/easyjson/jwriter"
)

func TestGetHomePage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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
	expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")

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
	url = conf.PathMain + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=z" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetHomePage(rr, req)
	assert.Equal(t, 400, rr.Code)
}

func TestGetProductsByCategory(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	testProducts := new([10]*model.Product)
	err := faker.FakeData(testProducts)
	assert.NoError(t, err)
	testProductsSlice := testProducts[:]
	mockLastItemID := 0
	mockCount := 10
	mockSort := ""
	mockCategory := "phones"
	for _, testProductsSliceIt := range testProductsSlice {
		testProductsSliceIt.Category = "phones"
	}

	//default sort
	productUsecaseMock.EXPECT().GetProductsWithCategory(mockCategory, mockLastItemID, mockCount, mockSort).Return(testProductsSlice, nil)
	url := conf.BasePath + "/category/" + mockCategory + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler := NewProductHandler(productUsecaseMock)
	productHandler.GetProductsByCategory(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")

	//err 500
	productUsecaseMock.EXPECT().GetProductsWithCategory(mockCategory, mockLastItemID, mockCount, mockSort).Return(nil, baseErrors.ErrServerError500)
	url = conf.BasePath + "/category/" + mockCategory + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetProductsByCategory(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 400
	url = conf.BasePath + "/category/" + mockCategory + "?lastitemid=" + fmt.Sprint(mockLastItemID) + "&count=z" + fmt.Sprint(mockCount) + "&sort=" + mockSort
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetProductsByCategory(rr, req)
	assert.Equal(t, 400, rr.Code)
}

func TestGetProductByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	testProduct := new(model.Product)
	err := faker.FakeData(testProduct)
	assert.NoError(t, err)
	mockItemID := 1

	//ok
	productUsecaseMock.EXPECT().GetProductByID(mockItemID).Return(testProduct, nil)
	url := conf.PathMain + "/" + fmt.Sprint(mockItemID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler := NewProductHandler(productUsecaseMock)
	productHandler.GetProductByID(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	expectedstr, err := json.Marshal(&model.Response{Body: testProduct})
	assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")

	//err 500
	productUsecaseMock.EXPECT().GetProductByID(mockItemID).Return(nil, baseErrors.ErrServerError500)
	url = conf.PathMain + "/" + fmt.Sprint(mockItemID)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetProductByID(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 400
	url = conf.PathMain + "/z" + fmt.Sprint(mockItemID)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetProductByID(rr, req)
	assert.Equal(t, 400, rr.Code)
}

func TestGetProductsBySearch(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	testProducts := new([10]*model.Product)
	err := faker.FakeData(testProducts)
	assert.NoError(t, err)
	testProductsSlice := testProducts[:]
	for _, testProductsSliceIt := range testProductsSlice {
		testProductsSliceIt.Name = "item" + fmt.Sprint(1)
	}

	//ok
	productUsecaseMock.EXPECT().GetProductsBySearch("item").Return(testProductsSlice, nil)
	url := conf.PathSeacrh
	data, _ := json.Marshal(&model.Search{Search: "item"})
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler := NewProductHandler(productUsecaseMock)
	productHandler.GetProductsBySearch(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")

	//err 500
	productUsecaseMock.EXPECT().GetProductsBySearch("item").Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetProductsBySearch(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 400
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetProductsBySearch(rr, req)
	assert.Equal(t, 400, rr.Code)
}

func TestGetSuggestions(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productUsecaseMock := mocks.NewMockProductUsecaseInterface(ctrl)
	testProducts := new([3]*model.Product)
	err := faker.FakeData(testProducts)
	assert.NoError(t, err)
	testProductsSlice := testProducts[:]
	var expstrings []string
	for i, testProductsSliceIt := range testProductsSlice {
		testProductsSliceIt.Name = "item" + fmt.Sprint(i)
		expstrings = append(expstrings, testProductsSliceIt.Name)
	}

	//ok
	productUsecaseMock.EXPECT().GetSuggestions("item").Return(expstrings, nil)
	url := conf.PathSuggestions
	data, _ := json.Marshal(&model.Search{Search: "item"})
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	productHandler := NewProductHandler(productUsecaseMock)
	productHandler.GetSuggestions(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	expected, err := json.Marshal(&model.Response{Body: expstrings})
	assert.Equal(t, rr.Body.String(), string(expected)+"\n")

	//err 500
	productUsecaseMock.EXPECT().GetSuggestions("item").Return(nil, baseErrors.ErrServerError500)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetSuggestions(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 400
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	productHandler.GetSuggestions(rr, req)
	assert.Equal(t, 400, rr.Code)
}
