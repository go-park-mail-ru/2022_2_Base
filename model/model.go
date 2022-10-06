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

type Product struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"lowprice"`
	Rating        float64 `json:"rating"`
	Imgsrc        string  `json:"imgsrc"`
}

type Error struct {
	Message string
}

func UserToString(user UserDB) string {
	res, _ := json.Marshal(user)
	return string(res)
}
