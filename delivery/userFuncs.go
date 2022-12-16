package delivery

import (
	"encoding/base64"
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
	usecase usecase.UserUsecaseInterface
}

func NewUserHandler(uuc usecase.UserUsecaseInterface) *UserHandler {
	return &UserHandler{
		usecase: uuc,
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
	if user := r.Context().Value("userdata").(*model.UserProfile); user == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	userProfile := r.Context().Value("userdata").(*model.UserProfile)

	userProfile.Email = sanitizer.Sanitize(userProfile.Email)
	userProfile.Username = sanitizer.Sanitize(userProfile.Username)
	userProfile.Phone = sanitizer.Sanitize(userProfile.Phone)
	userProfile.Avatar = sanitizer.Sanitize(userProfile.Avatar)
	for _, addr := range userProfile.Address {
		addr.City = sanitizer.Sanitize(addr.City)
		addr.House = sanitizer.Sanitize(addr.House)
		addr.Street = sanitizer.Sanitize(addr.Street)
	}

	for _, paym := range userProfile.PaymentMethods {
		paym.PaymentType = sanitizer.Sanitize(paym.PaymentType)
		paym.Number = sanitizer.Sanitize(paym.Number)
	}
	userProfile.ID = 0
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

	err = api.usecase.ChangeUser(oldUserData, &req)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	session, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	if req.Email != "" {
		err = api.usecase.ChangeEmail(session.Value, req.Email)
		if err != nil {
			log.Println("error with auth microservice: ", err)
			ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}
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
	if oldUserData := r.Context().Value("userdata").(*model.UserProfile); oldUserData == nil {
		log.Println("err get user from context ")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	oldUserData := r.Context().Value("userdata").(*model.UserProfile)

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("error parse file")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	defer file.Close()
	err = api.usecase.SetAvatar(oldUserData.ID, file)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	fileName := "/img/avatar" + strconv.FormatUint(uint64(oldUserData.ID), 10) + ".jpg"
	newUserData := model.UserProfile{Avatar: fileName}

	err = api.usecase.ChangeUser(oldUserData, &newUserData)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// ChangePassword godoc
// @Summary changes user password
// @Description changes user parameters
// @ID changeUserPassword
// @Accept  json
// @Produce  json
// @Tags User
// @Param userpassword body model.ChangePassword true "ChangePassword params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /user/password [post]
func (api *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.ChangePassword
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

	if len(req.NewPassword) < 6 {
		log.Println("validation error ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	user, err := api.usecase.GetUserByUsername(oldUserData.Email)
	if err != nil {
		log.Println("get GetUserByUsername ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if user.Email == "" {
		log.Println("get Email ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	byteUserPass, err := base64.RawStdEncoding.DecodeString(user.Password)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	if !checkPass(byteUserPass, req.OldPassword) {
		log.Println("Old Password is incorrect ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	salt := []byte("Base2022")
	hashedPass := hashPass(salt, req.NewPassword)
	b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)

	err = api.usecase.ChangeUserPassword(oldUserData.ID, b64Pass)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
