package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository

}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil{
		return user,err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil{
		return  newUser,err
	}

	return newUser, nil

}

// mapping struct unput ke struct User
// simpan struct User melalui repository

func (s *service) Login(input LoginInput) (User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil{
		return user, err
	}

	// pengecekan user ada nilainya atau tidak
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	// mencocokan password yg dimasukan user / yg sudah diregistrasi
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil{
		return user, err
	}


	// jika passwordnya  matched
	return user, nil





}