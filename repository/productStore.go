package repository

import (
	"database/sql"
	"log"
	"math"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"strconv"
	"strings"
)

type ProductStoreInterface interface {
	GetProductsFromStore(lastitemid int, count int, sort string) ([]*model.Product, error)
	GetProductsWithCategoryFromStore(category string, lastitemid int, count int, sort string) ([]*model.Product, error)
	GetProductsWithBiggestDiscountFromStore(lastitemid int, count int) ([]*model.Product, error)
	GetProductFromStoreByID(itemsID int) (*model.Product, error)
	GetProductsRatingAndCommsCountFromStore(itemsID int) (float64, int, error)
	GetProductPropertiesFromStore(itemID int, itemCategory string) ([]*model.Property, error)
	CheckIsProductInFavoritesDB(userID int, itemID int) (bool, error)
	GetProductsBySearchFromStore(search string) ([]*model.Product, error)
	GetSuggestionsFromStore(search string) ([]string, error)
	GetOrderItemsFromStore(orderID int) ([]*model.OrderItem, error)
	UpdatePricesOrderItemsInStore(userID int, category string, discount int) error
	CheckPromocodeUsage(userID int, promocode string) error
	SetPromocodeDB(userID int, promocode string) error
	CreateCart(userID int) error
	GetCart(userID int) (*model.Order, error)
	UpdateCart(userID int, items *[]int) error
	InsertItemIntoCartById(userID int, itemID int) error
	DeleteItemFromCartById(userID int, itemID int) error
	GetCommentsFromStore(productID int) ([]*model.CommentDB, error)
	CreateCommentInStore(in *model.CreateComment) error
	UpdateProductRatingInStore(itemID int) error
	GetRecommendationProductsFromStore(itemID int) ([]*model.Product, error)
	GetFavoritesDB(userID int, lastitemid int, count int, sort string) ([]*model.Product, error)
	InsertItemIntoFavoritesDB(userID int, itemID int) error
	DeleteItemFromFavoritesDB(userID int, itemID int) error
}

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) ProductStoreInterface {
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
	var rows *sql.Rows

	if sort == "priceup" {
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (price, id) > ($1, $2) ORDER BY (price, id) LIMIT $3;`, lastProduct.Price, lastitemid, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (price, id) < ($1, $2) ORDER BY (price, id) DESC LIMIT $3;`, lastProduct.Price, lastitemid, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (rating, id) > ($1, $2) ORDER BY (rating, id) ASC LIMIT $3;`, lastProduct.Rating, lastitemid, count)
	} else if sort == "ratingdown" {
		if lastitemid == 0 {
			lastitemid = 1e9
			lastProduct.Rating = 10
		}
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (rating, id) < ($1, $2) ORDER BY (rating, id) DESC LIMIT $3;`, lastProduct.Rating, lastitemid, count)
	} else {
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id > $1 ORDER BY id LIMIT $2;`, lastitemid, count)
	}

	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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
	var rows *sql.Rows
	lastProduct, err := ps.GetProductFromStoreByID(lastitemid)
	if err != nil {
		return nil, err
	}

	if sort == "priceup" {
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (price, id) > ($2, $3) ORDER BY (price, id) LIMIT $4;`, category, lastProduct.Price, lastitemid, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (price, id) < ($2, $3) ORDER BY (price, id) DESC LIMIT $4;`, category, lastProduct.Price, lastitemid, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (rating, id) > ($2, $3) ORDER BY (rating, id) ASC LIMIT $4;`, category, lastProduct.Rating, lastitemid, count)
	} else if sort == "ratingdown" {
		if lastitemid == 0 {
			lastitemid = 1e9
			lastProduct.Rating = 10
		}
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (rating, id) < ($2, $3) ORDER BY (rating, id) DESC LIMIT $4;`, category, lastProduct.Rating, lastitemid, count)
	} else {
		rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND id > $2 ORDER BY id LIMIT $3;`, category, lastitemid, count)
	}

	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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

func (ps *ProductStore) GetProductsWithBiggestDiscountFromStore(lastitemid int, count int) ([]*model.Product, error) {
	products := []*model.Product{}
	lastProduct, err := ps.GetProductFromStoreByID(lastitemid)
	if err != nil {
		return nil, err
	}
	if lastitemid == 0 {
		lastProduct.Price = 0
		lastProduct.NominalPrice = 1
		lastProduct.ID = 1e9
	}
	rows, err := ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE (1 - price/nominalprice, id) < ($1, $2) ORDER BY (1 - price/nominalprice, id) DESC LIMIT $3;`, 1-lastProduct.Price/lastProduct.NominalPrice, lastProduct.ID, count)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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
	rows, err := ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE id = $1;`, itemsID)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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
	rows, err := ps.db.Query(`SELECT COUNT(id), AVG(rating) FROM comments WHERE itemid = $1;`, itemsID)
	if err != nil {
		log.Println("err get rows: ", err)
		return 0, 0, nil
	}
	defer rows.Close()
	log.Println("got product Rating And Comms from db")
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

func (ps *ProductStore) GetProductPropertiesFromStore(itemID int, itemCategory string) ([]*model.Property, error) {
	properties := []*model.Property{}
	var rows *sql.Rows
	var err error

	propertiesDB := make([]model.Property, 6)
	switch itemCategory {
	case "monitors":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN monitors cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)
	case "phones":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN phones cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)
	case "tvs":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN tvs cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)
	case "computers":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN computers cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)
	case "watches":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN watches cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)
	case "tablets":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN tablets cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)
	case "accessories":
		rows, err = ps.db.Query(`SELECT propname1, propname2, propname3, propname4, propname5, propname6, propdesc1, propdesc2, propdesc3, propdesc4, propdesc5, propdesc6 FROM properties JOIN accessories cattable ON properties.category = $1 WHERE cattable.itemid = $2;`, itemCategory, itemID)

	}

	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("got product properties from db")
	for rows.Next() {
		err := rows.Scan(&propertiesDB[0].Name, &propertiesDB[1].Name, &propertiesDB[2].Name, &propertiesDB[3].Name, &propertiesDB[4].Name, &propertiesDB[5].Name, &propertiesDB[0].Description, &propertiesDB[1].Description, &propertiesDB[2].Description, &propertiesDB[3].Description, &propertiesDB[4].Description, &propertiesDB[5].Description)
		if err != nil {
			return nil, err
		}
		properties = append(properties, &propertiesDB[0], &propertiesDB[1], &propertiesDB[2], &propertiesDB[3], &propertiesDB[4], &propertiesDB[5])
	}
	return properties, nil
}

func (ps *ProductStore) CheckIsProductInFavoritesDB(userID int, itemID int) (bool, error) {
	rows, err := ps.db.Query(`SELECT id FROM favorites WHERE userid = $1 AND itemid = $2;`, userID, itemID)
	if err != nil {
		log.Println("err get rows: ", err)
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return false, err
		}
		if id != 0 {
			return true, nil
		}
	}
	return false, nil
}

func (ps *ProductStore) GetProductsBySearchFromStore(search string) ([]*model.Product, error) {
	products := []*model.Product{}
	searchWords := strings.Split(search, " ")
	searchWordsUnite := strings.Join(searchWords, "")
	searchLetters := strings.Split(searchWordsUnite, "")
	searchString := strings.ToLower(`%` + strings.Join(searchLetters, "%") + `%`)
	rows, err := ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE LOWER(name) LIKE $1 LIMIT 20;`, searchString)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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
	search = strconv.Quote(search)
	searchWords := strings.Split(search, " ")
	searchString := strings.ToLower(`%` + strings.Join(searchWords, " ") + `%`)
	rows, err := ps.db.Query(`SELECT name FROM products WHERE LOWER(name) LIKE $1 LIMIT 3;`, searchString)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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
	rows, err := ps.db.Query(`SELECT count, pr.id, pr.name, pr.category, orderitems.price, pr.nominalprice, pr.rating, pr.imgsrc FROM orderitems JOIN orders ON orderitems.orderid=orders.id JOIN products pr ON orderitems.itemid = pr.id WHERE orderid = $1;`, orderID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
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

func (ps *ProductStore) UpdatePricesOrderItemsInStore(userID int, category string, discount int) error {
	cart, err := ps.GetCart(userID)
	if err != nil {
		return err
	}
	orderID := cart.ID
	orderItems, err := ps.GetOrderItemsFromStore(orderID)
	if err != nil {
		return err
	}
	for _, item := range orderItems {
		_, err = ps.db.Exec(`UPDATE orderItems SET price = (SELECT price FROM products WHERE id = $1) WHERE orderID = $2 AND itemID = $3;`, item.Item.ID, orderID, item.Item.ID)
		if err != nil {
			return err
		}
		switch category {
		case "all":
			_, err = ps.db.Exec(`UPDATE orderItems SET price = $1 WHERE orderID = $2 AND itemID = $3;`, math.Min(math.Ceil(item.Item.NominalPrice*float64(100-discount)/100), item.Item.Price), orderID, item.Item.ID)
			if err != nil {
				return err
			}
		default:
			if item.Item.Category == category {
				_, err = ps.db.Exec(`UPDATE orderItems SET price = $1 WHERE orderID = $2 AND itemID = $3;`, math.Min(math.Ceil(item.Item.NominalPrice*float64(100-discount)/100), item.Item.Price), orderID, item.Item.ID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ps *ProductStore) CheckPromocodeUsage(userID int, promocode string) error {
	prom := strconv.Quote(promocode)
	rows, err := ps.db.Query(`SELECT id, userid, promocode FROM usedpromocodes WHERE userid = $1 AND promocode = $2;`, userID, prom)
	if err != nil {
		log.Println("err get rows: ", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var usid int
		var prom string
		err := rows.Scan(&id, &usid, &prom)
		if err != nil {
			return err
		}
		if id != 0 {
			return baseErrors.ErrConflict409
		}
	}
	return nil
}

func (ps *ProductStore) SetPromocodeDB(userID int, promocode string) error {
	cart, err := ps.GetCart(userID)
	promocode = strconv.Quote(promocode)
	if err != nil {
		return err
	}
	if promocode == "" {
		_, err = ps.db.Exec(`DELETE FROM usedpromocodes WHERE userid = $1 AND promocode = $2;`, userID, cart.Promocode)
		if err != nil {
			return err
		}
		_, err = ps.db.Exec(`UPDATE orders SET promocode = $1 WHERE id = $2;`, nil, cart.ID)
		if err != nil {
			return err
		}
		return nil
	} else if cart.Promocode != nil {
		_, err = ps.db.Exec(`UPDATE usedpromocodes SET promocode = $1 WHERE userid = $2 AND promocode = $3;`, promocode, userID, cart.Promocode)
		if err != nil {
			return err
		}
	} else {
		_, err = ps.db.Exec(`INSERT INTO usedpromocodes (userid, promocode) VALUES ($1, $2);`, userID, promocode)
		if err != nil {
			return err
		}
	}
	_, err = ps.db.Exec(`UPDATE orders SET promocode = $1 WHERE id = $2;`, promocode, cart.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) CreateCart(userID int) error {
	_, err := ps.db.Exec(`INSERT INTO orders (userID, orderStatus, paymentStatus, addressID, paymentcardID) VALUES ($1, $2, $3, 1, 1);`, userID, "cart", "not started")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetCart(userID int) (*model.Order, error) {
	rows, err := ps.db.Query(`SELECT ID, userID, orderStatus, paymentStatus, addressID, paymentcardID, creationDate, deliveryDate, promocode  FROM orders WHERE userID = $1 AND orderStatus = $2;`, userID, "cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cart := &model.Order{}
	for rows.Next() {
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.OrderStatus, &cart.PaymentStatus, &cart.AddressID, &cart.PaymentcardID, &cart.CreationDate, &cart.DeliveryDate, &cart.Promocode)
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
	_, err = ps.db.Exec(`DELETE FROM orderItems WHERE orderID = $1;`, cart.ID)
	if err != nil {
		return err
	}
	for _, item := range *items {
		err = ps.InsertItemIntoCartById(userID, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *ProductStore) InsertItemIntoCartById(userID int, itemID int) error {
	cart, err := ps.GetCart(userID)
	if err != nil {
		return err
	}
	for _, prod := range cart.Items {
		if prod.Item.ID == itemID {
			_, err = ps.db.Exec(`UPDATE orderItems SET count = count+1 WHERE orderID = $1 AND itemID = $2;`, cart.ID, itemID)
			if err != nil {
				return err
			}
			return nil
		}
	}
	product, err := ps.GetProductFromStoreByID(itemID)
	if err != nil {
		return nil
	}
	_, err = ps.db.Exec(`INSERT INTO orderItems (itemID, orderID, price, count) VALUES ($1, $2, $3, $4);`, itemID, cart.ID, product.Price, 1)
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
	for _, prod := range cart.Items {
		if prod.Item.ID == itemID {
			if prod.Count != 1 {
				_, err = ps.db.Exec(`UPDATE orderItems SET count = count-1 WHERE orderID = $1 AND itemID = $2;`, cart.ID, itemID)
				if err != nil {
					return err
				}
				return nil
			}

			_, err = ps.db.Exec(`DELETE FROM orderItems WHERE itemID = $1 AND orderID = $2;`, itemID, cart.ID)
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
	rows, err := ps.db.Query(`SELECT userid, pros, cons, comment, rating FROM comments WHERE itemid = $1;`, productID)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
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
	pros := strconv.Quote(in.Pros)
	cons := strconv.Quote(in.Cons)
	comment := strconv.Quote(in.Comment)
	_, err := ps.db.Exec(`INSERT INTO comments (itemID, userID, pros, cons, comment, rating) VALUES ($1, $2, $3, $4, $5, $6);`, in.ItemID, in.UserID, pros, cons, comment, in.Rating)
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
	_, err = ps.db.Exec(`UPDATE products SET rating = $1 WHERE id = $2;`, rating, itemID)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetRecommendationProductsFromStore(itemID int) ([]*model.Product, error) {
	products := []*model.Product{}
	product, err := ps.GetProductFromStoreByID(itemID)
	if err != nil {
		return nil, err
	}
	var rows *sql.Rows
	lastitemid := 1e9
	lastProductRating := 10
	categoryproductsCount := 10
	accessoriesCount := 20

	// products from same category
	rows, err = ps.db.Query(`SELECT id, name, category, price, nominalprice, rating, imgsrc FROM products WHERE category = $1 AND (rating, id) < ($2, $3) ORDER BY (rating, id) DESC LIMIT $4;`, product.Category, lastProductRating, lastitemid, categoryproductsCount)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("got products from db 1")
	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Category, &dat.Price, &dat.NominalPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, err
		}
		products = append(products, &dat)
	}
	// accessories for product
	rows, err = ps.db.Query(`SELECT products.id, name, products.category, price, nominalprice, rating, imgsrc FROM products JOIN accessories ON products.id=accessories.itemID WHERE accessories.category = $1 ORDER BY (rating, products.id) DESC LIMIT $2;`, product.Category, accessoriesCount)
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("got products from db 2")
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

func (ps *ProductStore) GetFavoritesDB(userID int, lastitemid int, count int, sort string) ([]*model.Product, error) {
	products := []*model.Product{}
	lastProduct, err := ps.GetProductFromStoreByID(lastitemid)
	if err != nil {
		return nil, err
	}
	var rows *sql.Rows

	if sort == "priceup" {
		rows, err = ps.db.Query(`SELECT products.id, name, category, price, nominalprice, rating, imgsrc FROM products JOIN favorites ON products.id = favorites.itemid WHERE (price, products.id) > ($1, $2) AND userid = $3 ORDER BY (price, products.id) LIMIT $4;`, lastProduct.Price, lastitemid, userID, count)
	} else if sort == "pricedown" {
		if lastProduct.Price == 0 {
			lastProduct.Price = 1e10
		}
		rows, err = ps.db.Query(`SELECT products.id, name, category, price, nominalprice, rating, imgsrc FROM products JOIN favorites ON products.id = favorites.itemid WHERE (price, products.id) < ($1, $2) AND userid = $3 ORDER BY (price, products.id) DESC LIMIT $4;`, lastProduct.Price, lastitemid, userID, count)
	} else if sort == "ratingup" {
		rows, err = ps.db.Query(`SELECT products.id, name, category, price, nominalprice, rating, imgsrc FROM products JOIN favorites ON products.id = favorites.itemid WHERE (rating, products.id) > ($1, $2) AND userid = $3 ORDER BY (rating, products.id) ASC LIMIT $4;`, lastProduct.Rating, lastitemid, userID, count)
	} else if sort == "ratingdown" {
		if lastitemid == 0 {
			lastitemid = 1e9
			lastProduct.Rating = 10
		}
		rows, err = ps.db.Query(`SELECT products.id, name, category, price, nominalprice, rating, imgsrc FROM products JOIN favorites ON products.id = favorites.itemid WHERE (rating, products.id) < ($1, $2) AND userid = $3 ORDER BY (rating, products.id) DESC LIMIT $4;`, lastProduct.Rating, lastitemid, userID, count)
	} else {
		rows, err = ps.db.Query(`SELECT products.id, name, category, price, nominalprice, rating, imgsrc FROM products JOIN favorites ON products.id = favorites.itemid WHERE products.id > $1 AND userid = $2 LIMIT $3;`, lastitemid, userID, count)
	}

	defer rows.Close()
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, err
	}
	log.Println("got favorites from db")
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

func (ps *ProductStore) InsertItemIntoFavoritesDB(userID int, itemID int) error {
	id := 0
	rows, err := ps.db.Query(`SELECT id FROM favorites WHERE userid = $1 AND itemid = $2`, userID, itemID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
	}
	if id != 0 {
		return nil
	}
	_, err = ps.db.Exec(`INSERT INTO favorites (userID, itemID) VALUES ($1, $2);`, userID, itemID)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) DeleteItemFromFavoritesDB(userID int, itemID int) error {
	_, err := ps.db.Exec(`DELETE FROM favorites WHERE userid = $1 AND itemid = $2;`, userID, itemID)
	if err != nil {
		return err
	}
	return nil
}
