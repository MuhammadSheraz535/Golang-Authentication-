package service

import (
	"net/http"

	log "github.com/MuhammadSheraz535/golang-authentication/logger"
	"github.com/gin-gonic/gin"
)

func (s *SignupService) Logout(c *gin.Context) {
	log.Info("Initializing Logout User handler function...")
	c.SetCookie("Authorization", "", -3600*2, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout"})

}
