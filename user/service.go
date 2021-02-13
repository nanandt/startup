package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
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
		return user, errors.New("User Not Found on that Email")
	}

	// mencocokan password yg dimasukan user / yg sudah diregistrasi
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil{
		return user, errors.New("Your Password is False")
	}


	// jika passwordnya  matched
	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput)(bool,error){
	// mengambil email dari input
	email := input.Email

	// cari email yg diinput oleh user
	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return false, err
	}

	// apakah email yg diinput oleh user ada usernya atau tidak
	if user.ID == 0{ // user tidak ditemukan (email is available)
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error){
	// dapatkan user berdasarkan ID
	// user update atribut avatar file name
	// simpan perubahan avatar file name

	user, err := s.repository.FindByID(ID)
	if err != nil{
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}