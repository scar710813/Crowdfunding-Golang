package main

import (
	// "fmt"
	"log"
	// "net/http"
	"nura-fund/handler"
	"nura-fund/user"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
func main() {
	// Database Connection
	dsn := "root:@tcp(127.0.0.1:3306)/nura_fund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Register User Repository
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// Routing Register User
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}