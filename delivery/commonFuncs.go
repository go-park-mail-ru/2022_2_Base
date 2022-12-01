package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

// @title Reozon API
// @version 1.0
// @description Reazon back server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 89.208.198.137:8080
// @BasePath  /api/v1

type OrderHandler struct {
	usHandler UserHandler
	prHandler ProductHandler
}

func NewOrderHandler(us *UserHandler, pr *ProductHandler) *OrderHandler {
	return &OrderHandler{
		usHandler: *us,
		prHandler: *pr,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
	return
}

// GetCart godoc
// @Summary gets user's cart
// @Description gets user's cart
// @ID GetCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Success 200 {object} model.Order
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart [get]
func (api *OrderHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value("userdata").(*model.UserProfile)

	cart, err := api.prHandler.usecase.GetCart(UserData.ID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	cart.OrderStatus = sanitizer.Sanitize(cart.OrderStatus)
	cart.PaymentStatus = sanitizer.Sanitize(cart.PaymentStatus)
	for _, prod := range cart.Items {
		if prod.Item.Imgsrc != nil {
			*prod.Item.Imgsrc = sanitizer.Sanitize(*prod.Item.Imgsrc)
		}
		prod.Item.Name = sanitizer.Sanitize(prod.Item.Name)
		prod.Item.Category = sanitizer.Sanitize(prod.Item.Category)
	}
	prodCart := model.Cart{ID: cart.ID, UserID: cart.UserID}
	for _, prod := range cart.Items {
		prodCart.Items = append(prodCart.Items, &model.CartProduct{ID: prod.Item.ID, Name: prod.Item.Name, Count: prod.Count, Price: prod.Item.Price, DiscountPrice: prod.Item.DiscountPrice, Imgsrc: prod.Item.Imgsrc})
	}
	json.NewEncoder(w).Encode(prodCart)
}

// UpdateCart godoc
// @Summary updates user's cart
// @Description updates user's cart
// @ID UpdateCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Param items body model.ProductCart true "ProductCart items"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart [post]
func (api *OrderHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.ProductCart
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value("userdata").(*model.UserProfile)

	err = api.prHandler.usecase.UpdateOrder(UserData.ID, &req.Items)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// AddItemToCart godoc
// @Summary Adds item to cart
// @Description Adds item to cart
// @ID AddItemToCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Param items body model.ProductCartItem true "ProductCart item"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/insertintocart [post]
func (api *OrderHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.ProductCartItem
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value("userdata").(*model.UserProfile)

	err = api.prHandler.usecase.AddToOrder(UserData.ID, req.ItemID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// DeleteItemFromCart godoc
// @Summary Deletes Item From cart
// @Description Deletes Item From cart
// @ID DeleteItemFromCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Param items body model.ProductCartItem true "ProductCart item"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "Not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/deletefromcart [post]
func (api *OrderHandler) DeleteItemFromCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.ProductCartItem
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value("userdata").(*model.UserProfile)

	err = api.prHandler.usecase.DeleteFromOrder(UserData.ID, req.ItemID)
	if err == baseErrors.ErrNotFound404 {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrNotFound404, 404)
		return
	}
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// MakeOrder godoc
// @Summary makes user's order
// @Description makes user's order
// @ID MakeOrder
// @Accept  json
// @Produce  json
// @Tags Order
// @Param order body model.MakeOrder true "MakeOrder params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/makeorder [post]
func (api *OrderHandler) MakeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.MakeOrder
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value("userdata").(*model.UserProfile)

	if oldUserData.ID != req.UserID {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	err = api.prHandler.usecase.MakeOrder(&req)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// GetOrders godoc
// @Summary gets user's orders
// @Description gets user's orders
// @ID GetOrder
// @Accept  json
// @Produce  json
// @Tags Order
// @Success 200 {object} model.Order
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/orders [get]
func (api *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value("userdata").(*model.UserProfile)

	orders, err := api.prHandler.usecase.GetOrders(UserData.ID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	var responseOrders []*model.OrderModelGetOrders
	for _, order := range orders.Orders {
		order.OrderStatus = sanitizer.Sanitize(order.OrderStatus)
		order.PaymentStatus = sanitizer.Sanitize(order.PaymentStatus)

		newOrder := model.OrderModelGetOrders{ID: int(order.ID), UserID: int(order.UserID), OrderStatus: order.OrderStatus, PaymentStatus: order.PaymentStatus}
		for _, prod := range order.Items {
			if prod.Imgsrc != nil {
				*prod.Imgsrc = sanitizer.Sanitize(*prod.Imgsrc)
			}
			prod.Name = sanitizer.Sanitize(prod.Name)
			newOrder.Items = append(newOrder.Items, &model.CartProduct{ID: int(prod.ID), Name: prod.Name, Count: int(prod.Count), Price: prod.Price, DiscountPrice: prod.DiscountPrice, Imgsrc: prod.Imgsrc})
		}
		t1 := time.Unix(order.CreationDate, 0)
		newOrder.CreationDate = &t1
		t2 := time.Unix(order.DeliveryDate, 0)
		newOrder.DeliveryDate = &t2

		newOrder.Address = model.Address{ID: int(order.Address.ID), City: order.Address.City, Street: order.Address.Street, House: order.Address.House, Flat: order.Address.Flat, Priority: order.Address.Priority}
		newOrder.Address.City = sanitizer.Sanitize(newOrder.Address.City)
		newOrder.Address.Street = sanitizer.Sanitize(newOrder.Address.Street)
		newOrder.Address.House = sanitizer.Sanitize(newOrder.Address.House)
		newOrder.Address.Flat = sanitizer.Sanitize(newOrder.Address.Flat)

		newOrder.Paymentcard = model.PaymentMethod{ID: int(order.PaymentMethod.ID), PaymentType: order.PaymentMethod.PaymentType, Number: order.PaymentMethod.Number, Priority: order.PaymentMethod.Priority}
		t3 := time.Unix(order.PaymentMethod.ExpiryDate, 0)
		newOrder.Paymentcard.ExpiryDate = t3
		newOrder.Paymentcard.PaymentType = sanitizer.Sanitize(newOrder.Paymentcard.PaymentType)
		newOrder.Paymentcard.Number = sanitizer.Sanitize(newOrder.Paymentcard.Number)

		responseOrders = append(responseOrders, &newOrder)
	}
	json.NewEncoder(w).Encode(&model.Response{Body: responseOrders})
}

// GetComments godoc
// @Summary gets product's comments
// @Description gets product's comments
// @ID GetComments
// @Accept  json
// @Produce  json
// @Tags Comments
// @Param id path string true "Id of product"
// @Success 200 {object} model.Comment
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /products/comments/{id} [get]
func (api *OrderHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	s := strings.Split(r.URL.Path, "/")
	idS := s[len(s)-1]
	id, err := strconv.Atoi(idS)
	commentsDB, err := api.prHandler.usecase.GetComments(id)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	comments, err := api.usHandler.usecase.SetUsernamesForComments(commentsDB)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	for _, comm := range comments {
		comm.Username = sanitizer.Sanitize(comm.Username)
		comm.Pros = sanitizer.Sanitize(comm.Pros)
		comm.Cons = sanitizer.Sanitize(comm.Cons)
		comm.Comment = sanitizer.Sanitize(comm.Comment)
	}
	json.NewEncoder(w).Encode(&model.Response{Body: comments})
}

// CreateComment godoc
// @Summary creates product's comment by user
// @Description creates product's comment by user
// @ID CreateComment
// @Accept  json
// @Produce  json
// @Tags Comments
// @Param comment body model.CreateComment true "Comment params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/makecomment [post]
func (api *OrderHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.CreateComment
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value("userdata").(*model.UserProfile)
	if oldUserData.ID != req.UserID {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	err = api.prHandler.usecase.CreateComment(&req)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}
