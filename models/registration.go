package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint64         `json:"id"`
	CreatedAt time.Time      `gorm:"<-:create" json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type SignUp struct {
	CommonModel
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email" gorm:"unique"`
	DOB             string `json:"date_of_birth" validate:"required"`
	Password        string `json:"Password" validate:"required,min=8"`
	PasswordConfirm string `json:"confirm_password" validate:"required,min=8"`
}

func (a *SignUp) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
