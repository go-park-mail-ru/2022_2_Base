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
	users, err := api.store.GetUsers()
	id := users[len(users)-1].ID + 1
	newUser := &model.UserDB{ID: id, Username: username, Password: password}
	id, err = api.store.AddUser(newUser)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	return id, nil
}

func (api *UserHandler) GetUserByUsername(username string) (model.UserDB, error) {

	users, err := api.store.GetUsers()
	if err != nil {
		return model.UserDB{ID: 0, Username: "", Password: ""}, baseErrors.ErrServerError500
	}

	var user *model.UserDB
	for _, u := range users {
		if u.Username == username {
			user = u
			break
		}
	}
	if user == nil {
		return model.UserDB{ID: 0, Username: "", Password: ""}, baseErrors.ErrNotFound404
	}
	return *user, nil
}
