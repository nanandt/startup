package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler{
	return &userHandler{userService, authService}
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

	if err != nil{
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil{
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	formatter := user.FormatUser(newUser, token)

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

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil{
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Loggin Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(h *userHandler) CheckEmailAvailability(c *gin.Context){
	// ada input email dari user
	// input email di mapping ke struct input
	// struct input di parsing ke service
	// service memanggil repository - apakah email sudah ada atau belum
	// repository query ke database

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {

		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	// metaMessage := "Email has been registered"

	// if IsEmailAvailable {
	// 	metaMessage = "Email is Available"
	// }

	var metaMessage string

	if IsEmailAvailable {
		metaMessage = "Email is Available"
	} else {
		metaMessage = "Email has been registered"
	}


	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK,response)
	return
}

func (h *userHandler) UploadAvatar(c *gin.Context){
	// input file  dari user
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// image/namaFile.png
	// menjadi
	// images/ID-namaFile.png

	// path := "images/" + file.Filename

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)


	// simpan gambar di folder "images/"
	// di service panggil repository utk menntukn siapa user yg mengakses
	// JWT (sementara hardcode, seakan2 user yg login ID = 1)
	// repo ambil data user yg id nya = 1
	// repo update data user simpan lokasi file (yg disimpan adalah pathnya)
}