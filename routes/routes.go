package routes

import (
	"os"
	"strconv"

	"github.com/MuhammadSheraz535/golang-authentication/service"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	isCorsEnabled, _ := strconv.ParseBool(os.Getenv("ENABLE_CORS"))
	if isCorsEnabled {
		_ = router.SetTrustedProxies(nil)

		router.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
				"Accept",
				"Origin",
				"Cache-Control",
			},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
			AllowCredentials: false,
		}))
	}

	s := service.NewSignupService()

	v1 := router.Group("/v1")

	user := v1.Group("/register")
	{
		user.POST("", s.Register)
	}

	return router
}