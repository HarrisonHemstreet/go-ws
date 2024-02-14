package model

type User struct {
	ID       *int    `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
}
