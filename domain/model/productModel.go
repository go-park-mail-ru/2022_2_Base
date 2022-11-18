package model

import "time"

type Product struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	Price         float64  `json:"price"`
	DiscountPrice *float64 `json:"lowprice,omitempty"`
	Rating        *float64 `json:"rating,omitempty"`
	Imgsrc        *string  `json:"imgsrc,omitempty"`
}

type ProductCart struct {
	Items []int `json:"items,omitempty"`
}

type ProductCartItem struct {
	ItemID int `json:"itemid"`
}

type OrderItem struct {
	Count int      `json:"count"`
	Item  *Product `json:"item"`
}

type Order struct {
	ID                int          `json:"id"`
	UserID            int          `json:"userid"`
	Items             []*OrderItem `json:"items"`
	OrderStatus       string       `json:"orderstatus"`
	PaymentStatus     string       `json:"paymentstatus"`
	Address           *string      `json:"address"`
	Paymentcardnumber *string      `json:"card"`
	CreationDate      *time.Time   `json:"creationDate"`
	DeliveryDate      *time.Time   `json:"deliveryDate"`
}

type CartProduct struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Count         int      `json:"count"`
	Price         float64  `json:"price"`
	DiscountPrice *float64 `json:"lowprice,omitempty"`
	Imgsrc        *string  `json:"imgsrc,omitempty"`
}

type Cart struct {
	ID     int            `json:"id"`
	UserID int            `json:"userid"`
	Items  []*CartProduct `json:"items"`
}

type MakeOrder struct {
	UserID            int       `json:"userid"`
	Items             []int     `json:"items"`
	Address           string    `json:"adress"`
	Paymentcardnumber string    `json:"card,omitempty"`
	DeliveryDate      time.Time `json:"deliveryDate"`
}
