package structs

// ToDo struct
type ToDo struct {
	UserID    float64 `json:"userId"`
	ID        float64 `json:"id"`
	Title     string  `json:"title"`
	Completed bool    `json:"completed"`
}

// ToDos List of ToDo struct
type ToDos []ToDo
