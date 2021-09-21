package main

import (
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
	userHandler := handler.NewHandler(userService)

	// userInput := user.LoginInput{
	// 	Email:    "zoro@gmail.com",
	// 	Password: "123456",
	// }

	// user, err := userService.Login(userInput)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(user.Name)
	// }

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/users", userHandler.RegisterUser)

	router.Run()
}
