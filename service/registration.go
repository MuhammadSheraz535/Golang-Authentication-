package service

import (
	"net/http"
	"strings"
	"time"

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
func (s *SignupService) Register(c *gin.Context) {

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
	time := time.Now()
	newUser := models.SignUp{
		CommonModel: models.CommonModel{
			CreatedAt: time,
			UpdatedAt: time,
		},
		Name:            user.Name,
		Email:           strings.ToLower(user.Email),
		DOB:             user.DOB,
		Password:        hashedPassword,
		PasswordConfirm: `json:"-"`,
	}
	newUser, err = controller.CreateUser(s.Db, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
