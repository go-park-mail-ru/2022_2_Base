package delivery

import (
	"testing"

	baseErrors "serv/errors"

	"github.com/stretchr/testify/assert"

	"serv/domain/model"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "serv/mocks"
)

func TestAdresses(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	//sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)

	testUser := new(model.UserProfile)
	err := faker.FakeData(testUser)
	assert.NoError(t, err)

	//ok get
	// userStoreMock.EXPECT().GetAddressesByUserIDFromDB(testUser.ID).Return(testUser.Address, nil)
	// adresses, err := userUsecase.GetAddressesByUserID(testUser.ID)
	// assert.NoError(t, err)
	// assert.Equal(t, adresses, testUser.Address)

	// //err get
	// userStoreMock.EXPECT().GetAddressesByUserIDFromDB(testUser.ID).Return(nil, baseErrors.ErrServerError500)
	// _, err = userUsecase.GetAddressesByUserID(testUser.ID)
	// assert.Equal(t, baseErrors.ErrServerError500, err)

	// testAddr := new([4]*model.Address)
	// err = faker.FakeData(testAddr)
	// assert.NoError(t, err)
	// testAddrSlice := testAddr[:]

	// //ok change
	// for _, addr := range testAddrSlice {
	// 	if addr == nil {
	// 		continue
	// 	}
	// 	flag := true
	// 	for _, addrFromDB := range testUser.Address {
	// 		if addr.ID == addrFromDB.ID {
	// 			userStoreMock.EXPECT().UpdateUsersAddress(addr.ID, addr).Return(nil)
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		userStoreMock.EXPECT().AddUsersAddress(testUser.ID, addr).Return(nil)
	// 	}
	// }
	// for _, addrFromDB := range testUser.Address {
	// 	flag := true
	// 	for _, addr := range testAddrSlice {
	// 		if addr == nil {
	// 			continue
	// 		}
	// 		if addr.ID == addrFromDB.ID {
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		userStoreMock.EXPECT().DeleteUsersAddress(addrFromDB.ID).Return(nil)
	// 	}
	// }

	// err = userUsecase.ChangeUserAddresses(testUser.ID, testUser.Address, testAddrSlice)
	// assert.NoError(t, err)
	// //assert.Equal(t, adresses, testUser.Address)
}

// var casesLogin = []struct {
// 	data     map[string]string
// 	wantCode int
// 	err      error
// }{
// 	{map[string]string{"email": "s", "username": "art", "password": "string"}, 401, baseErrors.ErrUnauthorized401},
// 	{map[string]string{"email": "string", "username": "art", "password": "s"}, 401, baseErrors.ErrUnauthorized401},
// }

// func TestLogin(t *testing.T) {
// 	t.Run("tests", func(t *testing.T) {
// 		data, _ := json.Marshal(map[string]string{"email": "art@art", "username": "art", "password": "art"})
// 		req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader(string(data)))
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		rr := httptest.NewRecorder()
// 		userHandler := NewUserHandler()
// 		userHandler.Login(rr, req)
// 		assert.Equal(t, 201, rr.Code)
// 	})

// }

// func TestLoginErrors(t *testing.T) {
// 	for _, c := range casesLogin {
// 		t.Run("tests", func(t *testing.T) {
// 			data, _ := json.Marshal(c.data)
// 			req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader(string(data)))
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			rr := httptest.NewRecorder()
// 			userHandler := NewUserHandler()
// 			userHandler.Login(rr, req)
// 			assert.Equal(t, c.wantCode, rr.Code)
// 		})
// 	}
// }

// func TestLoginErr400(t *testing.T) {
// 	t.Run("tests", func(t *testing.T) {
// 		req, err := http.NewRequest("POST", conf.PathSignUp, strings.NewReader("ASDSAD"))
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		rr := httptest.NewRecorder()
// 		userHandler := NewUserHandler()
// 		userHandler.Login(rr, req)
// 		assert.Equal(t, 400, rr.Code)
// 	})
// }
