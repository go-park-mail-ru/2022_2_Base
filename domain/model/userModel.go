package model

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

type Adress struct {
	ID       int    `json:"id"`
	City     string `json:"city"`
	Street   string `json:"street"`
	House    string `json:"house"`
	Priority bool   `json:"primary"`
}

type PaymentMethod struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Number     string `json:"number"`
	ExpiryDate string `json:"expirydate"`
	Priority   bool   `json:"priority"`
}

type UserProfile struct {
	Email          string           `json:"email"`
	Username       string           `json:"username"`
	Phone          string           `json:"phone,omitempty"`
	Avatar         string           `json:"avatar,omitempty"`
	Adress         []*Adress        `json:"adress,omitempty"`
	PaymentMethods []*PaymentMethod `json:"paymentmethods,omitempty"`
}

type Session struct {
	ID       int    `json:"id"`
	UserUUID string `json:"useruuid"`
}
