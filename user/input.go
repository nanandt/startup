package user

// RegisterUserInput is ...
type RegisterUserInput struct {
	// struct ini mewakili apa yg diinputkan oleh user (form yg ada di frontend)
	Name       string
	Occupation string
	Email      string
	Password   string
}
