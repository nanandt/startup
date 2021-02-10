package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas diparsing sbg parameter service

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil{

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser,err := h.userService.RegisterUser(input)

	formatter := user.FormatUser(newUser, "tempat token")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	if err != nil{
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context){
	// user melakukan input email dan password
	// input ditangkap di handler
	// mapping dari input user ke input struct
	// input struct passing ke service
	// di service mencari dgn bantuan repositori user dgn email tertentu
	// mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil{
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "token")

	response := helper.APIResponse("Loggin Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func CheckEmailAvailability(c *gin.Context){
	// ada input email dari user
	// input email di mapping ke struct input
	// struct input di parsing ke service
	// service memanggil repository - apakah email sudah ada atau belum
	// repository query ke database

}