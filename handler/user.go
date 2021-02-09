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
}