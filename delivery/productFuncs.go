package delivery

import (
	"encoding/json"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	usecase "serv/usecase"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(puc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: *puc,
	}
}

// GetHomePage godoc
// @Summary Gets products for main page
// @Description Gets products for main page
// @ID getMain
// @Accept  json
// @Produce  json
// @Tags Products
// @Success 200 {object} model.Product
// @Failure 404 {object} model.Error "Products not found"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /products [get]
func (api *ProductHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	products, err := api.usecase.GetProducts()
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if len(products) == 0 {
		ReturnErrorJSON(w, baseErrors.ErrNotFound404, 404)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{Body: products})
}
