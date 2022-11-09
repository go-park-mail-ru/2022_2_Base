package delivery

import (
	"encoding/json"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	usecase "serv/usecase"
	"strings"

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
		if prod.Imgsrc != nil {
			*prod.Imgsrc = sanitizer.Sanitize(*prod.Imgsrc)
		}
		prod.Name = sanitizer.Sanitize(prod.Name)
		prod.Category = sanitizer.Sanitize(prod.Category)
	}
	json.NewEncoder(w).Encode(&model.Response{Body: products})
}

// GetProductsByCategory godoc
// @Summary Gets products by category
// @Description Gets products by category
// @ID GetProductsByCategory
// @Accept  json
// @Produce  json
// @Tags Products
// @Param category path string true "The category of products"
// @Success 200 {object} model.Product
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /products/{category} [get]
func (api *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	s := strings.Split(r.URL.Path, "/")
	category := s[len(s)-1]
	sanitizer := bluemonday.UGCPolicy()
	products, err := api.usecase.GetProductsWithCategory(category)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	for _, prod := range products {

		if prod.Imgsrc != nil {
			*prod.Imgsrc = sanitizer.Sanitize(*prod.Imgsrc)
		}
		prod.Name = sanitizer.Sanitize(prod.Name)
		prod.Category = sanitizer.Sanitize(prod.Category)
	}
	json.NewEncoder(w).Encode(&model.Response{Body: products})
}
