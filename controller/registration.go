package controller

import (
	"errors"
	"fmt"

	log "github.com/MuhammadSheraz535/golang-authentication/logger"
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

//Get All Users

func GetAllUsers(db *gorm.DB, name string, users []models.UserResponse) ([]models.UserResponse, error) {
	log.Info("Get all register users")
	db = db.Model(models.SignUp{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}

func CheckUserExist(db *gorm.DB, user models.SignUp, id uint64) error {
	log.Info("Check user exist by ID")

	err := db.Model(&models.SignUp{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil

}

func DeleteRegisterUser(db *gorm.DB, user models.SignUp) error {
	log.Info("Delete User")
	err := db.Delete(&user).Error
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil

}
