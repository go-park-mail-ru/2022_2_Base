package model

import "encoding/json"

type UserDB struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserProfile struct {
	Email               string   `json:"email"`
	Username            string   `json:"username"`
	Phone               string   `json:"phone,omitempty"`
	Avatar              string   `json:"avatar,omitempty"`
	PaymentMethodsUUIDs []string `json:"paymentmethods,omitempty"`
}

type Product struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description,omitempty"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"lowprice,omitempty"`
	Rating        float64 `json:"rating,omitempty"`
	Imgsrc        string  `json:"imgsrc,omitempty"`
}

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

func UserToString(user UserDB) string {
	res, _ := json.Marshal(user)
	return string(res)
}
