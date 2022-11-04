package repository

import (
	"database/sql"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (ps *ProductStore) GetProductsFromStore() ([]*model.Product, error) {

	products := []*model.Product{}
	rows, err := ps.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}
	defer rows.Close()

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
