package usecase

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	rep "serv/repository"
	"strconv"
)

type UserUsecase struct {
	sessions map[string]string
	store    rep.UserStore
}

func NewUserUsecase(us *rep.UserStore) *UserUsecase {
	return &UserUsecase{
		sessions: make(map[string]string),
		store:    *us,
	}
}

func (uh *UserUsecase) SetSession(key string, value string) {
	uh.sessions[key] = value
}
func (uh *UserUsecase) GetSession(key string) (string, error) {
	if res, ok := uh.sessions[key]; ok {
		return res, nil
	}
	return "", baseErrors.ErrUnauthorized401
}

func (uh *UserUsecase) DeleteSession(value string) {
	delete(uh.sessions, value)
}

func (api *UserUsecase) AddUser(params *model.UserCreateParams) error {
	username := params.Username
	password := params.Password
	email := params.Email
	newUser := &model.UserDB{ID: 0, Email: email, Username: username, Password: password}
	return api.store.AddUser(newUser)
}

func (api *UserUsecase) GetUserByUsername(email string) (model.UserDB, error) {
	user, err := api.store.GetUserByUsernameFromDB(email)
	if err != nil {
		return model.UserDB{ID: 0, Email: "", Username: "", Password: ""}, err
	}

	if user == nil {
		return model.UserDB{ID: 0, Email: "", Username: "", Password: ""}, baseErrors.ErrNotFound404
	}
	return *user, nil
}

func (api *UserUsecase) GetAdressesByUserID(userID int) ([]*model.Adress, error) {
	return api.store.GetAdressesByUserIDFromDB(userID)
}

func (api *UserUsecase) GetPaymentMethodByUserID(userID int) ([]*model.PaymentMethod, error) {
	return api.store.GetPaymentMethodByUserIDFromDB(userID)
}

func (api *UserUsecase) ChangeUser(oldUserData *model.UserDB, params *model.UserProfile) error {
	adresses, err := api.GetAdressesByUserID(oldUserData.ID)
	if err != nil {
		return err
	}
	payments, err := api.GetPaymentMethodByUserID(oldUserData.ID)
	if err != nil {
		return err
	}
	newUser := &model.UserProfile{Email: oldUserData.Email, Username: oldUserData.Username, Adress: adresses, PaymentMethods: payments}
	if oldUserData.Avatar != nil {
		newUser.Avatar = *oldUserData.Avatar
	}
	if oldUserData.Phone != nil {
		newUser.Phone = *oldUserData.Phone
	}
	if params.Email != "" {
		newUser.Email = params.Email
	}
	if params.Username != "" {
		newUser.Username = params.Username
	}
	if params.Phone != "" {
		newUser.Phone = params.Phone
	}
	if params.Avatar != "" {
		newUser.Avatar = params.Avatar
	}
	if len(params.Adress) > 0 {
		newUser.Adress = params.Adress
	}
	if len(params.PaymentMethods) > 0 {
		newUser.PaymentMethods = params.PaymentMethods
	}

	err = api.store.UpdateUser(oldUserData.ID, newUser)
	if err != nil {
		return err
	}
	for _, addr := range params.Adress {
		err = api.store.UpdateUsersAdress(addr.ID, addr)
		if err != nil {
			return err
		}
	}
	for _, paym := range params.PaymentMethods {
		err = api.store.UpdateUsersPayment(paym.ID, paym)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *UserUsecase) SetAvatar(usedID int, file multipart.File) error {
	fileName := "/app/img/avatars/avatar" + strconv.FormatUint(uint64(usedID), 10) + ".jpg"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error create/open file")
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("error copy to new file")
		return err
	}
	return nil
}
