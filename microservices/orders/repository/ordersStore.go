package orders

import (
	"database/sql"
	"log"
	"serv/domain/model"
	orders "serv/microservices/orders/gen_files"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type OrderStoreInterface interface {
	MakeOrder(in *orders.MakeOrderType) error
	GetOrdersFromStore(userID int) ([]*model.Order, error)
	GetOrdersAddressFromStore(addressID int) (*model.Address, error)
	GetOrdersPaymentFromStore(paymentID int) (*model.PaymentMethod, error)
	GetOrderItemsFromStore(orderID int) ([]*model.OrderItem, error)
}

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) OrderStoreInterface {
	return &OrderStore{
		db: db,
	}
}

func (os *OrderStore) MakeOrder(in *orders.MakeOrderType) error {
	log.Println("call MakeOrder store")
	delivDate := time.Unix(in.DeliveryDate, 0)
	_, err := os.db.Exec(`UPDATE orders SET orderStatus = $1, paymentStatus = $2, addressID = $3, paymentcardID = $4, creationDate = $5, deliveryDate = $6  WHERE userID = $7 AND orderStatus = $8;`, "created", "not started", in.AddressID, in.PaymentcardID, time.Now().Format("2006.01.02 15:04:05"), delivDate.Format("2006.01.02 15:04:05"), in.UserID, "cart")
	if err != nil {
		return err
	}
	return nil
}

func (os *OrderStore) GetOrdersFromStore(userID int) ([]*model.Order, error) {
	log.Println("call orders store")
	orders := []*model.Order{}

	rows, err := os.db.Query(`SELECT id, userid, orderstatus, paymentstatus, addressid, paymentcardid, creationdate, deliverydate, promocode FROM orders WHERE userid = $1 AND orderstatus <> 'cart';`, userID)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("wwwww")
	for rows.Next() {
		dat := model.Order{}
		err := rows.Scan(&dat.ID, &dat.UserID, &dat.OrderStatus, &dat.PaymentStatus, &dat.AddressID, &dat.PaymentcardID, &dat.CreationDate, &dat.DeliveryDate, &dat.Promocode)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &dat)
	}
	if orders == nil {
		log.Println("empty")
	}
	for _, order := range orders {
		orderItems, err := os.GetOrderItemsFromStore(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = orderItems
	}
	log.Println("got orders from db, count: ", len(orders))
	return orders, nil
}

func (os *OrderStore) GetOrdersAddressFromStore(addressID int) (*model.Address, error) {
	adress := model.Address{}
	rows, err := os.db.Query(`SELECT id, city, street, house, flat, priority FROM address WHERE id  = $1`, addressID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&adress.ID, &adress.City, &adress.Street, &adress.House, &adress.Flat, &adress.Priority)
		if err != nil {
			return nil, err
		}
	}
	return &adress, nil
}

func (os *OrderStore) GetOrdersPaymentFromStore(paymentID int) (*model.PaymentMethod, error) {
	payment := model.PaymentMethod{}
	rows, err := os.db.Query(`SELECT id, paymentType, number, expiryDate, priority FROM payment WHERE id  = $1`, paymentID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&payment.ID, &payment.PaymentType, &payment.Number, &payment.ExpiryDate, &payment.Priority)
		if err != nil {
			return nil, err
		}
	}
	return &payment, nil
}

func (os *OrderStore) GetOrderItemsFromStore(orderID int) ([]*model.OrderItem, error) {
	products := []*model.OrderItem{}
	rows, err := os.db.Query(`SELECT count, pr.id, pr.name, pr.category, pr.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems JOIN orders ON orderitems.orderid=orders.id JOIN products pr ON orderitems.itemid = pr.id WHERE orderid = $1;`, orderID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var count int
		dat := model.Product{}
		err := rows.Scan(&count, &dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.NominalPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		orderItem := model.OrderItem{Count: count, Item: &dat}
		products = append(products, &orderItem)
	}
	return products, nil
}
