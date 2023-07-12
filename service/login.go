package service

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MuhammadSheraz535/golang-authentication/controller"
	log "github.com/MuhammadSheraz535/golang-authentication/logger"
	"github.com/MuhammadSheraz535/golang-authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login

func (s *SignupService) Login(c *gin.Context) {

	log.Info("Initializing Login User handler function...")

	// binding user
	var loginuser *models.Login
	if err := c.ShouldBindJSON(&loginuser); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := loginuser.Validate()
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

	//check user by email
	var user *models.SignUp
	if s.Db.Model(models.SignUp{}).Where("email = ?", strings.ToLower(loginuser.Email)).Find(&user).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
	// compare user password  with requested user password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginuser.Password)); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"name":user.Name,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}
	//set it cookie
	c.SetSameSite(http.SameSiteLaxMode)
	//set cookie and send it back
	c.SetCookie("Authorization", tokenString, 3600*2, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": tokenString})
}

//validate user credentials

func (s *SignupService) Validate(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user})
}
