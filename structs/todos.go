package structs

// ToDo struct
type ToDo struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ToDos List of ToDo struct
type ToDos []ToDo
