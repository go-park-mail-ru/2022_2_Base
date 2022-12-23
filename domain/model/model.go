package model

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

type CreateComment struct {
	ItemID  int     `json:"itemid"`
	UserID  int     `json:"userid"`
	Pros    string  `json:"pros,omitempty"`
	Cons    string  `json:"cons,omitempty"`
	Comment string  `json:"comment,omitempty"`
	Rating  float64 `json:"rating"`
}

type CommentDB struct {
	ID      int     `json:"id"`
	ItemID  int     `json:"itemid"`
	UserID  int     `json:"userid"`
	Pros    string  `json:"pros,omitempty"`
	Cons    string  `json:"cons,omitempty"`
	Comment string  `json:"comment,omitempty"`
	Rating  float64 `json:"rating"`
}

type Comment struct {
	UserID     int     `json:"userid"`
	Username   string  `json:"username"`
	UserAvatar string  `json:"avatar,omitempty"`
	Pros       string  `json:"pros,omitempty"`
	Cons       string  `json:"cons,omitempty"`
	Comment    string  `json:"comment,omitempty"`
	Rating     float64 `json:"rating"`
}

type Promocode struct {
	Promocode string `json:"promocode"`
}

type Mail struct {
	Type        string `json:"type"`
	Username    string `json:"username"`
	Useremail   string `json:"usermail"`
	Promocode   string `json:"promocode,omitempty"`
	OrderStatus string `json:"orderstatus,omitempty"`
	OrderID     int    `json:"orderid,omitempty"`
}
