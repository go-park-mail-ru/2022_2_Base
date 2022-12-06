package repository

import (
	"context"
	"log"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// type ProductStor interface {
//     db
// }

// type ProductStorePGXPOOl struct {
// 	db ProductStore
// }

type ProductStore struct {
	db *pgxpool.Pool
}

func NewProductStore(db *pgxpool.Pool) ProductStoreInterface {
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

	if sort == "priceup" {
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (price, id) > ($1, $2) ORDER BY (price, id) LIMIT $3;`, lastProduct.Price, lastitemid, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (price, id) < ($1, $2) ORDER BY (price, id) DESC LIMIT $3;`, lastProduct.Price, lastitemid, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (rating, id) > ($1, $2) ORDER BY (rating, id) ASC LIMIT $3;`, lastProduct.Rating, lastitemid, count)
	} else if sort == "ratingdown" {
		if lastitemid == 0 {
			lastitemid = 1e9
			lastProduct.Rating = 10
		}
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (rating, id) < ($1, $2) ORDER BY (rating, id) DESC LIMIT $3;`, lastProduct.Rating, lastitemid, count)
	} else {
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id > $1 LIMIT $2;`, lastitemid, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got products from db")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.NominalPrice, &dat.Rating, &dat.Imgsrc)
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
	if sort == "priceup" {
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (price, id) > ($2, $3) ORDER BY (price, id) LIMIT $4;`, category, lastProduct.Price, lastitemid, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (price, id) < ($2, $3) ORDER BY (price, id) DESC LIMIT $4;`, category, lastProduct.Price, lastitemid, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (rating, id) > ($2, $3) ORDER BY (rating, id) ASC LIMIT $4;`, category, lastProduct.Rating, lastitemid, count)
	} else if sort == "ratingdown" {
		if lastitemid == 0 {
			lastitemid = 1e9
			lastProduct.Rating = 10
		}
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (rating, id) < ($2, $3) ORDER BY (rating, id) DESC LIMIT $4;`, category, lastProduct.Rating, lastitemid, count)
	} else {
		rows, err = ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND id > $2 ORDER BY id LIMIT $3;`, category, lastitemid, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got products from db")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.NominalPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		products = append(products, &dat)
	}
	return products, nil
}

func (ps *ProductStore) GetProductFromStoreByID(itemsID int) (*model.Product, error) {
	product := model.Product{}
	rows, err := ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id = $1;`, itemsID)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got product by id from db")
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.NominalPrice, &product.Rating, &product.Imgsrc)
		if err != nil {
			return nil, err
		}
	}
	return &product, nil
}

func (ps *ProductStore) GetProductsRatingAndCommsCountFromStore(itemsID int) (float64, int, error) {
	var rating *float64
	var commsCount *int
	rows, err := ps.db.Query(context.Background(), `SELECT COUNT(id), AVG(rating) FROM comments WHERE itemid = $1;`, itemsID)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return 0, 0, nil
	}
	for rows.Next() {
		err := rows.Scan(&commsCount, &rating)
		if err != nil {
			return 0, 0, err
		}
	}
	if rating == nil || commsCount == nil {
		return 0, 0, nil
	}
	return *rating, *commsCount, nil
}

func (ps *ProductStore) GetProductsBySearchFromStore(search string) ([]*model.Product, error) {
	products := []*model.Product{}
	searchWords := strings.Split(search, " ")
	searchWordsUnite := strings.Join(searchWords, "")
	searchLetters := strings.Split(searchWordsUnite, "")
	searchString := strings.ToLower(`%` + strings.Join(searchLetters, "%") + `%`)
	rows, err := ps.db.Query(context.Background(), `SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE LOWER(name) LIKE $1 LIMIT 20;`, searchString)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got products from db")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.NominalPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		products = append(products, &dat)
	}
	return products, nil
}

func (ps *ProductStore) GetSuggestionsFromStore(search string) ([]string, error) {
	suggestions := []string{}
	searchWords := strings.Split(search, " ")
	searchString := strings.ToLower(`%` + strings.Join(searchWords, " ") + `%`)
	rows, err := ps.db.Query(context.Background(), `SELECT name FROM products WHERE LOWER(name) LIKE $1 LIMIT 3;`, searchString)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got products from db")
	for rows.Next() {
		var dat string
		err := rows.Scan(&dat)
		if err != nil {
			return nil, err
		}
		suggestions = append(suggestions, dat)
	}
	return suggestions, nil
}

func (ps *ProductStore) GetOrderItemsFromStore(orderID int) ([]*model.OrderItem, error) {
	products := []*model.OrderItem{}
	rows, err := ps.db.Query(context.Background(), `SELECT count, pr.id, pr.name, pr.category, pr.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems JOIN orders ON orderitems.orderid=orders.id JOIN products pr ON orderitems.itemid = pr.id WHERE orderid = $1;`, orderID)
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

func (ps *ProductStore) GetCommentsFromStore(productID int) ([]*model.CommentDB, error) {
	comments := []*model.CommentDB{}
	rows, err := ps.db.Query(context.Background(), `SELECT userid, pros, cons, comment, rating FROM comments WHERE itemid = $1;`, productID)
	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got comments from db")
	for rows.Next() {
		dat := model.CommentDB{}
		err := rows.Scan(&dat.UserID, &dat.Pros, &dat.Cons, &dat.Comment, &dat.Rating)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &dat)
	}
	return comments, nil
}

func (ps *ProductStore) CreateCommentInStore(in *model.CreateComment) error {
	_, err := ps.db.Exec(context.Background(), `INSERT INTO comments (itemID, userID, pros, cons, comment, rating) VALUES ($1, $2, $3, $4, $5, $6);`, in.ItemID, in.UserID, in.Pros, in.Cons, in.Comment, in.Rating)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) UpdateProductRatingInStore(itemID int) error {
	rating, _, err := ps.GetProductsRatingAndCommsCountFromStore(itemID)
	if err != nil {
		return err
	}
	_, err = ps.db.Exec(context.Background(), `UPDATE products SET rating = $1 WHERE id = $2;`, rating, itemID)
	if err != nil {
		return err
	}
	return nil
}
