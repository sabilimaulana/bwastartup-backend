package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("Connection to database is success")

	// userRepository := user.NewRepository(db)
	// userService := user.NewService(userRepository)

	// userInput := user.RegisterUserInput{
	// 	Name:       "Shanks",
	// 	Occupation: "Pirate",
	// 	Email:      "shanks@gmail.com",
	// 	Password:   "123456",
	// }

	// userService.RegisterUser(userInput)
}
