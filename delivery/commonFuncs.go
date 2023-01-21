package delivery

import (
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"strconv"
	"strings"
	"time"

	"github.com/mailru/easyjson"
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
	_, _, _ = easyjson.MarshalToHTTPResponseWriter(&model.Error{Error: err.Error()}, w)
}

// GetCart godoc
// @Summary gets user's cart
// @Description gets user's cart
// @ID GetCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Success 200 {object} model.Cart
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart [get]
func (api *OrderHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

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
		if prod.Item.NominalPrice == prod.Item.Price {
			prod.Item.Price = 0
		}
	}
	prodCart := model.Cart{ID: cart.ID, UserID: cart.UserID}
	if cart.Promocode != nil {
		prodCart.Promocode = *cart.Promocode
	}
	for _, prod := range cart.Items {
		prodCart.Items = append(prodCart.Items, &model.CartProduct{ID: prod.Item.ID, Name: prod.Item.Name, Count: prod.Count, Price: prod.Item.Price, NominalPrice: prod.Item.NominalPrice, Imgsrc: prod.Item.Imgsrc, IsFavorite: prod.IsFavorite})
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(prodCart, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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

	var req model.ProductCart
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	err = api.prHandler.usecase.UpdateOrder(UserData.ID, &req.Items)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// SetPromocode godoc
// @Summary Sets promocode for cart
// @Description Sets promocode for cart
// @ID SetPromocode
// @Accept  json
// @Produce  json
// @Tags Order
// @Param promo body model.Promocode true "Promocode"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 403 {object} model.Error "Forbidden"
// @Failure 409 {object} model.Error "Conflict - UserDB already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/setpromocode [post]
func (api *OrderHandler) SetPromocode(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	var req model.Promocode
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	err = api.prHandler.usecase.SetPromocode(UserData.ID, req.Promocode)
	if err == baseErrors.ErrConflict409 {
		log.Println("promocode is already used ")
		ReturnErrorJSON(w, baseErrors.ErrConflict409, 409)
		return
	}
	if err == baseErrors.ErrForbidden403 {
		log.Println("promocode is invalid ")
		ReturnErrorJSON(w, baseErrors.ErrForbidden403, 403)
		return
	}
	if err == baseErrors.ErrUnauthorized401 {
		log.Println("wrong promocode")
		ReturnErrorJSON(w, baseErrors.ErrForbidden403, 403)
		return
	}
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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

	var req model.ProductCartItem
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	err = api.prHandler.usecase.AddToOrder(UserData.ID, req.ItemID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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

	var req model.ProductCartItem
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

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

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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

	var req model.MakeOrder
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	if oldUserData.ID != req.UserID {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	orderID, err := api.prHandler.usecase.MakeOrder(&req)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	RegisterMail := model.Mail{Type: "orderstatus", Username: oldUserData.Username, Useremail: oldUserData.Email, OrderID: orderID, OrderStatus: "created"}

	err = api.usHandler.usecase.SendMail(RegisterMail)
	if err != nil {
		log.Println("error sending email ", err)
	}

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// ChangeOrderStatus godoc
// @Summary changess order's status
// @Description changess order's status
// @ID ChangeOrderStatus
// @Accept  json
// @Produce  json
// @Tags Order
// @Param order body model.ChangeOrderStatus true "SetOrderStatus params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/changeorderstatus [post]
func (api *OrderHandler) ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	var req model.ChangeOrderStatus
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	err = api.prHandler.usecase.ChangeOrderStatus(oldUserData.ID, &req)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	Mail := model.Mail{Type: "orderstatus", Username: oldUserData.Username, Useremail: oldUserData.Email, OrderID: req.OrderID, OrderStatus: req.OrderStatus}

	err = api.usHandler.usecase.SendMail(Mail)
	if err != nil {
		log.Println("error sending email ", err)
	}

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// GetOrders godoc
// @Summary gets user's orders
// @Description gets user's orders
// @ID GetOrder
// @Accept  json
// @Produce  json
// @Tags Order
// @Success 200 {object} model.OrderModelGetOrders
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart/orders [get]
func (api *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

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
		newOrder.Items = []*model.CartProduct{}
		for _, prod := range order.Items {
			if prod.Imgsrc != nil {
				*prod.Imgsrc = sanitizer.Sanitize(*prod.Imgsrc)
			}
			prod.Name = sanitizer.Sanitize(prod.Name)
			if prod.NominalPrice == prod.Price {
				prod.Price = 0
			}
			newOrder.Items = append(newOrder.Items, &model.CartProduct{ID: int(prod.ID), Name: prod.Name, Count: int(prod.Count), Price: prod.Price, NominalPrice: prod.NominalPrice, Imgsrc: prod.Imgsrc})
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

		if order.Promocode != nil {
			newOrder.Promocode = *order.Promocode
		}
		newOrder.Promocode = sanitizer.Sanitize(newOrder.Promocode)

		responseOrders = append(responseOrders, &newOrder)
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: responseOrders}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{Body: comments}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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

	var req model.CreateComment
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)
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
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// GetFavorites
// @Summary Gets user's favorites
// @Description Gets user's favorites
// @ID GetFavorites
// @Accept  json
// @Produce  json
// @Tags User
// @Param   lastitemid    query     string  true  "lastitemid"
// @Param   count         query     string  true  "count"
// @Param   sort         query     string  false  "sort"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/favorites [get]
func (api *ProductHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
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

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	products, err := api.usecase.GetFavorites(UserData.ID, lastitemid, count, sanitizer.Sanitize(sort))
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

// InsertItemIntoFav godoc
// @Summary Inserts Item into favorite
// @Description Inserts Item into favorite
// @ID InsertItemIntoFav
// @Accept  json
// @Produce  json
// @Tags User
// @Param item body model.ProductCartItem true "Favorite item"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/insertintofav [post]
func (api *ProductHandler) InsertItemIntoFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	var req model.ProductCartItem
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	err = api.usecase.InsertItemIntoFavorites(UserData.ID, req.ItemID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}

// DeleteItemFromFav godoc
// @Summary Deletes Item From favorite
// @Description Deletes Item From favorite
// @ID DeleteItemFromFav
// @Accept  json
// @Produce  json
// @Tags User
// @Param item body model.ProductCartItem true "Favorite item"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/deletefromfav [post]
func (api *ProductHandler) DeleteItemFromFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	var req model.ProductCartItem
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	if r.Context().Value(KeyUserdata{"userdata"}) == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	UserData := r.Context().Value(KeyUserdata{"userdata"}).(*model.UserProfile)

	err = api.usecase.DeleteItemFromFavorites(UserData.ID, req.ItemID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	_, _, err = easyjson.MarshalToHTTPResponseWriter(&model.Response{}, w)
	if err != nil {
		log.Println("serialize error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
}
