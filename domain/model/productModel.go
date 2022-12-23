package model

import "time"

type Property struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type Product struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	Category      string      `json:"category"`
	Price         float64     `json:"lowprice,omitempty"`
	NominalPrice  float64     `json:"price"`
	Rating        float64     `json:"rating"`
	Imgsrc        *string     `json:"imgsrc,omitempty"`
	CommentsCount *int        `json:"commentscount,omitempty"`
	Properties    []*Property `json:"properties"`
	IsFavorite    bool        `json:"isfavorite"`
}

type ProductCart struct {
	Items []int `json:"items,omitempty"`
}

type ProductCartItem struct {
	ItemID int `json:"itemid"`
}

type OrderItem struct {
	Count      int      `json:"count"`
	Item       *Product `json:"item"`
	IsFavorite bool     `json:"isfavorite"`
}

type Order struct {
	ID            int          `json:"id"`
	UserID        int          `json:"userid"`
	Items         []*OrderItem `json:"items"`
	OrderStatus   string       `json:"orderstatus"`
	PaymentStatus string       `json:"paymentstatus"`
	AddressID     int          `json:"address"`
	PaymentcardID int          `json:"card"`
	CreationDate  *time.Time   `json:"creationdate"`
	DeliveryDate  *time.Time   `json:"deliverydate"`
	Promocode     *string      `json:"promo,omitempty"`
}

type CartProduct struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Count        int     `json:"count"`
	Price        float64 `json:"lowprice,omitempty"`
	NominalPrice float64 `json:"price"`
	Imgsrc       *string `json:"imgsrc,omitempty"`
	IsFavorite   bool    `json:"isfavorite"`
}

type Cart struct {
	ID        int            `json:"id"`
	UserID    int            `json:"userid"`
	Items     []*CartProduct `json:"items"`
	Promocode string         `json:"promocode,omitempty"`
}

type MakeOrder struct {
	UserID        int       `json:"userid"`
	Items         []int     `json:"items"`
	AddressID     int       `json:"address"`
	PaymentcardID int       `json:"card"`
	DeliveryDate  time.Time `json:"deliverydate"`
}

type ChangeOrderStatus struct {
	UserID      int    `json:"userid"`
	OrderID     int    `json:"orderid"`
	OrderStatus string `json:"orderstatus"`
}

type OrderModelGetOrders struct {
	ID            int            `json:"id"`
	UserID        int            `json:"userid"`
	Items         []*CartProduct `json:"items"`
	OrderStatus   string         `json:"orderstatus"`
	PaymentStatus string         `json:"paymentstatus"`
	Address       Address        `json:"address"`
	Paymentcard   PaymentMethod  `json:"card"`
	CreationDate  *time.Time     `json:"creationdate"`
	DeliveryDate  *time.Time     `json:"deliverydate"`
	Promocode     string         `json:"promocode,omitempty"`
}

type Search struct {
	Search string `json:"search"`
}
