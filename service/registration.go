package service

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/MuhammadSheraz535/golang-authentication/controller"
	"github.com/MuhammadSheraz535/golang-authentication/database"
	log "github.com/MuhammadSheraz535/golang-authentication/logger"
	"github.com/MuhammadSheraz535/golang-authentication/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignupService struct {
	Db *gorm.DB
}

func NewSignupService() *SignupService {
	db := database.DB
	err := db.AutoMigrate(&models.SignUp{})
	if err != nil {
		panic(err)
	}
	return &SignupService{Db: db}
}

// user Signup
func (s *SignupService) RegisterUser(c *gin.Context) {
	log.Info("Initializing Register User handler function...")
	// binding user
	var user *models.SignUp

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.Validate()
	if err != nil {
		errs, ok := controller.ErrValidationSlice(err)
		if !ok {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Error(err.Error())
		if len(errs) > 1 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errs[0]})
		}
		return
	}

	//check password and confirm password are same

	if user.Password != user.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return

	}
	hashedPassword, err := controller.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.SignUp{

		Name:     user.Name,
		Email:    strings.ToLower(user.Email),
		DOB:      user.DOB,
		Password: hashedPassword,
	}
	_, err = controller.CreateUser(s.Db, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

//Get all register users

func (s *SignupService) GetAllRegisterUsers(c *gin.Context) {
	log.Info("Initializing Get All Register User handler function...")
	var users []models.UserResponse
	name := c.Query("name")
	user, err := controller.GetAllUsers(s.Db, name, users)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

// GET users by id

func (s *SignupService) GetUsersById(c *gin.Context) {
	log.Info("Initializing Get User by id handler function...")
	var user models.UserResponse
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	users, err := controller.CheckUserExist(s.Db, user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("User record not found against the given id")
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

// Delete User
func (s *SignupService) DeleteRegisterUser(c *gin.Context) {

	log.Info("Initializing Delete User handler function...")

	var user models.UserResponse
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	user.ID = id
	//check user exists in database
	_, err := controller.CheckUserExist(s.Db, user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("User record not found against the given id")
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var users models.SignUp
	user_id := user.ID
	// delete user from database
	err = controller.DeleteRegisterUser(s.Db, users,user_id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
