package model

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ID            uint `json:"id"`
	name          string
	description   string
	price         uint
	discountPrice uint
	rating        uint
	imgsrc        string
}

type Error struct {
	Message string
}
