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

	// conn, err := ps.db.Acquire(context.Background())
	// if err != nil {
	// 	log.Println("Unable to acquire a database connection: ", err)
	// 	return nil, err
	// }
	// defer conn.Release()

	rows, err := ps.db.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		log.Println("err get rows: ", err)
		return nil, baseErrors.ErrServerError500
	}
	log.Println("got rows")
	//defer rows.Close()

	for rows.Next() {
		dat := model.Product{}
		// err := rows.Scan(&dat.ID, &dat.Name, &dat.Description, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
		// if err != nil {
		// 	return nil, err
		// }
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset ", err)
			return nil, baseErrors.ErrServerError500
		}
		dat.ID = values[0].(uint)
		dat.Name = values[1].(string)
		dat.Description = values[2].(string)
		dat.Price = values[3].(float64)
		dat.DiscountPrice = values[4].(float64)
		dat.Imgsrc = values[5].(string)
		products = append(products, &dat)
	}

	return products, nil
}
