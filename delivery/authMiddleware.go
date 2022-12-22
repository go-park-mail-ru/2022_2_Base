package delivery

import (
	"context"
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	usecase "serv/usecase"
)

type AuthenticationMiddlewareInterface interface {
	CheckAuthMiddleware(next http.Handler) http.Handler
}

type AuthMiddleware struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewAuthMiddleware(uc usecase.UserUsecaseInterface) AuthenticationMiddlewareInterface {
	return &AuthMiddleware{
		userUsecase: uc,
	}
}

type KeyUserdata struct {
	key string
}

func WithUser(ctx context.Context, user *model.UserProfile) context.Context {
	return context.WithValue(ctx, KeyUserdata{"userdata"}, user)
}

func (amw *AuthMiddleware) CheckAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			log.Println("no session")
			ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
			return
		}
		usName, err := amw.userUsecase.CheckSession(session.Value)
		if err != nil {
			log.Println("no session2")
			ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
			return
		}

		// hashTok := HashToken{Secret: []byte("Base")}
		// token := r.Header.Get("csrf")
		// curSession := model.Session{ID: 0, UserUUID: session.Value}
		// flag, err := hashTok.CheckCSRFToken(&curSession, token)
		// if err != nil || !flag {
		// 	log.Println("no csrf token")
		// 	ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		// 	return
		// }

		user, err := amw.userUsecase.GetUserByUsername(usName)
		if err != nil {
			log.Println("err get user ", err)
			ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}
		addresses, err := amw.userUsecase.GetAddressesByUserID(user.ID)
		if err != nil {
			log.Println("err get adresses ", err)
			ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}
		payments, err := amw.userUsecase.GetPaymentMethodByUserID(user.ID)
		if err != nil {
			log.Println("err get payments ", err)
			ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}

		if user.Email == "" {
			log.Println("err get Email ", err)
			ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
			return
		}

		userData := model.UserProfile{ID: user.ID, Email: user.Email, Username: user.Username, Address: addresses, PaymentMethods: payments}
		if user.Phone != nil {
			userData.Phone = *user.Phone
		}
		if user.Avatar != nil {
			userData.Avatar = *user.Avatar
		}
		next.ServeHTTP(w, r.WithContext(WithUser(r.Context(), &userData)))
	})
}
