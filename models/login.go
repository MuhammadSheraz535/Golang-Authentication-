package models

import "github.com/go-playground/validator/v10"

type Login struct {
	CommonModel
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=8"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	DOB      string `json:"date_of_birth"`
	Password string `json:"Password"`
}

func (a *Login) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
