package controller

import (
	"errors"
	"fmt"

	"github.com/MuhammadSheraz535/golang-authentication/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

// create user
func CreateUser(db *gorm.DB, user models.SignUp) (models.SignUp, error) {
	if db.Model(models.SignUp{}).Where("email = ?", user.Email).Find(&user).RowsAffected > 0 {
		return user, errors.New("email is already registered")
	}

	if err := db.Model(models.SignUp{}).Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
