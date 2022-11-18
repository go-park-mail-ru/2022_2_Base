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

func (api *UserUsecase) GetAddressesByUserID(userID int) ([]*model.Address, error) {
	return api.store.GetAddressesByUserIDFromDB(userID)
}

func (api *UserUsecase) GetPaymentMethodByUserID(userID int) ([]*model.PaymentMethod, error) {
	return api.store.GetPaymentMethodByUserIDFromDB(userID)
}

func (api *UserUsecase) ChangeUser(oldUserData *model.UserDB, params *model.UserProfile) error {
	addresses, err := api.GetAddressesByUserID(oldUserData.ID)
	if err != nil {
		return err
	}
	payments, err := api.GetPaymentMethodByUserID(oldUserData.ID)
	if err != nil {
		return err
	}
	newUser := &model.UserProfile{Email: oldUserData.Email, Username: oldUserData.Username, Address: addresses, PaymentMethods: payments}
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
	if len(params.Address) > 0 {
		newUser.Address = params.Address
	}
	if len(params.PaymentMethods) > 0 {
		newUser.PaymentMethods = params.PaymentMethods
	}

	err = api.store.UpdateUser(oldUserData.ID, newUser)
	if err != nil {
		return err
	}
	err = api.ChangeUserAddresses(oldUserData.ID, addresses, params.Address)
	if err != nil {
		return err
	}
	err = api.ChangeUserPayments(oldUserData.ID, payments, params.PaymentMethods)
	if err != nil {
		return err
	}

	return nil
}

func (api *UserUsecase) ChangeUserAddresses(userID int, userAddresses []*model.Address, queryAddresses []*model.Address) error {
	for _, addr := range queryAddresses {
		flag := true
		for _, addrFromDB := range userAddresses {
			if addr.ID == addrFromDB.ID {
				err := api.store.UpdateUsersAddress(addr.ID, addr)
				if err != nil {
					return err
				}
				flag = false
				break
			}
		}
		if flag {
			err := api.store.AddUsersAddress(userID, addr)
			if err != nil {
				return err
			}
		}
	}
	for _, addrFromDB := range userAddresses {
		flag := true
		for _, addr := range queryAddresses {
			if addr.ID == addrFromDB.ID {
				flag = false
				break
			}
		}
		if flag {
			err := api.store.DeleteUsersAddress(addrFromDB.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (api *UserUsecase) ChangeUserPayments(userID int, userPayments []*model.PaymentMethod, queryPayments []*model.PaymentMethod) error {
	for _, paym := range queryPayments {
		flag := true
		for _, paymFromDB := range userPayments {
			if paym.ID == paymFromDB.ID {
				err := api.store.UpdateUsersPayment(paym.ID, paym)
				if err != nil {
					return err
				}
				flag = false
				break
			}
		}
		if flag {
			err := api.store.AddUsersPayment(userID, paym)
			if err != nil {
				return err
			}
		}
	}
	for _, paymFromDB := range userPayments {
		flag := true
		for _, addr := range queryPayments {
			if addr.ID == paymFromDB.ID {
				flag = false
				break
			}
		}
		if flag {
			err := api.store.DeleteUsersPayment(paymFromDB.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (api *UserUsecase) SetAvatar(usedID int, file multipart.File) error {
	//fileName := "./img/avatars/avatar" + strconv.FormatUint(uint64(usedID), 10) + ".jpg"
	fileName := "/avatars/avatar" + strconv.FormatUint(uint64(usedID), 10) + ".jpg"
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
