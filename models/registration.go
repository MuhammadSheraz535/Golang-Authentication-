package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SignUp struct {
	ID              uint64         `json:"id"`
	CreatedAt       time.Time      `gorm:"<-:create" json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	Name            string         `json:"name" validate:"required,min=3"`
	Email           string         `json:"email" validate:"required,email"`
	DOB             string         `json:"date_of_birth" validate:"required"`
	Password        string         `json:"Password" validate:"required,min=8"`
	PasswordConfirm string         `json:"confirm_password" validate:"required,min=8"`
}
type UserResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	DOB      string `json:"date_of_birth"`
	Password string `json:"Password"`
}

func (a *SignUp) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
