package delivery

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	//"net/http/httptest"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	conf "serv/config"
	baseErrors "serv/domain/errors"
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
	userHandler := NewUserHandler(userUsecaseMock)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	testQueryUserProfile := new(model.UserProfile)
	err = faker.FakeData(testQueryUserProfile)
	assert.NoError(t, err)
	testQueryUserProfile.ID = testUserProfile.ID
	testQueryUserProfile.Address = nil
	testQueryUserProfile.PaymentMethods = nil
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)

	//ok
	userUsecaseMock.EXPECT().ChangeUser(testUserProfile, testQueryUserProfile).Return(nil)
	userUsecaseMock.EXPECT().ChangeEmail(testsessID.ID, testQueryUserProfile.Email).Return(nil)

	url := "/api/v1/user" + conf.PathProfile
	data, _ := json.Marshal(testQueryUserProfile)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
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

	userHandler.ChangeProfile(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 401 no cookie
	data, _ = json.Marshal(testQueryUserProfile)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	userHandler.ChangeProfile(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 401, rr.Code)

	//err 500 db
	userUsecaseMock.EXPECT().ChangeUser(testUserProfile, testQueryUserProfile).Return(baseErrors.ErrServerError500)

	data, _ = json.Marshal(testQueryUserProfile)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()

	userHandler.ChangeProfile(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 500 mircoservice
	userUsecaseMock.EXPECT().ChangeUser(testUserProfile, testQueryUserProfile).Return(nil)
	userUsecaseMock.EXPECT().ChangeEmail(testsessID.ID, testQueryUserProfile.Email).Return(baseErrors.ErrServerError500)

	data, _ = json.Marshal(testQueryUserProfile)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()

	userHandler.ChangeProfile(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}

func TestChangePassword(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	userHandler := NewUserHandler(userUsecaseMock)

	testUserProfile := new(model.UserProfile)
	err := faker.FakeData(testUserProfile)
	assert.NoError(t, err)
	//testQueryChangePassword := new(model.ChangePassword)
	//err = faker.FakeData(testQueryChangePassword)
	//assert.NoError(t, err)
	//testQueryUserProfile.ID = testUserProfile.ID
	//testQueryUserProfile.Address = nil
	//testQueryUserProfile.PaymentMethods = nil
	// testUserDB := new(model.UserDB)
	// err = faker.FakeData(testUserDB)
	// assert.NoError(t, err)
	// testUserDB.ID = testUserProfile.ID
	// testUserDB.Email = testUserProfile.Email

	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)

	var mockOldPass string = "123456"
	var mockNewPass string = "12345678"
	testUserDB := model.UserDB{ID: testUserProfile.ID, Username: testUserProfile.Username, Email: testUserProfile.Email, Phone: &testUserProfile.Phone, Avatar: &testUserProfile.Avatar}
	salt := []byte("Base2022")
	hashedPass := hashPass(salt, mockOldPass)
	b64OldPass := base64.RawStdEncoding.EncodeToString(hashedPass)
	testUserDB.Password = b64OldPass
	hashedNewPass := hashPass(salt, mockNewPass)
	b64NewPass := base64.RawStdEncoding.EncodeToString(hashedNewPass)
	testQueryChangePassword := model.ChangePassword{OldPassword: mockOldPass, NewPassword: mockNewPass}

	//ok
	userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(testUserDB, nil)
	userUsecaseMock.EXPECT().ChangeUserPassword(testUserProfile.ID, b64NewPass).Return(nil)

	url := "/api/v1/user" + conf.PathPassword
	data, _ := json.Marshal(testQueryChangePassword)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
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

	userHandler.ChangePassword(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, http.StatusOK, rr.Code)

	//err 401 err validation pass
	testQueryChangePasswordErr := model.ChangePassword{OldPassword: mockOldPass, NewPassword: "12"}
	data, _ = json.Marshal(testQueryChangePasswordErr)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()

	userHandler.ChangePassword(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 401, rr.Code)

	//err 500 db
	userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(testUserDB, baseErrors.ErrServerError500)
	data, _ = json.Marshal(testQueryChangePassword)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()

	userHandler.ChangePassword(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)

	//err 401 wrong pass
	testQueryChangePasswordErr = model.ChangePassword{OldPassword: mockOldPass + "123", NewPassword: mockNewPass}
	data, _ = json.Marshal(testQueryChangePasswordErr)
	userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(testUserDB, nil)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()

	userHandler.ChangePassword(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 401, rr.Code)

	//err 500 microserv
	userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(testUserDB, nil)
	userUsecaseMock.EXPECT().ChangeUserPassword(testUserProfile.ID, b64NewPass).Return(baseErrors.ErrServerError500)
	data, _ = json.Marshal(testQueryChangePassword)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()

	userHandler.ChangePassword(rr, req.WithContext(WithUser(req.Context(), testUserProfile)))
	assert.Equal(t, 500, rr.Code)
}
