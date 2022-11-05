package usecase

import (
	"database/sql"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	rep "serv/repository"
)

type UserUsecase struct {
	sessions map[string]string
	store    rep.UserStore
}

func NewUserUsecase(db *sql.DB) *UserUsecase {
	return &UserUsecase{
		sessions: make(map[string]string),
		store:    *rep.NewUserStore(db),
	}
}

func SetSession(uh *UserUsecase, key string, value string) {
	uh.sessions[key] = value
}
func GetSession(uh *UserUsecase, key string) (string, error) {
	if res, ok := uh.sessions[key]; ok {
		return res, nil
	}
	return "", baseErrors.ErrUnauthorized401
}

func DeleteSession(uh *UserUsecase, value string) {
	delete(uh.sessions, value)
}

func (api *UserUsecase) AddUser(params *model.UserCreateParams) (uint, error) {
	username := params.Username
	password := params.Password
	email := params.Email
	newUser := &model.UserDB{ID: 0, Email: email, Username: username, Password: password}
	id, err := api.store.AddUser(newUser)
	if err != nil {
		return 0, err
	}
	return id, nil
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

func (api *UserUsecase) ChangeUser(oldEmail string, params model.UserProfile) (int64, error) {
	username := params.Username
	email := params.Email
	phone := params.Phone
	avatar := params.Avatar
	newUser := &model.UserProfile{Email: email, Username: username, Phone: phone, Avatar: avatar}
	count, err := api.store.UpdateUser(oldEmail, newUser)
	if err != nil {
		return 0, err
	}
	return count, nil
}
