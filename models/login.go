package models

import "github.com/go-playground/validator/v10"

type Login struct {
	CommonModel
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=8"`
}

func (a *Login) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
