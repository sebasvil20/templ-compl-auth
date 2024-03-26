package model

type User struct {
	ID       string
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Admin    bool
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
