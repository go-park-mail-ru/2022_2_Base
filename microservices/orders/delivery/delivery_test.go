package orders

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	baseErrors "serv/domain/errors"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	mocks "serv/mocks"
)

func TestGetOrders(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	orderUsecaseMock := mocks.NewMockOrderUsecaseInterface(ctrl)
	orderHandler := NewOrdersManager(orderUsecaseMock)

	testUserID := new(orders.UserID)
	err := faker.FakeData(testUserID)
	assert.NoError(t, err)

	testOrders := new([1]*model.Order)
	err = faker.FakeData(testOrders)
	assert.NoError(t, err)
	testOrdersSlice := testOrders[:]
	testAddr := new(model.Address)
	err = faker.FakeData(testAddr)
	assert.NoError(t, err)
	testPaym := new(model.PaymentMethod)
	err = faker.FakeData(testPaym)
	assert.NoError(t, err)

	//ok
	orderUsecaseMock.EXPECT().GetOrders(int(testUserID.UserID)).Return(testOrdersSlice, nil)
	orderUsecaseMock.EXPECT().GetOrdersAddress(testOrdersSlice[0].AddressID).Return(*testAddr, nil)
	orderUsecaseMock.EXPECT().GetOrdersPayment(testOrdersSlice[0].PaymentcardID).Return(*testPaym, nil)

	orderHandler.GetOrders(context.Background(), testUserID)
	assert.NoError(t, err)

	// err 500
	orderUsecaseMock.EXPECT().GetOrders(int(testUserID.UserID)).Return(nil, baseErrors.ErrServerError500)
	_, err = orderHandler.GetOrders(context.Background(), testUserID)
	//assert.NoError(t, err)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	// err 500
	orderUsecaseMock.EXPECT().GetOrders(int(testUserID.UserID)).Return(testOrdersSlice, nil)
	orderUsecaseMock.EXPECT().GetOrdersAddress(testOrdersSlice[0].AddressID).Return(*testAddr, baseErrors.ErrServerError500)
	_, err = orderHandler.GetOrders(context.Background(), testUserID)
	assert.Equal(t, baseErrors.ErrServerError500, err)

	// err 500
	orderUsecaseMock.EXPECT().GetOrders(int(testUserID.UserID)).Return(testOrdersSlice, nil)
	orderUsecaseMock.EXPECT().GetOrdersAddress(testOrdersSlice[0].AddressID).Return(*testAddr, nil)
	orderUsecaseMock.EXPECT().GetOrdersPayment(testOrdersSlice[0].PaymentcardID).Return(*testPaym, baseErrors.ErrServerError500)
	_, err = orderHandler.GetOrders(context.Background(), testUserID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestMakeOrder(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	orderUsecaseMock := mocks.NewMockOrderUsecaseInterface(ctrl)
	orderHandler := NewOrdersManager(orderUsecaseMock)

	testUserID := new(orders.UserID)
	err := faker.FakeData(testUserID)
	assert.NoError(t, err)

	testOrder := new(orders.MakeOrderType)
	err = faker.FakeData(testOrder)
	assert.NoError(t, err)

	//ok
	orderUsecaseMock.EXPECT().MakeOrder(context.Background(), testOrder).Return(nil)
	orderHandler.MakeOrder(context.Background(), testOrder)
	assert.NoError(t, err)

	// err 500
	orderUsecaseMock.EXPECT().MakeOrder(context.Background(), testOrder).Return(baseErrors.ErrServerError500)
	_, err = orderHandler.MakeOrder(context.Background(), testOrder)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}
