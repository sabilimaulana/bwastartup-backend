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

	// userByEmail, err := userRepository.FindByEmail("zoroa@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

	router := gin.Default()

	v1 := router.Group("/api/v1")
	v1.POST("/users", userHandler.RegisterUser)

	router.Run()
}
