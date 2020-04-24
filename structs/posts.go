package dummyapis

// Post struct
type Post struct {
	UserID float64 `json:"userId"`
	ID     float64 `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
}

// Posts list of Post struct
type Posts []Post
