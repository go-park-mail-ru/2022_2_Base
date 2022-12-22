package delivery

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
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
)

func TestLogin(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	testUser := new(model.UserLogin)
	err := faker.FakeData(testUser)
	assert.NoError(t, err)
	testUserDB := new(model.UserDB)
	err = faker.FakeData(testUserDB)
	assert.NoError(t, err)
	testUserDB.Email = testUser.Email
	salt := []byte("Base2022")
	hashedPass := hashPass(salt, testUser.Password)
	b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	testUserDB.Password = b64Pass
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)

	//ok
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*testUserDB, nil)
	userUsecaseMock.EXPECT().SetSession(testUser.Email).Return(testsessID, nil)
	url := conf.PathLogin
	data, _ := json.Marshal(testUser)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	userHandler := NewSessionHandler(userUsecaseMock)
	userHandler.Login(rr, req)
	assert.Equal(t, 201, rr.Code)

	//err 400 query err
	url = conf.PathLogin
	data, _ = json.Marshal("sfdsd")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.Login(rr, req)
	assert.Equal(t, 400, rr.Code)

	//err 500 db err
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*testUserDB, baseErrors.ErrServerError500)
	url = conf.PathLogin
	data, _ = json.Marshal(testUser)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.Login(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 500 err microservice
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*testUserDB, nil)
	userUsecaseMock.EXPECT().SetSession(testUser.Email).Return(testsessID, baseErrors.ErrServerError500)
	url = conf.PathLogin
	data, _ = json.Marshal(testUser)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.Login(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 401 wrong pass
	returnedModel := model.UserDB{Email: testUserDB.Email, Username: testUserDB.Username, Password: testUserDB.Password, Phone: testUserDB.Phone, Avatar: testUserDB.Avatar}
	hashedPass = hashPass(salt, testUser.Password+"2134312")
	b64Pass = base64.RawStdEncoding.EncodeToString(hashedPass)
	returnedModel.Password = b64Pass
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(returnedModel, nil)
	url = conf.PathLogin
	data, _ = json.Marshal(testUser)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.Login(rr, req)
	assert.Equal(t, 401, rr.Code)

	//err 500 err Email
	returnedModel = model.UserDB{Email: "", Username: testUserDB.Username, Password: testUserDB.Password, Phone: testUserDB.Phone, Avatar: testUserDB.Avatar}
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(returnedModel, nil)
	url = conf.PathLogin
	data, _ = json.Marshal(testUser)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.Login(rr, req)
	assert.Equal(t, 401, rr.Code)
}

func TestSignUp(t *testing.T) {
	//t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)

	testUser := model.UserCreateParams{Email: "art@art", Username: "art", Password: "12345678"}
	testUser2 := model.UserCreateParams{Email: "art@art", Username: "art"}
	salt := []byte("Base2022")
	hashedPass := hashPass(salt, testUser.Password)
	b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	testUser2.Password = b64Pass
	testsessID := new(auth.SessionID)
	err := faker.FakeData(testsessID)
	assert.NoError(t, err)
	testUserID := 2
	testPromocode := "Az2"
	testMail := model.Mail{Type: "greeting", Username: testUser.Username, Useremail: testUser.Email}
	testMailPromo := model.Mail{Type: "promocode", Username: testUser.Username, Useremail: testUser.Email, Promocode: testPromocode}

	//ok
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*&model.UserDB{}, nil)
	userUsecaseMock.EXPECT().AddUser(&testUser2).Return(testUserID, nil)
	userUsecaseMock.EXPECT().SetSession(testUser.Email).Return(testsessID, nil)
	userUsecaseMock.EXPECT().GenPromocode(testUserID).Return(testPromocode)
	wg := sync.WaitGroup{}
	wg.Add(2)
	userUsecaseMock.EXPECT().SendMail(testMail).Do(func(arg1 interface{}) {
		defer wg.Done()
	})
	userUsecaseMock.EXPECT().SendMail(testMailPromo).Do(func(arg2 interface{}) {
		defer wg.Done()
	})

	url := conf.PathLogin
	data, _ := json.Marshal(testUser)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	userHandler := NewSessionHandler(userUsecaseMock)
	userHandler.SignUp(rr, req)
	wg.Wait()
	assert.Equal(t, 201, rr.Code)

	//err 400 query err
	url = conf.PathLogin
	data, _ = json.Marshal("fsfsf")
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.SignUp(rr, req)
	assert.Equal(t, 400, rr.Code)

	//err 500 db err
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*&model.UserDB{}, baseErrors.ErrServerError500)
	url = conf.PathLogin
	data, _ = json.Marshal(testUser)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.SignUp(rr, req)
	assert.Equal(t, 500, rr.Code)

	//err 409 user exists
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*&model.UserDB{Email: testUser.Email}, nil)
	url = conf.PathLogin
	data, _ = json.Marshal(testUser)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.SignUp(rr, req)
	assert.Equal(t, 401, rr.Code)

	//err 401 err validation
	testUser3 := model.UserCreateParams{Email: "a", Username: "art", Password: "12345678"}
	userUsecaseMock.EXPECT().GetUserByUsername(testUser3.Email).Return(*&model.UserDB{}, nil)
	url = conf.PathLogin

	data, _ = json.Marshal(testUser3)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.SignUp(rr, req)
	assert.Equal(t, 401, rr.Code)

	//err 401 err validation
	userUsecaseMock.EXPECT().GetUserByUsername(testUser.Email).Return(*&model.UserDB{}, nil)
	url = conf.PathLogin
	testUser3 = model.UserCreateParams{Email: "art@art", Username: "art", Password: "1"}
	data, _ = json.Marshal(testUser3)
	req, err = http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler = NewSessionHandler(userUsecaseMock)
	userHandler.SignUp(rr, req)
	assert.Equal(t, 401, rr.Code)
}

func TestGetSession(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)

	testUser := model.UserCreateParams{Email: "art@art", Username: "art", Password: "12345678"}
	testUserDB := new(model.UserDB)
	err := faker.FakeData(testUserDB)
	assert.NoError(t, err)
	testUserDB.Email = testUser.Email
	salt := []byte("Base2022")
	hashedPass := hashPass(salt, testUser.Password)
	b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	testUserDB.Password = b64Pass
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)
	testsess := new(auth.Session)
	err = faker.FakeData(testsess)
	assert.NoError(t, err)

	//ok
	userHandler := NewSessionHandler(userUsecaseMock)
	userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	url := conf.PathSessions
	req, err := http.NewRequest("GET", url, nil)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	userHandler.GetSession(rr, req)
	assert.Equal(t, 200, rr.Code)

	//err 401 no cookies
	url = conf.PathSessions
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler.GetSession(rr, req)
	assert.Equal(t, 401, rr.Code)

	//err 401 no sess
	userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("", baseErrors.ErrServerError500)
	url = conf.PathSessions
	req, err = http.NewRequest("GET", url, nil)
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler.GetSession(rr, req)
	assert.Equal(t, 401, rr.Code)
}

func TestLogout(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)

	testUser := model.UserCreateParams{Email: "art@art", Username: "art", Password: "12345678"}
	testUserDB := new(model.UserDB)
	err := faker.FakeData(testUserDB)
	assert.NoError(t, err)
	testUserDB.Email = testUser.Email
	salt := []byte("Base2022")
	hashedPass := hashPass(salt, testUser.Password)
	b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	testUserDB.Password = b64Pass
	testsessID := new(auth.SessionID)
	err = faker.FakeData(testsessID)
	assert.NoError(t, err)

	//ok
	userHandler := NewSessionHandler(userUsecaseMock)
	userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	userUsecaseMock.EXPECT().DeleteSession(testsessID.ID).Return(nil)
	url := conf.PathSessions
	req, err := http.NewRequest("DELETE", url, nil)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	userHandler.Logout(rr, req)
	assert.Equal(t, 200, rr.Code)

	//err 401 no cookies
	url = conf.PathSessions
	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler.Logout(rr, req)
	assert.Equal(t, 401, rr.Code)

	//err 401 no sess
	userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("", baseErrors.ErrServerError500)
	url = conf.PathSessions
	req, err = http.NewRequest("DELETE", url, nil)
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler.Logout(rr, req)
	assert.Equal(t, 401, rr.Code)

	//err 401 delete sess
	userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	userUsecaseMock.EXPECT().DeleteSession(testsessID.ID).Return(baseErrors.ErrServerError500)
	url = conf.PathSessions
	req, err = http.NewRequest("DELETE", url, nil)
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    testsessID.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	req.AddCookie(cookie)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	userHandler.Logout(rr, req)
	assert.Equal(t, 500, rr.Code)
}
