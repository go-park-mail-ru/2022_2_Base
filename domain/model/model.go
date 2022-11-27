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
	Username string  `json:"username"`
	Pros     string  `json:"pros,omitempty"`
	Cons     string  `json:"cons,omitempty"`
	Comment  string  `json:"comment,omitempty"`
	Rating   float64 `json:"rating"`
}
