package user

// RegisterUserInput is ...
type RegisterUserInput struct {
	// struct ini mewakili apa yg diinputkan oleh user (form yg ada di frontend)
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// LoginInput is ...
type LoginInput struct{
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}

// CheckEmailInput is ...
type CheckEmailInput struct{
	Email string `json:"email" binding:"required,email"`
}

type FormCreateUserInput struct {
	Name		string `form:"name" binding:"required"`
	Email		string `form:"email" binding:"required,email"`
	Occupation	string `form:"occupation" binding:"required"`
	Password	string `form:"password" binding:"required"`
	Error		error
}