package handlers

import (
	baseErrors "serv/errors"
	"serv/model"
)

type UserHandler struct {
	sessions map[string]uint
	store    UserStore
}

func (api *UserHandler) AddUser(params *model.UserCreateParams) (uint, error) {
	username := params.Username
	password := params.Password
	var id uint = 7
	newUser := &model.User{ID: id, Username: username, Password: password}
	id, err := api.store.AddUser(newUser)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	return id, nil
}

func (api *UserHandler) GetUserByUsername(username string) (model.User, error) {

	users, err := api.store.GetUsers()
	if err != nil {
		return model.User{ID: 0, Username: "", Password: ""}, baseErrors.ErrServerError500
	}

	var user *model.User
	for _, u := range users {
		if u.Username == username {
			user = u
			break
		}
	}
	if user == nil {
		return model.User{ID: 0, Username: "", Password: ""}, baseErrors.ErrNotFound404
	}

	return *user, nil
}
