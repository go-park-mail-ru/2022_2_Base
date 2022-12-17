package delivery

import (
	"net/http"
	"net/http/httptest"

	//"net/http/httptest"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	conf "serv/config"
	"serv/domain/model"
	auth "serv/microservices/auth/gen_files"
	mocks "serv/mocks"
	//mocks "serv/mocks"
	//"github.com/mailru/easyjson/jwriter"
)

// type authenticationMiddlewareType struct {
// 	userUsecase mocks.MockUserUsecaseInterface
// }

func TestGetUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	//authMiddleware := mocks.NewMockAuthenticationMiddlewareInterface(ctrl)

	//testUser := model.UserCreateParams{Email: "art@art", Username: "art", Password: "12345678"}
	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	// testUserDB := new(model.UserDB)
	// err = faker.FakeData(testUserDB)
	// assert.NoError(t, err)
	// testUserDB.ID = testUserProfile.ID
	// testUserDB.Email = testUserProfile.Email
	// salt := []byte("Base2022")
	// hashedPass := hashPass(salt, "12345678")
	// b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	// testUserDB.Password = b64Pass
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)

	//ok
	//userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	//url := conf.PathSessions
	//req, err := http.NewRequest("GET", url, nil)

	//userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	//userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(*testUserDB, nil)
	//userUsecaseMock.EXPECT().GetAddressesByUserID(testUserDB.ID).Return(testUserProfile.Address, nil)
	//userUsecaseMock.EXPECT().GetPaymentMethodByUserID(testUserDB.ID).Return(testUserProfile.PaymentMethods, nil)
	userHandler := NewUserHandler(userUsecaseMock)

	url := "/api/v1/user" + conf.PathProfile
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()

	//authMiddleware.EXPECT().CheckAuthMiddleware(myRouter).Return(myRouter)

	userHandler.GetUser(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	//router.ServeHTTP(w, req)
	// myRouter := mux.NewRouter()
	// // amw := authenticationMiddlewareType{*userUsecaseMock}
	// myRouter.Use(authMiddleware.CheckAuthMiddleware)
	// myRouter.HandleFunc(url, userHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)

	// authMiddleware.EXPECT().CheckAuthMiddleware(myRouter).Return(myRouter)

	// //next.ServeHTTP(w, r.WithContext(WithUser(r.Context(), &userData)))

	// //myRouter.ServeHTTP(rr, req)
	// myRouter.ServeHTTP(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)
	//expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	//assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")
}

func TestChangeProfile(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	//authMiddleware := mocks.NewMockAuthenticationMiddlewareInterface(ctrl)

	//testUser := model.UserCreateParams{Email: "art@art", Username: "art", Password: "12345678"}
	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	// testUserDB := new(model.UserDB)
	// err = faker.FakeData(testUserDB)
	// assert.NoError(t, err)
	// testUserDB.ID = testUserProfile.ID
	// testUserDB.Email = testUserProfile.Email
	// salt := []byte("Base2022")
	// hashedPass := hashPass(salt, "12345678")
	// b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	// testUserDB.Password = b64Pass
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)

	//ok
	//userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	//url := conf.PathSessions
	//req, err := http.NewRequest("GET", url, nil)

	//userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	//userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(*testUserDB, nil)
	//userUsecaseMock.EXPECT().GetAddressesByUserID(testUserDB.ID).Return(testUserProfile.Address, nil)
	//userUsecaseMock.EXPECT().GetPaymentMethodByUserID(testUserDB.ID).Return(testUserProfile.PaymentMethods, nil)
	userHandler := NewUserHandler(userUsecaseMock)

	url := "/api/v1/user" + conf.PathProfile
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()

	//authMiddleware.EXPECT().CheckAuthMiddleware(myRouter).Return(myRouter)

	userHandler.GetUser(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	//router.ServeHTTP(w, req)
	// myRouter := mux.NewRouter()
	// // amw := authenticationMiddlewareType{*userUsecaseMock}
	// myRouter.Use(authMiddleware.CheckAuthMiddleware)
	// myRouter.HandleFunc(url, userHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)

	// authMiddleware.EXPECT().CheckAuthMiddleware(myRouter).Return(myRouter)

	// //next.ServeHTTP(w, r.WithContext(WithUser(r.Context(), &userData)))

	// //myRouter.ServeHTTP(rr, req)
	// myRouter.ServeHTTP(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)
	//expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	//assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")
}
