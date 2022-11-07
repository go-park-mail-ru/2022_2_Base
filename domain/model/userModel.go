package model

type UserDB struct {
	ID                  int      `json:"id"`
	Email               string   `json:"email"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	Phone               *string  `json:"phone,omitempty"`
	Avatar              *string  `json:"avatar,omitempty"`
	PaymentMethodsUUIDs []string `json:"paymentmethods,omitempty"`
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

type UserProfile struct {
	Email               string   `json:"email"`
	Username            string   `json:"username"`
	Phone               string   `json:"phone,omitempty"`
	Avatar              string   `json:"avatar,omitempty"`
	PaymentMethodsUUIDs []string `json:"paymentmethods,omitempty"`
}
