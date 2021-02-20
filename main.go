package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	campaignRepository := campaign.NewRepository(db)

	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)


	// test campaign
	// campaigns, err := campaignRepository.FindAll()
	// campaigns, err = campaignRepository.FindByUserID(1)
	// fmt.Println("debug")
	// fmt.Println("debug")
	// fmt.Println("debug")
	// fmt.Println(len(campaigns))

	// for _, campaign := range campaigns{
	// 	fmt.Println(campaign.Name)
	// 	if len(campaign.CampaignImages) > 0 {
	// 		fmt.Println(campaign.CampaignImages[0].FileName)

	// 	}
	// }

	// test generate token jwt
	authService := auth.NewService()
	// fmt.Println(authService.GenerateToken(2000))

	// test save avatar file name
	// userService.SaveAvatar(1, "images/1-profile-png")

	// cek email user
	// userByEmail, err := userRepository.FindByEmail("jakaaa@gmail.com")
	// if err != nil{
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0{
	// 	fmt.Println("User tidak ditemukan")
	// }
	// fmt.Println(userByEmail.Name)


	// test email dan password login
	// input := user.LoginInput{
	// 	Email: "harunnn@gmail.com",
	// 	Password: "password",
	// }
	// user, err := userService.Login(input)
	// if err != nil{
	// 	fmt.Println("terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)
	// test validate token dan ex.token invalid
	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0fQ.4VyKdWGh0T4g6_9zbr4U_FJqPuFUi6nOOEoFmff0KiA")
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// }

	// if token.Valid{
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// } else {
	// 	fmt.Println("INVALID")
	// }


	// input := campaign.CreateCampaignInput{}
	// input.Name = "Penggalangan Dana Banjir"
	// input.ShortDescription = "short"
	// input.Description = "panjaaaaaang"
	// input.GoalAmount = 100000000
	// input.Perks = "banjir di jakarta, gempa bumi di jogja, tsunami di palu"
	// inputUser, _ := userService.GetUserByID(1)
	// input.User = inputUser

	// _, err = campaignService.CreateCampaign(input)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }


	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	router.Static("/images", "./images")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService),campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService),campaignHandler.UpdateCampaign)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc{
	return func (c *gin.Context){
	authHeader := c.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer"){
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized,response)
		return
	}

	// Bearer tokentokentoken
	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2{
		tokenString = arrayToken[1]
	}

	token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		c.Set("currentUser", user)
	}
}

// ambil nilai header Authorization: Bearer tokentokentoken
// dari header Authorization, kita ambil nilai tokennya saja
// validasi token menggunakan ValidateToken yg suda dibuat
// jika valid ambil user_id
// ambil user dari db  berdasarkan user_id lewat service
// jika usernya ada, set context isinya user