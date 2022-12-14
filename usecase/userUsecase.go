package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	auth "serv/microservices/auth/gen_files"
	mail "serv/microservices/mail/gen_files"
	rep "serv/repository"
	"strconv"
)

type UserUsecaseInterface interface {
	SetSession(userEmail string) (*auth.SessionID, error)
	CheckSession(sessID string) (string, error)
	ChangeEmail(sessID string, newEmail string) error
	DeleteSession(sessID string) error
	SendMail(in model.Mail) error
	GenPromocode(userID int) string
	AddUser(params *model.UserCreateParams) (int, error)
	GetUserByUsername(email string) (model.UserDB, error)
	GetAddressesByUserID(userID int) ([]*model.Address, error)
	GetPaymentMethodByUserID(userID int) ([]*model.PaymentMethod, error)
	ChangeUser(oldUserData *model.UserProfile, params *model.UserProfile) error
	ChangeUserAddresses(userID int, userAddresses []*model.Address, queryAddresses []*model.Address) error
	ChangeUserPayments(userID int, userPayments []*model.PaymentMethod, queryPayments []*model.PaymentMethod) error
	ChangeUserPassword(userID int, newPass string) error
	SetAvatar(usedID int, file multipart.File) error
	SetUsernamesForComments(comms []*model.CommentDB) ([]*model.Comment, error)
}

type UserUsecase struct {
	sessManager auth.AuthCheckerClient
	mailManager mail.MailServiceClient
	store       rep.UserStoreInterface
}

func NewUserUsecase(us rep.UserStoreInterface, sessManager auth.AuthCheckerClient, mailManager mail.MailServiceClient) UserUsecaseInterface {
	return &UserUsecase{
		sessManager: sessManager,
		mailManager: mailManager,
		store:       us,
	}
}

func (uh *UserUsecase) SetSession(userEmail string) (*auth.SessionID, error) {
	sess, err := uh.sessManager.Create(
		context.Background(),
		&auth.Session{
			Login: userEmail,
		})
	return sess, err
}
func (uh *UserUsecase) CheckSession(sessID string) (string, error) {
	sess, err := uh.sessManager.Check(
		context.Background(),
		&auth.SessionID{
			ID: sessID,
		})
	if err != nil {
		return "", err
	}
	return sess.Login, nil
}

func (uh *UserUsecase) ChangeEmail(sessID string, newEmail string) error {
	ans, err := uh.sessManager.ChangeEmail(
		context.Background(),
		&auth.NewLogin{
			ID:    sessID,
			Login: newEmail,
		})
	if err != nil || !ans.IsSuccessful {
		return err
	}
	return nil
}

func (uh *UserUsecase) DeleteSession(sessID string) error {
	ans, err := uh.sessManager.Delete(
		context.Background(),
		&auth.SessionID{
			ID: sessID,
		})
	if err != nil || !ans.IsSuccessful {
		return err
	}
	return nil
}

func (api *UserUsecase) SendMail(in model.Mail) error {
	orderid := int32(in.OrderID)
	ans, err := api.mailManager.SendMail(
		context.Background(),
		&mail.Mail{Type: in.Type, Username: in.Username, Useremail: in.Useremail, OrderStatus: &in.OrderStatus, Promocode: &in.Promocode, OrderID: &orderid})
	if err != nil || !ans.IsSuccessful {
		return err
	}
	return nil
}

func (api *UserUsecase) GenPromocode(userID int) string {
	h := hmac.New(sha256.New, []byte("Base2022"))
	data := fmt.Sprintf("%d", userID)
	h.Write([]byte(data))
	hashedStr := hex.EncodeToString(h.Sum(nil))
	promocode := "A10" + hashedStr[:5]
	return promocode
}

func (api *UserUsecase) AddUser(params *model.UserCreateParams) (int, error) {
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

func (api *UserUsecase) ChangeUser(oldUserData *model.UserProfile, params *model.UserProfile) error {
	newUser := &model.UserProfile{Email: oldUserData.Email, Username: oldUserData.Username, Phone: oldUserData.Phone, Avatar: oldUserData.Avatar, Address: oldUserData.Address, PaymentMethods: oldUserData.PaymentMethods}
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
	if params.Address != nil {
		err := api.ChangeUserAddresses(oldUserData.ID, oldUserData.Address, params.Address)
		if err != nil {
			return err
		}
	}
	if params.PaymentMethods != nil {
		err := api.ChangeUserPayments(oldUserData.ID, oldUserData.PaymentMethods, params.PaymentMethods)
		if err != nil {
			return err
		}
	}

	err := api.store.UpdateUser(oldUserData.ID, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (api *UserUsecase) ChangeUserAddresses(userID int, userAddresses []*model.Address, queryAddresses []*model.Address) error {
	for _, addr := range queryAddresses {
		if addr == nil {
			continue
		}
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
			if addr == nil {
				continue
			}
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
		if paym == nil {
			continue
		}
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
		for _, paym := range queryPayments {
			if paym == nil {
				continue
			}
			if paym.ID == paymFromDB.ID {
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

func (api *UserUsecase) ChangeUserPassword(userID int, newPass string) error {
	return api.store.ChangeUserPasswordDB(userID, newPass)
}

func (api *UserUsecase) SetAvatar(usedID int, file multipart.File) error {
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

func (api *UserUsecase) SetUsernamesForComments(comms []*model.CommentDB) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	for _, comm := range comms {
		usName, avatar, err := api.store.GetUsernameAndAvatarByIDFromDB(comm.UserID)
		if err != nil {
			return nil, err
		}
		comment := &model.Comment{Username: usName, UserID: comm.UserID, Pros: comm.Pros, Cons: comm.Cons, Comment: comm.Comment, Rating: comm.Rating, UserAvatar: avatar}

		comments = append(comments, comment)
	}
	return comments, nil
}
