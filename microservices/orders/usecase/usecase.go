package orders

import (
	"context"
	"log"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	orderst "serv/microservices/orders/repository"
)

type OrderUsecaseInterface interface {
	MakeOrder(ctx context.Context, in *orders.MakeOrderType) error
	GetOrders(userID int) ([]*model.Order, error)
	GetOrdersAddress(addressID int) (model.Address, error)
	GetOrdersPayment(paymentID int) (model.PaymentMethod, error)
}

type OrderUsecase struct {
	store orderst.OrderStoreInterface
}

func NewOrderUsecase(os orderst.OrderStoreInterface) OrderUsecaseInterface {
	return &OrderUsecase{
		store: os,
	}
}

func (om *OrderUsecase) MakeOrder(ctx context.Context, in *orders.MakeOrderType) error {
	log.Println("call MakeOrder usecase")
	return om.store.MakeOrder(in)
}

func (api *OrderUsecase) GetOrders(userID int) ([]*model.Order, error) {
	log.Println("call GetOrders usecase")
	orders, err := api.store.GetOrdersFromStore(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (api *OrderUsecase) GetOrdersAddress(addressID int) (model.Address, error) {
	address, err := api.store.GetOrdersAddressFromStore(addressID)
	if err != nil {
		return model.Address{}, err
	}
	return *address, nil
}

func (api *OrderUsecase) GetOrdersPayment(paymentID int) (model.PaymentMethod, error) {
	payment, err := api.store.GetOrdersPaymentFromStore(paymentID)
	if err != nil {
		return model.PaymentMethod{}, err
	}
	return *payment, nil
}
