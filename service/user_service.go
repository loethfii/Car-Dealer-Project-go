package service

import (
	"tugas_akhir_course_net/auth"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type UserService interface {
	RegisterUser(userReq models.UserRequest) (models.User, error)
	LoginUser(user models.TokenRequest) (string, error)
	CredentialEmailPassword(user models.TokenRequest) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(userReq models.UserRequest) (models.User, error) {
	hashUser := helper.User{}
	password := hashUser.HashPassword(userReq.Password)

	NewUser := models.User{
		Name:     userReq.Name,
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: password,
		Role:     userReq.Role,
	}

	NewUser, err := s.userRepository.RegisterUser(NewUser)

	return NewUser, err
}

func (s *userService) LoginUser(user models.TokenRequest) (string, error) {

	role, err := s.userRepository.LoginUser(user)

	//generate token
	tokenString, err := auth.GenereteJWT(user.Email, user.Password, role)

	return tokenString, err
}

func (s *userService) CredentialEmailPassword(login models.TokenRequest) error {

	err := s.userRepository.CredentialEmailPassword(login)

	return err
}
