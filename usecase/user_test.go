package usecase

import (
	"context"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	auth "serv/microservices/auth/gen_files"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "serv/mocks"
)

func TestAddUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	userUsecase := NewUserUsecase(userStoreMock, sessManager, mailManager)

	testUser := new(model.UserDB)
	err := faker.FakeData(testUser)
	assert.NoError(t, err)

	testUser.ID = 0
	testUser.Phone = nil
	testUser.Avatar = nil

	//ok
	userStoreMock.EXPECT().AddUser(testUser).Return(testUser.ID, nil)
	us := model.UserCreateParams{Email: testUser.Email, Username: testUser.Username, Password: testUser.Password}
	_, err = userUsecase.AddUser(&us)
	assert.NoError(t, err)
}

func TestSessions(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	userUsecase := NewUserUsecase(userStoreMock, sessManager, mailManager)

	testSessID := new(auth.SessionID)
	err := faker.FakeData(testSessID)
	assert.NoError(t, err)
	usEmail := "art@art"

	sessManager.EXPECT().Create(context.Background(),
		&auth.Session{
			Login: usEmail,
		}).Return(testSessID, nil)
	id, err := userUsecase.SetSession(usEmail)
	assert.NoError(t, err)
	assert.Equal(t, id, testSessID)

	//ok check
	testSess := new(auth.Session)
	err = faker.FakeData(testSess)
	assert.NoError(t, err)

	sessManager.EXPECT().Check(context.Background(),
		&auth.SessionID{
			ID: testSessID.ID,
		}).Return(testSess, nil)
	usname, err := userUsecase.CheckSession(testSessID.ID)
	assert.NoError(t, err)
	assert.Equal(t, usname, testSess.Login)

	//err check
	sessManager.EXPECT().Check(context.Background(),
		&auth.SessionID{
			ID: testSessID.ID,
		}).Return(nil, baseErrors.ErrServerError500)
	_, err = userUsecase.CheckSession(testSessID.ID)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	//ok delete
	testNoth := new(auth.Nothing)
	err = faker.FakeData(testNoth)
	assert.NoError(t, err)

	sessManager.EXPECT().Delete(context.Background(),
		&auth.SessionID{
			ID: testSessID.ID,
		}).Return(testNoth, nil)
	err = userUsecase.DeleteSession(testSessID.ID)
	assert.NoError(t, err)

	//err delete
	sessManager.EXPECT().Delete(context.Background(),
		&auth.SessionID{
			ID: testSessID.ID,
		}).Return(nil, baseErrors.ErrServerError500)
	err = userUsecase.DeleteSession(testSessID.ID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetUserByUsername(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	userUsecase := NewUserUsecase(userStoreMock, sessManager, mailManager)

	testUser := new(model.UserDB)
	err := faker.FakeData(testUser)
	assert.NoError(t, err)

	//ok
	userStoreMock.EXPECT().GetUserByUsernameFromDB(testUser.Email).Return(testUser, nil)
	user, err := userUsecase.GetUserByUsername(testUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, &user, testUser)

	//err
	userStoreMock.EXPECT().GetUserByUsernameFromDB(testUser.Email).Return(nil, baseErrors.ErrServerError500)
	_, err = userUsecase.GetUserByUsername(testUser.Email)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestAdresses(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	userUsecase := NewUserUsecase(userStoreMock, sessManager, mailManager)

	testUser := new(model.UserProfile)
	err := faker.FakeData(testUser)
	assert.NoError(t, err)

	//ok get
	userStoreMock.EXPECT().GetAddressesByUserIDFromDB(testUser.ID).Return(testUser.Address, nil)
	adresses, err := userUsecase.GetAddressesByUserID(testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, adresses, testUser.Address)

	//err get
	userStoreMock.EXPECT().GetAddressesByUserIDFromDB(testUser.ID).Return(nil, baseErrors.ErrServerError500)
	_, err = userUsecase.GetAddressesByUserID(testUser.ID)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	testAddr := new([4]*model.Address)
	err = faker.FakeData(testAddr)
	assert.NoError(t, err)
	testAddrSlice := testAddr[:]

	//ok change
	for _, addr := range testAddrSlice {
		if addr == nil {
			continue
		}
		flag := true
		for _, addrFromDB := range testUser.Address {
			if addr.ID == addrFromDB.ID {
				userStoreMock.EXPECT().UpdateUsersAddress(addr.ID, addr).Return(nil)
				flag = false
				break
			}
		}
		if flag {
			userStoreMock.EXPECT().AddUsersAddress(testUser.ID, addr).Return(nil)
		}
	}
	for _, addrFromDB := range testUser.Address {
		flag := true
		for _, addr := range testAddrSlice {
			if addr == nil {
				continue
			}
			if addr.ID == addrFromDB.ID {
				flag = false
				break
			}
		}
		if flag {
			userStoreMock.EXPECT().DeleteUsersAddress(addrFromDB.ID).Return(nil)
		}
	}

	err = userUsecase.ChangeUserAddresses(testUser.ID, testUser.Address, testAddrSlice)
	assert.NoError(t, err)
	//assert.Equal(t, adresses, testUser.Address)
}

func TestPayments(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	userUsecase := NewUserUsecase(userStoreMock, sessManager, mailManager)

	testUser := new(model.UserProfile)
	err := faker.FakeData(testUser)
	assert.NoError(t, err)

	//ok get
	userStoreMock.EXPECT().GetPaymentMethodByUserIDFromDB(testUser.ID).Return(testUser.PaymentMethods, nil)
	payms, err := userUsecase.GetPaymentMethodByUserID(testUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, payms, testUser.PaymentMethods)

	//err get
	userStoreMock.EXPECT().GetPaymentMethodByUserIDFromDB(testUser.ID).Return(nil, baseErrors.ErrServerError500)
	_, err = userUsecase.GetPaymentMethodByUserID(testUser.ID)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	testPayms := new([4]*model.PaymentMethod)
	err = faker.FakeData(testPayms)
	assert.NoError(t, err)
	testPaymSlice := testPayms[:]

	//ok change
	for _, paym := range testPaymSlice {
		if paym == nil {
			continue
		}
		flag := true
		for _, paymFromDB := range testUser.PaymentMethods {
			if paym.ID == paymFromDB.ID {
				userStoreMock.EXPECT().UpdateUsersPayment(paym.ID, paym).Return(nil)
				flag = false
				break
			}
		}
		if flag {
			userStoreMock.EXPECT().AddUsersPayment(testUser.ID, paym).Return(nil)
		}
	}
	for _, paymFromDB := range testUser.PaymentMethods {
		flag := true
		for _, paym := range testPaymSlice {
			if paym == nil {
				continue
			}
			if paym.ID == paymFromDB.ID {
				flag = false
				break
			}
		}
		if flag {
			userStoreMock.EXPECT().DeleteUsersPayment(paymFromDB.ID).Return(nil)
		}
	}

	err = userUsecase.ChangeUserPayments(testUser.ID, testUser.PaymentMethods, testPaymSlice)
	assert.NoError(t, err)
}

func TestChangeUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStoreMock := mocks.NewMockUserStoreInterface(ctrl)
	sessManager := mocks.NewMockAuthCheckerClient(ctrl)
	mailManager := mocks.NewMockMailServiceClient(ctrl)
	userUsecase := NewUserUsecase(userStoreMock, sessManager, mailManager)

	oldUser := new(model.UserProfile)
	err := faker.FakeData(oldUser)
	assert.NoError(t, err)

	testUser := new(model.UserProfile)
	err = faker.FakeData(testUser)
	assert.NoError(t, err)
	testUser.ID = oldUser.ID

	//ok change
	if len(testUser.Address) >= 0 {
		for _, addr := range testUser.Address {
			if addr == nil {
				continue
			}
			flag := true
			for _, addrFromDB := range oldUser.Address {
				if addr.ID == addrFromDB.ID {
					userStoreMock.EXPECT().UpdateUsersAddress(addr.ID, addr).Return(nil)
					flag = false
					break
				}
			}
			if flag {
				userStoreMock.EXPECT().AddUsersAddress(testUser.ID, addr).Return(nil)
			}
		}
		for _, addrFromDB := range oldUser.Address {
			flag := true
			for _, addr := range testUser.Address {
				if addr == nil {
					continue
				}
				if addr.ID == addrFromDB.ID {
					flag = false
					break
				}
			}
			if flag {
				userStoreMock.EXPECT().DeleteUsersAddress(addrFromDB.ID).Return(nil)
			}
		}
	}

	if len(testUser.PaymentMethods) >= 0 {
		for _, paym := range testUser.PaymentMethods {
			if paym == nil {
				continue
			}
			flag := true
			for _, paymFromDB := range oldUser.PaymentMethods {
				if paym.ID == paymFromDB.ID {
					userStoreMock.EXPECT().UpdateUsersPayment(paym.ID, paym).Return(nil)
					flag = false
					break
				}
			}
			if flag {
				userStoreMock.EXPECT().AddUsersPayment(testUser.ID, paym).Return(nil)
			}
		}
		for _, paymFromDB := range oldUser.PaymentMethods {
			flag := true
			for _, paym := range testUser.PaymentMethods {
				if paym == nil {
					continue
				}
				if paym.ID == paymFromDB.ID {
					flag = false
					break
				}
			}
			if flag {
				userStoreMock.EXPECT().DeleteUsersPayment(paymFromDB.ID).Return(nil)
			}
		}
	}
	testUser2 := &model.UserProfile{Email: testUser.Email, Username: testUser.Username, Phone: testUser.Phone, Avatar: testUser.Avatar, Address: oldUser.Address, PaymentMethods: oldUser.PaymentMethods}
	userStoreMock.EXPECT().UpdateUser(oldUser.ID, testUser2).Return(nil)
	err = userUsecase.ChangeUser(oldUser, testUser)
	assert.NoError(t, err)
}
