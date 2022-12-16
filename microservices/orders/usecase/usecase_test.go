package orders

import (
	"context"
	"serv/domain/model"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	baseErrors "serv/domain/errors"
	orders "serv/microservices/orders/gen_files"
	mocks "serv/mocks"
)

func TestGetOrders(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ordStoreMock := mocks.NewMockOrderStoreInterface(ctrl)

	ordersUsecase := NewOrderUsecase(ordStoreMock)
	testOrders := new([5]*model.Order)
	err := faker.FakeData(testOrders)
	assert.NoError(t, err)
	testOrdersSlice := testOrders[:]

	var userID int = 1
	//ok
	ordStoreMock.EXPECT().GetOrdersFromStore(userID).Return(testOrdersSlice, nil)
	orders, err := ordersUsecase.GetOrders(userID)
	assert.NoError(t, err)
	assert.Equal(t, testOrdersSlice, orders)

	//error
	ordStoreMock.EXPECT().GetOrdersFromStore(userID).Return(nil, baseErrors.ErrServerError500)
	_, err = ordersUsecase.GetOrders(userID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetOrdersAddress(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ordStoreMock := mocks.NewMockOrderStoreInterface(ctrl)
	ordersUsecase := NewOrderUsecase(ordStoreMock)

	testAddr := new(model.Address)
	err := faker.FakeData(testAddr)
	assert.NoError(t, err)
	var addrID int = 1

	//ok
	ordStoreMock.EXPECT().GetOrdersAddressFromStore(addrID).Return(testAddr, nil)
	addr, err := ordersUsecase.GetOrdersAddress(addrID)
	assert.NoError(t, err)
	assert.Equal(t, testAddr, &addr)

	//error
	ordStoreMock.EXPECT().GetOrdersAddressFromStore(addrID).Return(nil, baseErrors.ErrServerError500)
	_, err = ordersUsecase.GetOrdersAddress(addrID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestGetOrdersPayment(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ordStoreMock := mocks.NewMockOrderStoreInterface(ctrl)
	ordersUsecase := NewOrderUsecase(ordStoreMock)

	testPaym := new(model.PaymentMethod)
	err := faker.FakeData(testPaym)
	assert.NoError(t, err)
	var paymID int = 1

	//ok
	ordStoreMock.EXPECT().GetOrdersPaymentFromStore(paymID).Return(testPaym, nil)
	paym, err := ordersUsecase.GetOrdersPayment(paymID)
	assert.NoError(t, err)
	assert.Equal(t, testPaym, &paym)

	//error
	ordStoreMock.EXPECT().GetOrdersPaymentFromStore(paymID).Return(nil, baseErrors.ErrServerError500)
	_, err = ordersUsecase.GetOrdersPayment(paymID)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}

func TestMakeOrder(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ordStoreMock := mocks.NewMockOrderStoreInterface(ctrl)
	ordersUsecase := NewOrderUsecase(ordStoreMock)

	testOrder := new(orders.MakeOrderType)
	err := faker.FakeData(testOrder)
	assert.NoError(t, err)

	//ok
	ordStoreMock.EXPECT().MakeOrder(testOrder).Return(nil)
	err = ordersUsecase.MakeOrder(context.Background(), testOrder)
	assert.NoError(t, err)

	//error
	ordStoreMock.EXPECT().MakeOrder(testOrder).Return(baseErrors.ErrServerError500)
	err = ordersUsecase.MakeOrder(context.Background(), testOrder)
	assert.Equal(t, baseErrors.ErrServerError500, err)
}
