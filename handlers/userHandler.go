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
	email := params.Email
	//users, err := api.store.GetUsers()
	// if err != nil {
	// 	return 0, baseErrors.ErrServerError500
	// }
	//id := users[len(users)-1].ID + 1
	newUser := &model.UserDB{ID: 0, Email: email, Username: username, Password: password}
	id, err := api.store.AddUser(newUser)
	if err != nil {
		return 0, baseErrors.ErrServerError500
	}
	return id, nil
}

func (api *UserHandler) GetUserByUsernameAndPassword(email string, password string) (model.UserDB, error) {
	user, err := api.store.GetUserByUsernameAndPasswordFromDB(email, password)
	if err != nil {
		return model.UserDB{ID: 0, Email: "", Username: "", Password: ""}, baseErrors.ErrServerError500
	}

	if user == nil {
		return model.UserDB{ID: 0, Email: "", Username: "", Password: ""}, baseErrors.ErrNotFound404
	}
	return *user, nil
}
