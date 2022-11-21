package model

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

type CreateComment struct {
	ItemID    int    `json:"itemid"`
	UserID    int    `json:"userid"`
	Worths    string `json:"worths"`
	Drawbacks string `json:"drawbacks"`
	Comment   string `json:"comment"`
}

type Comment struct {
	ID        int     `json:"id"`
	ItemID    int     `json:"itemid"`
	UserID    int     `json:"userid"`
	Worths    string  `json:"worths"`
	Drawbacks string  `json:"drawbacks"`
	Comment   string  `json:"comment"`
	Rating    float64 `json:"rating,omitempty"`
}
