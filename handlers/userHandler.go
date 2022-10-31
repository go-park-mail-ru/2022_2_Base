package handlers

import (
	baseErrors "serv/errors"
	"serv/model"
)

type UserHandler struct {
	sessions map[string]string
	store    UserStore
}

func (api *UserHandler) AddUser(params *model.UserCreateParams) (uint, error) {
	username := params.Username
	password := params.Password
	email := params.Email
	newUser := &model.UserDB{ID: 0, Email: email, Username: username, Password: password}
	id, err := api.store.AddUser(newUser)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	return id, nil
}

func (api *UserHandler) GetUserByUsername(email string) (model.UserDB, error) {
	user, err := api.store.GetUserByUsernameFromDB(email)
	if err != nil {
		return model.UserDB{ID: 0, Email: "", Username: "", Password: ""}, baseErrors.ErrServerError500
	}

	if user == nil {
		return model.UserDB{ID: 0, Email: "", Username: "", Password: ""}, baseErrors.ErrNotFound404
	}
	return *user, nil
}

func (api *UserHandler) ChangeUser(oldEmail string, params model.UserProfile) (int64, error) {
	username := params.Username
	email := params.Email
	newUser := &model.UserProfile{Email: email, Username: username}
	count, err := api.store.UpdateUser(oldEmail, newUser)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	return count, nil
}
