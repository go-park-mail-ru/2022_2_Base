package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"strconv"

	usecase "serv/usecase"

	"github.com/microcosm-cc/bluemonday"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uuc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: *uuc,
	}
}

// GetUser godoc
// @Summary Get current user
// @Description gets user by username in cookies
// @ID getUser
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} model.UserProfile
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/profile [get]
func (api *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	if user := r.Context().Value("userdata").(*model.UserDB); user == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	user := r.Context().Value("userdata").(*model.UserDB)

	addresses, err := api.usecase.GetAddressesByUserID(user.ID)
	if err != nil {
		log.Println("err get adresses ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	payments, err := api.usecase.GetPaymentMethodByUserID(user.ID)
	if err != nil {
		log.Println("err get payments ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	userProfile := model.UserProfile{Email: user.Email, Username: user.Username}
	if user.Phone != nil {
		userProfile.Phone = *user.Phone
	} else {
		userProfile.Phone = ""
	}
	if user.Avatar != nil {
		userProfile.Avatar = *user.Avatar
	} else {
		userProfile.Avatar = ""
	}

	userProfile.Email = sanitizer.Sanitize(userProfile.Email)
	userProfile.Username = sanitizer.Sanitize(userProfile.Username)
	userProfile.Phone = sanitizer.Sanitize(userProfile.Phone)
	userProfile.Avatar = sanitizer.Sanitize(userProfile.Avatar)
	for _, addr := range addresses {
		addr.City = sanitizer.Sanitize(addr.City)
		addr.House = sanitizer.Sanitize(addr.House)
		addr.Street = sanitizer.Sanitize(addr.Street)
	}

	for _, paym := range payments {
		paym.PaymentType = sanitizer.Sanitize(paym.PaymentType)
		paym.Number = sanitizer.Sanitize(paym.Number)
		//paym.ExpiryDate = sanitizer.Sanitize(paym.ExpiryDate)

	}

	userProfile.Address = addresses
	userProfile.PaymentMethods = payments
	json.NewEncoder(w).Encode(userProfile)
}

// ChangeUser godoc
// @Summary changes user parameters
// @Description changes user parameters
// @ID changeUserParameters
// @Accept  json
// @Produce  json
// @Tags User
// @Param user body model.UserProfile true "UserProfile params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/profile [post]
func (api *UserHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.UserProfile
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	if oldUserData := r.Context().Value("userdata").(*model.UserDB); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value("userdata").(*model.UserDB)
	//log.Println("zzz")
	err = api.usecase.ChangeUser(oldUserData, &req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// SetAvatar godoc
// @Summary Set user's avatar
// @Description sets user's avatar
// @ID setAvatar
// @Accept  multipart/form-data
// @Produce  json
// @Tags User
// @Param file formData file true "user's avatar"
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/avatar [post]
func (api *UserHandler) SetAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return

	}
	if oldUserData := r.Context().Value("userdata").(*model.UserDB); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value("userdata").(*model.UserDB)

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("error parse file")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	defer file.Close()
	err = api.usecase.SetAvatar(oldUserData.ID, file)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	fileName := "/img/avatar" + strconv.FormatUint(uint64(oldUserData.ID), 10) + ".jpg"
	newUserData := model.UserProfile{Avatar: fileName}
	//newUserData.Avatar = fileName
	// if userDB.Phone != nil {
	// 	newUserData.Phone = *userDB.Phone
	// } else {
	// 	newUserData.Phone = ""
	// }

	err = api.usecase.ChangeUser(oldUserData, &newUserData)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
