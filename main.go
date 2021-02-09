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
	dsn := "root:root@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// cek email user
	// userByEmail, err := userRepository.FindByEmail("jakaaa@gmail.com")
	// if err != nil{
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0{
	// 	fmt.Println("User tidak ditemukan")
	// }
	// fmt.Println(userByEmail.Name)

	input := user.LoginInput{
		Email: "harunnn@gmail.com",
		Password: "password",
	}
	user, err := userService.Login(input)
	if err != nil{
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()


	// input dari user
	//handler : mapping input dari user -> struct input
	//service : melakukan mapping dari struct input ke struct User
	//repository
	//db


}