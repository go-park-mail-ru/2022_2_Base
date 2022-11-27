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

// func TimestampToOrdersData(in *time.Time) orders.Date {

// }

func (om *OrderManager) GetOrders(ctx context.Context, userID *orders.UserID) (*orders.OrdersResponse, error) {
	log.Println("call GetOrders ", userID)
	ordersArr, err := om.usecase.GetOrders(int(userID.UserID))
	if err != nil {
		log.Println("error ", err)
		return nil, err
	}
	var responseOrders orders.OrdersResponse
	for _, order := range ordersArr {
		log.Println(order.CreationDate)
		//newOrder := orders.Order{ID: order.ID, UserID: order.UserID, OrderStatus: order.OrderStatus, PaymentStatus: order.PaymentStatus, CreationDate: order.CreationDate, DeliveryDate: order.DeliveryDate}
		orderResponse := orders.Order{ID: int32(order.ID), UserID: int32(order.UserID), OrderStatus: order.OrderStatus, PaymentStatus: order.PaymentStatus}
		//orderResponse.CreationDate = &orders.Date{Year: int32(order.CreationDate.Year()), Month: int32(order.CreationDate.Month()), Day: int32(order.CreationDate.Day()), Hour: int32(order.CreationDate.Hour()), Minutes: int32(order.CreationDate.Minute())}
		//orderResponse.DeliveryDate = &orders.Date{Year: int32(order.DeliveryDate.Year()), Month: int32(order.DeliveryDate.Month()), Day: int32(order.DeliveryDate.Day()), Hour: int32(order.DeliveryDate.Hour()), Minutes: int32(order.DeliveryDate.Minute())}
		orderResponse.CreationDate = order.CreationDate.Unix()
		orderResponse.DeliveryDate = order.DeliveryDate.Unix()
		for _, prod := range order.Items {
			orderResponse.Items = append(orderResponse.Items, &orders.CartProduct{ID: int32(prod.Item.ID), Name: prod.Item.Name, Count: int32(prod.Count), Price: prod.Item.Price, DiscountPrice: prod.Item.DiscountPrice, Imgsrc: prod.Item.Imgsrc})
		}
		//orderResponse.Address, err = om.usecase.GetOrdersAddress(order.AddressID)
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
	//log.Println("list ", len(responseOrders.Orders))
	return &responseOrders, nil
}
