package orders

import (
	"context"
	"log"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	orderst "serv/microservices/orders/repository"
)

type OrderUsecase struct {
	store orderst.OrderStore
}

func NewOrderUsecase(os *orderst.OrderStore) *OrderUsecase {
	return &OrderUsecase{
		store: *os,
	}
}

func (om *OrderUsecase) MakeOrder(ctx context.Context, in *orders.MakeOrderType) error {
	log.Println("call MakeOrder usecase")
	return om.store.MakeOrder(ctx, in)
}

// func (om *OrderUsecase) GetOrders(ctx context.Context, userID *orders.UserID) (*orders.OrdersResponse, error) {
// 	log.Println("call GetOrders usecase")
// 	return om.store.GetOrdersFromStore(ctx, userID)
// }

func (api *OrderUsecase) GetOrders(userID int) ([]*model.Order, error) {
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
