package main

import (
	"os"
	"reflect"
	"strings"

	"github.com/MuhammadSheraz535/golang-authentication/database"
	"github.com/MuhammadSheraz535/golang-authentication/logger"
	"github.com/MuhammadSheraz535/golang-authentication/routes"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load(".env")

	// Convert fe.Field() from StructField to json field for custom validation messages
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// Connect to the database
	database.Connect()

	// Initializing logger
	logger.TextLogInit()

	// Register all the routes
	server := routes.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = server.Run(":" + port)
}
