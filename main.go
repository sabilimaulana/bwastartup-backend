package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is success")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewHandler(userService, authService)

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsJnR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.m4nqYzcy0cQ2k8QGyHoeiN7rP8H8h1EgkQcRQupYUjk")
	if err != nil {
		fmt.Println(err.Error())
	}

	if token.Valid {
		fmt.Println("Valid Token")
	} else {
		fmt.Println("Invalid Token")
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/users", userHandler.RegisterUser)
	v1.POST("/sessions", userHandler.Login)
	v1.POST("/email_checkers", userHandler.CheckEmailAvailability)
	v1.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
