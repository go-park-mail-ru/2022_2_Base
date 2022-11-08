package model

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description,omitempty"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"lowprice,omitempty"`
	Rating        float64 `json:"rating,omitempty"`
	Imgsrc        string  `json:"imgsrc,omitempty"`
}

type ProductCart struct {
	Items []int `json:"items,omitempty"`
}

type Order struct {
	ID     int        `json:"id"`
	UserID uint       `json:"userid"`
	Items  []*Product `json:"items"`
	//Items         []string `json:"items"`
	OrderStatus   string `json:"orderstatus"`
	PaymentStatus string `json:"paymentstatus"`
	Adress        string `json:"adress"`
}
