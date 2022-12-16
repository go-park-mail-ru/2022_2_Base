package delivery

import (
	"net/http"
	//"net/http/httptest"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	conf "serv/config"
	auth "serv/microservices/auth/gen_files"
	//mocks "serv/mocks"
	//"github.com/mailru/easyjson/jwriter"
)

func TestGetUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//userUsecaseMock := mocks.NewMockUserUsecaseInterface(ctrl)
	//authenticationMiddleware := mocks.NewMockUserUsecaseInterface(ctrl)

	//testUser := model.UserCreateParams{Email: "art@art", Username: "art", Password: "12345678"}
	// testUserProfile := new(model.UserProfile)
	// err := faker.FakeData(testUserProfile)
	// assert.NoError(t, err)
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
	err := faker.FakeData(testsessID)
	assert.NoError(t, err)

	//ok
	//userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	//url := conf.PathSessions
	//req, err := http.NewRequest("GET", url, nil)

	//userUsecaseMock.EXPECT().CheckSession(testsessID.ID).Return("zz", nil)
	//userUsecaseMock.EXPECT().GetUserByUsername(testUserProfile.Email).Return(*testUserDB, nil)
	//userUsecaseMock.EXPECT().GetAddressesByUserID(testUserDB.ID).Return(testUserProfile.Address, nil)
	//userUsecaseMock.EXPECT().GetPaymentMethodByUserID(testUserDB.ID).Return(testUserProfile.PaymentMethods, nil)
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
	//rr := httptest.NewRecorder()
	//userHandler := NewUserHandler(userUsecaseMock)
	//userHandler.GetUser(rr, req)
	//router.ServeHTTP(w, req)
	// myRouter := mux.NewRouter()
	// amw := authenticationMiddleware{userUsecaseMock}
	// myRouter.Use(amw.checkAuthMiddleware)
	// myRouter.ServeHTTP(rr, req)
	// assert.Equal(t, http.StatusOK, rr.Code)
	//expectedstr, err := json.Marshal(&model.Response{Body: testProducts})
	//assert.Equal(t, rr.Body.String(), string(expectedstr)+"\n")

}
