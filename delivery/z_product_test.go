package delivery

import (
	"encoding/json"
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
	//log.Println(rr.Body.String())
	expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	//log.Println(string(expectedstr))
	//log.Println(easyjson.Marshal(&model.Response{Body: testProducts}))
	//assert.Equal(t, rr.Body.String(), expected+"\n")
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
