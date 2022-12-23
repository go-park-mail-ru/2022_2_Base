package delivery

import (
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	usecase "serv/usecase"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
)

type ProductHandler struct {
	usecase     usecase.ProductUsecaseInterface
	userUsecase usecase.UserUsecaseInterface
}

func NewProductHandler(puc usecase.ProductUsecaseInterface, uuc usecase.UserUsecaseInterface) *ProductHandler {
	return &ProductHandler{
		usecase:     puc,
		userUsecase: uuc,
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
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
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

	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}

	products, err := api.usecase.GetProducts(lastitemid, count, sort, userID)
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
		if prod.NominalPrice == prod.Price {
			prod.Price = 0
		}
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: products}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
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

	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}

	products, err := api.usecase.GetProductsWithCategory(category, lastitemid, count, sort, userID)
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
		if prod.NominalPrice == prod.Price {
			prod.Price = 0
		}
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: products}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// GetProductsWithBiggestDiscount godoc
// @Summary Gets products with biggest discount for main page
// @Description Gets products with biggest discount for main page
// @ID getProductsWithBiggestDiscount
// @Accept  json
// @Produce  json
// @Tags Products
// @Param   lastitemid    query     string  true  "lastitemid"
// @Param   count         query     string  true  "count"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /productswithdiscount [get]
func (api *ProductHandler) GetProductsWithBiggestDiscount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	lastitemidS := r.URL.Query().Get("lastitemid")
	countS := r.URL.Query().Get("count")
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

	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}

	products, err := api.usecase.GetProductsWithBiggestDiscount(lastitemid, count, userID)
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
		if prod.NominalPrice == prod.Price {
			prod.Price = 0
		}
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: products}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
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
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}

	product, err := api.usecase.GetProductByID(id, userID)
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
	if product.NominalPrice == product.Price {
		product.Price = 0
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: product}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
	var req model.Search
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}

	products, err := api.usecase.GetProductsBySearch(req.Search, userID)
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
		if prod.NominalPrice == prod.Price {
			prod.Price = 0
		}
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: products}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
	var req model.Search
	err := easyjson.UnmarshalFromReader(r.Body, &req)
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
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: suggestions}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// GetRecommendations godoc
// @Summary Gets recommendations for product
// @Description  Gets recommendations for product by id
// @ID getRecommendations
// @Accept  json
// @Produce  json
// @Tags Products
// @Param id path string true "Id of product"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /recommendations/{id} [get]
func (api *ProductHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	s := strings.Split(r.URL.Path, "/")
	idS := s[len(s)-1]
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}

	products, err := api.usecase.GetRecommendationProducts(id, userID)
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
		if prod.NominalPrice == prod.Price {
			prod.Price = 0
		}
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: products}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// GetBestProductInCategory godoc
// @Summary Gets random from 10 best products in category
// @Description  Gets random from 10 best products in category
// @ID GetBestProductInCategory
// @Accept  json
// @Produce  json
// @Tags Products
// @Param category path string true "The category of products"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /bestproduct/{category} [get]
func (api *ProductHandler) GetBestProductInCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	s := strings.Split(r.URL.Path, "/")
	category := s[len(s)-1]
	var userID int = 0
	session, err := r.Cookie("session_id")
	if err == nil {
		usName, err := api.userUsecase.CheckSession(session.Value)
		if err == nil {
			user, err := api.userUsecase.GetUserByUsername(usName)
			if err == nil {
				userID = user.ID
			}
		}
	}
	product, err := api.usecase.GetBestProductInCategory(category, userID)
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
	if product.NominalPrice == product.Price {
		product.Price = 0
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: product}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}
