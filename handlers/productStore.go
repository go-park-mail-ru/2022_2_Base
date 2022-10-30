package handlers

import (
	"database/sql"
	baseErrors "serv/errors"
	"serv/model"
)

type ProductStore struct {
	DB *sql.DB
}

func (ps *ProductStore) GetProductsFromStore() ([]*model.Product, error) {

	products := []*model.Product{}
	rows, err := ps.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, baseErrors.ErrServerError500
	}
	defer rows.Close()

	for rows.Next() {
		dat := model.Product{}
		err := rows.Scan(&dat.ID, &dat.Name, &dat.Description, &dat.Price, &dat.DiscountPrice, &dat.Rating, &dat.Imgsrc)
		if err != nil {
			return nil, baseErrors.ErrServerError500
		}
		products = append(products, &dat)
	}

	return products, nil
}
