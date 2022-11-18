package repository

import (
	"context"
	"log"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"time"

	"github.com/jackc/pgx/v5"
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

func (ps *ProductStore) GetProductsFromStore(lastitemid int, count int, sort string) ([]*model.Product, error) {
	products := []*model.Product{}
	var rows pgx.Rows
	var err error
	if sort == "" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE id > $1 LIMIT $2;`, lastitemid, count)
	} else if sort == "price" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE id > $1 ORDER BY price LIMIT $2;`, lastitemid, count)
	} else if sort == "rating" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE id > $1 ORDER BY rating DESC LIMIT $2;`, lastitemid, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, baseErrors.ErrServerError500
	}
	log.Println("got products from db")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		products = append(products, &dat)
	}
	return products, nil
}

func (ps *ProductStore) GetProductsWithCategoryFromStore(category string, lastitemid int, count int, sort string) ([]*model.Product, error) {
	products := []*model.Product{}
	var rows pgx.Rows
	var err error
	if sort == "" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND id > $2 LIMIT $3;`, category, lastitemid, count)
	} else if sort == "price" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND id > $2 ORDER BY price LIMIT $3;`, category, lastitemid, count)
	} else if sort == "rating" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND id > $2 ORDER BY rating DESC LIMIT $3;`, category, lastitemid, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, baseErrors.ErrServerError500
	}
	log.Println("got products from db")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
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
	rows, err := ps.db.Query(context.Background(), `SELECT count, pr.id, pr.name, pr.category, pr.price, pr.discountprice, pr.rating, pr.imgsrc FROM orderitems JOIN orders ON orderitems.orderid=orders.id JOIN products pr ON orderitems.itemid = pr.id WHERE orderid = $1;`, orderID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var count int
		dat := model.Product{}
		err := rows.Scan(&count, &dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		orderItem := model.OrderItem{Count: count, Item: &dat}
		products = append(products, &orderItem)
	}
	return products, nil
}

func (ps *ProductStore) CreateCart(userID int) error {
	_, err := ps.db.Exec(context.Background(), `INSERT INTO orders (userID, orderStatus, paymentStatus) VALUES ($1, $2, $3);`, userID, "cart", "not started")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetCart(userID int) (*model.Order, error) {
	rows, err := ps.db.Query(context.Background(), `SELECT ID, userID, orderStatus, paymentStatus, address, paymentcardnumber, creationDate, deliveryDate  FROM orders WHERE userID = $1 AND orderStatus = $2;`, userID, "cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cart := model.Order{}
	for rows.Next() {
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.OrderStatus, &cart.PaymentStatus, &cart.Address, &cart.Paymentcardnumber, &cart.CreationDate, &cart.DeliveryDate)
		if err != nil {
			return nil, err
		}
	}

	orderItems, err := ps.GetOrderItemsFromStore(cart.ID)
	if err != nil {
		return nil, err
	}
	cart.Items = orderItems
	return &cart, nil
}

func (ps *ProductStore) UpdateCart(userID int, items *[]int) error {
	cart, err := ps.GetCart(userID)
	if err != nil {
		return err
	}
	_, err = ps.db.Exec(context.Background(), `DELETE FROM orderItems WHERE orderID = $1;`, cart.ID)
	if err != nil {
		return err
	}
	for _, item := range *items {
		ps.InsertItemIntoCartById(userID, item)
	}
	return nil
}

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
			_, err = ps.db.Exec(context.Background(), `UPDATE orderItems SET count = count+1 WHERE orderID = $1;`, cart.ID)
			if err != nil {
				return err
			}
			return nil
		}
	}
	_, err = ps.db.Exec(context.Background(), `INSERT INTO orderItems (itemID, orderID, count) VALUES ($1, $2, $3);`, itemID, cart.ID, 1)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) DeleteItemFromCartById(userID int, itemID int) error {
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
			if prod.Count != 1 {
				_, err = ps.db.Exec(context.Background(), `UPDATE orderItems SET count = count-1 WHERE orderID = $1;`, cart.ID)
				if err != nil {
					return err
				}
				return nil
			}

			_, err = ps.db.Exec(context.Background(), `DELETE FROM orderItems WHERE itemID = $1 AND orderID = $2;`, itemID, cart.ID)
			if err != nil {
				return err
			}
			return nil

		}
	}
	return baseErrors.ErrNotFound404
}

func (ps *ProductStore) MakeOrder(in *model.MakeOrder) error {
	_, err := ps.db.Exec(context.Background(), `UPDATE orders SET orderStatus = $1, paymentStatus = $2, address = $3, paymentcardnumber = $4, creationDate = $5, deliveryDate = $6  WHERE userID = $7 AND orderStatus = $8;`, "created", "not started", in.Address, in.Paymentcardnumber, time.Now().Format("2006.01.02 15:04:05"), in.DeliveryDate, in.UserID, "cart")
	if err != nil {
		return err
	}
	return nil
}
