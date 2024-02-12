package model

type User struct {
	ID       *int   `json:"id,omitempty"` // Pointer to int, making it optional and omit it in JSON if nil
	Username string `json:"username"`
	Email    string `json:"email"`
}
