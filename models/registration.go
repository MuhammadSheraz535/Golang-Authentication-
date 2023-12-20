package models

import (
	"github.com/go-playground/validator/v10"
)

type SignUp struct {
	CommonModel
	Name            string `json:"name" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	DOB             string `json:"date_of_birth" validate:"required"`
	Password        string `json:"Password" validate:"required,min=8"`
	PasswordConfirm string `json:"confirm_password,omitempty" validate:"required,min=8"`
}

func (a *SignUp) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
