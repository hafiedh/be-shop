package models

type (
	User struct {
		ID       int    `json:"id,omitempty"`
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,min=8,max=20,uppercase,lowercase,number,specialchar"`
	}
)
