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
	lastProduct, err := ps.GetProductFromStoreByID(lastitemid)
	if err != nil {
		return nil, err
	}
	var rows pgx.Rows
	if sort == "" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE id > $1 LIMIT $2;`, lastitemid, count)
	} else if sort == "priceup" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE (price, id) > ($1, $2) ORDER BY price LIMIT $3;`, lastProduct.Price, lastitemid, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE (price, id) < ($1, $2) ORDER BY (price, id) DESC LIMIT $3;`, lastProduct.Price, lastitemid, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE (rating, id) > ($1, $2) ORDER BY rating ASC LIMIT $3;`, lastProduct.Rating, lastitemid, count)
	} else if sort == "ratingdown" {
		if lastProduct.Rating == 0 {
			lastProduct.Rating = 6
		}
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE (rating, id) < ($1, $2) ORDER BY (rating, id) DESC LIMIT $3;`, lastProduct.Rating, lastitemid, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
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
	lastProduct, err := ps.GetProductFromStoreByID(lastitemid)
	if err != nil {
		return nil, err
	}
	if sort == "" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND id > $2 LIMIT $3;`, category, lastitemid, count)
	} else if sort == "priceup" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND (price, id) > ($2, $3) ORDER BY price LIMIT $4;`, category, lastProduct.Price, lastitemid, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND (price, id) < ($2, $3) ORDER BY price DESC LIMIT $4;`, category, lastProduct.Price, lastitemid, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND (rating, id) > ($2, $3) ORDER BY rating ASC LIMIT $4;`, category, lastProduct.Rating, lastitemid, count)
	} else if sort == "ratingdown" {
		if lastProduct.Rating == 0 {
			lastProduct.Rating = 6
		}
		rows, err = ps.db.Query(context.Background(), `SELECT * FROM products WHERE category = $1 AND (rating, id) < ($2, $3) ORDER BY (rating, id) DESC LIMIT $4;`, category, lastProduct.Rating, lastitemid, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
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

func (ps *ProductStore) GetProductFromStoreByID(itemsID int) (*model.Product, error) {
	product := model.Product{}
	rows, err := ps.db.Query(context.Background(), `SELECT * FROM products WHERE id = $1;`, itemsID)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got product by id from db")
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.DiscountPrice, &product.Rating, &product.Imgsrc)
		if err != nil {
			return nil, err
		}
	}
	return &product, nil
}

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
	_, err := ps.db.Exec(context.Background(), `INSERT INTO orders (userID, orderStatus, paymentStatus, addressID, paymentcardID) VALUES ($1, $2, $3, 1, 1);`, userID, "cart", "not started")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetCart(userID int) (*model.Order, error) {
	rows, err := ps.db.Query(context.Background(), `SELECT ID, userID, orderStatus, paymentStatus, addressID, paymentcardID, creationDate, deliveryDate  FROM orders WHERE userID = $1 AND orderStatus = $2;`, userID, "cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cart := &model.Order{}
	for rows.Next() {
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.OrderStatus, &cart.PaymentStatus, &cart.AddressID, &cart.PaymentcardID, &cart.CreationDate, &cart.DeliveryDate)
		if err != nil {
			return nil, err
		}
		if cart.ID == 0 {
			err = ps.CreateCart(userID)
			if err != nil {
				return nil, err
			}
			cart, err = ps.GetCart(userID)
			if err != nil {
				return nil, err
			}
		}
	}

	orderItems, err := ps.GetOrderItemsFromStore(cart.ID)
	if err != nil {
		return nil, err
	}
	cart.Items = orderItems
	return cart, nil
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
			_, err = ps.db.Exec(context.Background(), `UPDATE orderItems SET count = count+1 WHERE orderID = $1 AND itemID = $2;`, cart.ID, itemID)
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
				_, err = ps.db.Exec(context.Background(), `UPDATE orderItems SET count = count-1 WHERE orderID = $1 AND itemID = $2;`, cart.ID, itemID)
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
	_, err := ps.db.Exec(context.Background(), `UPDATE orders SET orderStatus = $1, paymentStatus = $2, addressID = $3, paymentcardID = $4, creationDate = $5, deliveryDate = $6  WHERE userID = $7 AND orderStatus = $8;`, "created", "not started", in.AddressID, in.PaymentcardID, time.Now().Format("2006.01.02 15:04:05"), in.DeliveryDate, in.UserID, "cart")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetOrdersFromStore(userID int) ([]*model.Order, error) {
	orders := []*model.Order{}
	rows, err := ps.db.Query(context.Background(), `SELECT * FROM orders WHERE userid = $1 AND orderstatus <> 'cart';`, userID)

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got orders from db")
	for rows.Next() {
		dat := model.Order{}
		err := rows.Scan(&dat.ID, &dat.UserID, &dat.OrderStatus, &dat.PaymentStatus, &dat.AddressID, &dat.PaymentcardID, &dat.CreationDate, &dat.DeliveryDate)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &dat)
	}
	for _, order := range orders {
		orderItems, err := ps.GetOrderItemsFromStore(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = orderItems
	}
	return orders, nil
}

func (us *ProductStore) GetOrdersAddressFromStore(addressID int) (*model.Address, error) {
	adress := model.Address{}
	rows, err := us.db.Query(context.Background(), `SELECT id, city, street, house, priority FROM address WHERE id  = $1`, addressID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&adress.ID, &adress.City, &adress.Street, &adress.House, &adress.Priority)
		if err != nil {
			return nil, err
		}
	}
	return &adress, nil
}

func (us *ProductStore) GetOrdersPaymentFromStore(paymentID int) (*model.PaymentMethod, error) {
	payment := model.PaymentMethod{}
	rows, err := us.db.Query(context.Background(), `SELECT id, paymentType, number, expiryDate, priority FROM payment WHERE id  = $1`, paymentID)
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

func (ps *ProductStore) GetCommentsFromStore(productID int) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	rows, err := ps.db.Query(context.Background(), `SELECT * FROM comments WHERE itemid = $1;`, productID)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got comments from db")
	for rows.Next() {
		dat := model.Comment{}
		err := rows.Scan(&dat.ID, &dat.ItemID, &dat.UserID, &dat.Worths, &dat.Drawbacks, &dat.Comment, &dat.Rating)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &dat)
	}
	return comments, nil
}

func (ps *ProductStore) CreateCommentInStore(in *model.Comment) error {
	_, err := ps.db.Exec(context.Background(), `INSERT INTO comments (itemID, userID, worths, drawbacks, comment, rating) VALUES ($1, $2, $3, $4, $5, 0);`, in.ItemID, in.UserID, in.Worths, in.Drawbacks, in.Comment)
	if err != nil {
		return err
	}
	return nil
}
