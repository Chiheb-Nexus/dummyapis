package structs

// UserAddressGeo is a part of User struct definition
type UserAddressGeo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

// UserAddress is a part of User struct definition
type UserAddress struct {
	Street  string         `json:"street"`
	Suite   string         `json:"suite"`
	City    string         `json:"city"`
	Zipcode string         `json:"zipcode"`
	Geo     UserAddressGeo `json:"geo"`
}

// UserCompagny is part of User struct definition
type UserCompagny struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

// User struct for API: https://jsonplaceholder.typicode.com/users
type User struct {
	ID       int64        `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Address  UserAddress  `json:"address"`
	Phone    string       `json:"phone"`
	Website  string       `json:"website"`
	Compagny UserCompagny `json:"company"`
}

// Users List of Users
type Users []User
