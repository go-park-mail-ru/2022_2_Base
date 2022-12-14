package model

import "time"

type UserDB struct {
	ID       int     `json:"id"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Phone    *string `json:"phone,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

type UserCreateParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Address struct {
	ID       int    `json:"id"`
	City     string `json:"city"`
	Street   string `json:"street"`
	House    string `json:"house"`
	Flat     string `json:"flat"`
	Priority bool   `json:"priority"`
}

type PaymentMethod struct {
	ID          int       `json:"id"`
	PaymentType string    `json:"type"`
	Number      string    `json:"number"`
	ExpiryDate  time.Time `json:"expirydate"`
	Priority    bool      `json:"priority"`
}

type UserProfile struct {
	ID             int              `json:"id,omitempty"`
	Email          string           `json:"email"`
	Username       string           `json:"username"`
	Phone          string           `json:"phone,omitempty"`
	Avatar         string           `json:"avatar,omitempty"`
	Address        []*Address       `json:"address,omitempty"`
	PaymentMethods []*PaymentMethod `json:"paymentmethods,omitempty"`
}

type Session struct {
	ID       int    `json:"id"`
	UserUUID string `json:"useruuid"`
}

type ChangePassword struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}
