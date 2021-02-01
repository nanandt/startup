package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}

	userInput.Name = "save from service"
	userInput.Email = "save@gmail.com"
	userInput.Occupation = "pedagang"
	userInput.Password = "password"

	userService.RegisterUserInput(userInput)



	// input dari user
	//handler : mapping input dari user -> struct input
	//service : melakukan mapping dari struct input ke struct User
	//repository
	//db


}