package usecase

import (
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	rep "serv/repository"
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
	err := api.store.AddUser(newUser)
	if err != nil {
		return err
	}
	return nil
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

func (api *UserUsecase) ChangeUser(oldEmail string, params model.UserProfile) error {
	username := params.Username
	email := params.Email
	phone := params.Phone
	avatar := params.Avatar
	newUser := &model.UserProfile{Email: email, Username: username, Phone: phone, Avatar: avatar}
	err := api.store.UpdateUser(oldEmail, newUser)
	if err != nil {
		return err
	}
	return nil
}
