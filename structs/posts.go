package structs

// Post struct
type Post struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Posts list of Post struct
type Posts []Post
