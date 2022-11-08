package repository

import (
	"context"
	"log"
	baseErrors "serv/domain/errors"
	"serv/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStore struct {
	db *pgxpool.Pool
}

func NewProductStore(db *pgxpool.Pool) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (ps *ProductStore) GetProductsFromStore() ([]*model.Product, error) {
	products := []*model.Product{}
	rows, err := ps.db.Query(context.Background(), `SELECT * FROM products LIMIT 5;`)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, baseErrors.ErrServerError500
	}
	log.Println("got products from db")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Description, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		products = append(products, &dat)
	}
	return products, nil
}

// func (ps *ProductStore) GetProductsFromStoreByIDs(itemsIDs *[]string) ([]*model.Product, error) {
// 	products := []*model.Product{}
// 	itemsIDsString := "{" + strings.Join(*itemsIDs, ",") + "}"
// 	rows, err := ps.db.Query(context.Background(), `SELECT * FROM products WHERE id = ANY($1::int[]);`, itemsIDsString)
// 	defer rows.Close()
// 	if err != nil {
// 		log.Println("err get rows: ", err)
// 		return nil, baseErrors.ErrServerError500
// 	}
// 	log.Println("got products by ids from db")
// 	for rows.Next() {
// 		dat := model.Product{}
// 		err := rows.Scan(&dat.ID, &dat.Name, &dat.Description, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, &dat)
// 	}
// 	return products, nil
// }

func (ps *ProductStore) GetOrderItemsFromStore(orderID int) ([]*model.OrderItem, error) {
	products := []*model.OrderItem{}
	//itemsIDsString := "{" + strings.Join(*itemsIDs, ",") + "}"
	rows, err := ps.db.Query(context.Background(), `SELECT count, pr.name, pr.description, pr.price, pr.discountprice, pr.rating, pr.imgsrc FROM orderitems JOIN ordertable ON orderitems.orderid=ordertable.id JOIN products pr ON orderitems.itemid = pr.id WHERE orderid = $1;`, orderID)
	defer rows.Close()
	if err != nil {
		//log.Println("err get rows: ", err)
		return nil, baseErrors.ErrServerError500
	}
	//log.Println("aaa")
	for rows.Next() {
		// dat := model.OrderItem{}
		// err := rows.Scan(&dat.Count, &dat.Item.Name, &dat.Item.Description, &dat.Item.Price, &dat.Item.DiscountPrice, &dat.Item.Rating, &dat.Item.Imgsrc)
		var count int
		dat := model.Product{}
		err := rows.Scan(&count, &dat.Name, &dat.Description, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		orderItem := model.OrderItem{Count: count, Item: &dat}
		//products = append(products, &dat)
		products = append(products, &orderItem)
	}
	return products, nil
}

func (ps *ProductStore) CreateCart(userID int) error {
	_, err := ps.db.Exec(context.Background(), `INSERT INTO ordertable (userID, orderStatus, paymentStatus, adress) VALUES ($1, $2, $3, $4);`, userID, "cart", "not started", "111")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetCart(userID int) (*model.Order, error) {
	//rows, err := ps.db.Query(context.Background(), `SELECT ID, userID, items, orderStatus, paymentStatus, adress FROM ordertable WHERE userID = $1 AND orderStatus = $2;`, userID, "cart")
	rows, err := ps.db.Query(context.Background(), `SELECT ID, userID, orderStatus, paymentStatus, adress FROM ordertable WHERE userID = $1 AND orderStatus = $2;`, userID, "cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cart := model.Order{}
	//var stringArrayOfItemsIDs []string
	for rows.Next() {
		//err := rows.Scan(&cart.ID, &cart.UserID, (*pq.StringArray)(&stringArrayOfItemsIDs), &cart.OrderStatus, &cart.PaymentStatus, &cart.Adress)
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.OrderStatus, &cart.PaymentStatus, &cart.Adress)
		if err != nil {
			return nil, err
		}
	}
	//products, err := ps.GetProductsFromStoreByIDs(&stringArrayOfItemsIDs)
	orderItems, err := ps.GetOrderItemsFromStore(cart.ID)

	if err != nil {
		return nil, err
	}
	// var orderItems []*model.OrderItem
	// for _, prod := range products {
	// 	orderItems = append(orderItems, &model.OrderItem{Item: prod, Amount: 1})
	// }
	cart.Items = orderItems
	log.Println("items ", cart.Items)
	return &cart, nil
}

// func (ps *ProductStore) UpdateCart(userID int, items *[]int) error {
// 	_, err := ps.db.Exec(context.Background(), `UPDATE ordertable SET items = $1 WHERE userID = $2 AND orderStatus = $3;`, items, userID, "cart")
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (ps *ProductStore) InsertItemIntoCartById(userID int, itemID int) error {
	cart, err := ps.GetCart(userID)
	if err != nil {
		return err
	}
	orderItems, err := ps.GetOrderItemsFromStore(cart.ID)
	if err != nil {
		return err
	}

	for _, prod := range orderItems {
		if prod.Item.ID == itemID {
			_, err = ps.db.Exec(context.Background(), `UPDATE orderItems SET count = $1 WHERE orderID = $2;`, prod.Count+1, cart.ID)
			if err != nil {
				return err
			}
			return nil
		}
	}
	// //items := cart.Items
	// //append()
	// items = append(items, itemID)
	_, err = ps.db.Exec(context.Background(), `INSERT INTO orderItems (userID, itemID, orderID, count) VALUES ($1, $2, $3, $4);`, userID, itemID, cart.ID, 1)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) MakeOrder(userID int) error {
	_, err := ps.db.Exec(context.Background(), `UPDATE ordertable SET orderStatus = $1, paymentStatus = $2  WHERE userID = $3 AND orderStatus = $4;`, "processed", "paid", userID, "cart")
	if err != nil {
		return err
	}
	return nil
}
