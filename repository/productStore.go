package repository

import (
	"context"
	"log"
	baseErrors "serv/domain/errors"
	"serv/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
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

	rows, err := ps.db.Query(context.Background(), "SELECT * FROM products")
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

func (ps *ProductStore) CreateCart(userID int) error {
	_, err := ps.db.Exec(context.Background(), `INSERT INTO ordertable (userID, items, orderStatus, paymentStatus, adress) VALUES ($1, $2, $3, $4, $5);`, userID, make([]int, 0), "cart", "not started", "111")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) GetCart(userID int) (*model.Order, error) {
	rows, err := ps.db.Query(context.Background(), `SELECT ID, userID, items, orderStatus, paymentStatus, adress FROM ordertable WHERE userID = $1 AND orderStatus = $2;`, userID, "cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cart := model.Order{}
	for rows.Next() {
		err := rows.Scan(&cart.ID, &cart.UserID, (*pq.StringArray)(&cart.Items), &cart.OrderStatus, &cart.PaymentStatus, &cart.Adress)
		if err != nil {
			return nil, err
		}
	}
	return &cart, nil
}

func (ps *ProductStore) UpdateCart(userID int, items *[]int) error {
	_, err := ps.db.Exec(context.Background(), `UPDATE ordertable SET items = $1 WHERE userID = $2 AND orderStatus = $3;`, items, userID, "cart")
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
