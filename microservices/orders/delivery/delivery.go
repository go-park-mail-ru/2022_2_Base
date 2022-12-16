package orders

import (
	"context"
	"log"
	orders "serv/microservices/orders/gen_files"
	orderuc "serv/microservices/orders/usecase"
)

type OrderManager struct {
	orders.UnimplementedOrdersWorkerServer

	usecase orderuc.OrderUsecase
}

func NewOrdersManager(ouc *orderuc.OrderUsecase) *OrderManager {
	return &OrderManager{
		usecase: *ouc,
	}
}

func (sm *OrderManager) MakeOrder(ctx context.Context, in *orders.MakeOrderType) (*orders.Nothing, error) {
	log.Println("call MakeOrder ", in)
	err := sm.usecase.MakeOrder(ctx, in)
	if err != nil {
		log.Println("error ", err)
		return &orders.Nothing{IsSuccessful: false}, err
	}
	return &orders.Nothing{IsSuccessful: true}, nil
}

func (om *OrderManager) GetOrders(ctx context.Context, userID *orders.UserID) (*orders.OrdersResponse, error) {
	log.Println("call GetOrders ", userID)
	ordersArr, err := om.usecase.GetOrders(int(userID.UserID))
	if err != nil {
		log.Println("error ", err)
		return nil, err
	}
	log.Println("got orders count ", len(ordersArr))
	var responseOrders orders.OrdersResponse
	for _, order := range ordersArr {

		orderResponse := orders.Order{ID: int32(order.ID), UserID: int32(order.UserID), OrderStatus: order.OrderStatus, PaymentStatus: order.PaymentStatus}
		orderResponse.CreationDate = order.CreationDate.Unix()
		orderResponse.DeliveryDate = order.DeliveryDate.Unix()
		for _, prod := range order.Items {
			if prod.Item.NominalPrice == prod.Item.Price {
				prod.Item.Price = 0
			}
			orderResponse.Items = append(orderResponse.Items, &orders.CartProduct{ID: int32(prod.Item.ID), Name: prod.Item.Name, Count: int32(prod.Count), Price: prod.Item.Price, NominalPrice: prod.Item.NominalPrice, Imgsrc: prod.Item.Imgsrc})
		}

		address, err := om.usecase.GetOrdersAddress(order.AddressID)
		if err != nil {
			log.Println("error ", err)
			return nil, err
		}
		orderResponse.Address = &orders.Address{ID: int32(address.ID), City: address.City, Street: address.Street, House: address.House, Flat: address.Flat, Priority: address.Priority}
		paymentcard, err := om.usecase.GetOrdersPayment(order.PaymentcardID)
		if err != nil {
			log.Println("error ", err)
			return nil, err
		}
		orderResponse.PaymentMethod = &orders.PaymentMethod{ID: int32(paymentcard.ID), PaymentType: paymentcard.PaymentType, Number: paymentcard.Number, Priority: paymentcard.Priority}
		orderResponse.PaymentMethod.ExpiryDate = paymentcard.ExpiryDate.Unix()

		responseOrders.Orders = append(responseOrders.Orders, &orderResponse)
	}
	return &responseOrders, nil
}
