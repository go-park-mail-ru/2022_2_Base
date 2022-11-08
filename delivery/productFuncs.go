package delivery

import (
	"encoding/json"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	usecase "serv/usecase"

	"github.com/microcosm-cc/bluemonday"
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
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /products [get]
func (api *ProductHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	products, err := api.usecase.GetProducts()
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	for _, prod := range products {
		prod.Name = sanitizer.Sanitize(prod.Name)
		prod.Description = sanitizer.Sanitize(prod.Description)
		prod.Imgsrc = sanitizer.Sanitize(prod.Imgsrc)
	}

	json.NewEncoder(w).Encode(&model.Response{Body: products})
}
