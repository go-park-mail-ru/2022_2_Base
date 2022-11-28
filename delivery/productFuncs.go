package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	usecase "serv/usecase"
	"strconv"
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
// @Param   lastitemid    query     string  true  "lastitemid"
// @Param   count         query     string  true  "count"
// @Param   sort         query     string  false  "sort"
// @Success 200 {object} model.Product
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /products [get]
func (api *ProductHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	lastitemidS := r.URL.Query().Get("lastitemid")
	countS := r.URL.Query().Get("count")
	sort := r.URL.Query().Get("sort")
	lastitemid, err := strconv.Atoi(lastitemidS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	count, err := strconv.Atoi(countS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	products, err := api.usecase.GetProducts(lastitemid, count, sort)
	if err != nil {
		log.Println("error: ", err)
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
// @Param   lastitemid    query     string  true  "lastitemid"
// @Param   count         query     string  true  "count"
// @Param   sort         query     string  false  "sort"
// @Success 200 {object} model.Product
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /category/{category} [get]
func (api *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	s := strings.Split(r.URL.Path, "/")
	category := s[len(s)-1]
	sanitizer := bluemonday.UGCPolicy()
	lastitemidS := r.URL.Query().Get("lastitemid")
	countS := r.URL.Query().Get("count")
	sort := r.URL.Query().Get("sort")
	lastitemid, err := strconv.Atoi(lastitemidS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	count, err := strconv.Atoi(countS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	products, err := api.usecase.GetProductsWithCategory(category, lastitemid, count, sort)
	if err != nil {
		log.Println("error: ", err)
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

// GetProductByID godoc
// @Summary Gets product by id
// @Description Gets product by id
// @ID getProductByID
// @Accept  json
// @Produce  json
// @Tags Products
// @Param id path string true "Id of product"
// @Success 200 {object} model.Product
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /products/{id} [get]
func (api *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	s := strings.Split(r.URL.Path, "/")
	idS := s[len(s)-1]
	id, err := strconv.Atoi(idS)
	product, err := api.usecase.GetProductByID(id)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if product.Imgsrc != nil {
		*product.Imgsrc = sanitizer.Sanitize(*product.Imgsrc)
	}
	product.Name = sanitizer.Sanitize(product.Name)
	product.Category = sanitizer.Sanitize(product.Category)

	json.NewEncoder(w).Encode(&model.Response{Body: product})
}

// GetProductBySearch godoc
// @Summary Gets product by search
// @Description Gets product by search
// @ID getProductBySearch
// @Accept  json
// @Produce  json
// @Tags Products
// @Param search body model.Search true "search string"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /search [post]
func (api *ProductHandler) GetProductsBySearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	decoder := json.NewDecoder(r.Body)
	var req model.Search
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	products, err := api.usecase.GetProductsBySearch(req.Search)
	if err != nil {
		log.Println("error: ", err)
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

// GetSuggestions godoc
// @Summary Gets suggestions
// @Description Gets uggestions
// @ID getSuggestions
// @Accept  json
// @Produce  json
// @Tags Products
// @Param search body model.Search true "search string"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /suggestions [post]
func (api *ProductHandler) GetSuggestions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	decoder := json.NewDecoder(r.Body)
	var req model.Search
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	suggestions, err := api.usecase.GetSuggestions(req.Search)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	for _, sugg := range suggestions {
		sugg = sanitizer.Sanitize(sugg)
	}
	json.NewEncoder(w).Encode(&model.Response{Body: suggestions})
}
