package repository

import (
	"errors"
	"gorm.io/gorm"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
)

type UserRepository interface {
	RegisterUser(user models.User) (models.User, error)
	LoginUser(tokenUser models.TokenRequest) (int, error)
	CredentialEmailPassword(reqTOken models.TokenRequest) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) RegisterUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) LoginUser(tokenUser models.TokenRequest) (int, error) {
	var user models.User

	checkEmail := r.db.Where("email = ?", tokenUser.Email).First(&user)
	if checkEmail.Error != nil {
		return 0, errors.New("Password Not Match")
	}

	return user.Role, nil
}

func (r *userRepository) CredentialEmailPassword(reqTOken models.TokenRequest) error {
	var user models.User
	r.db.Where("email = ?", reqTOken.Email).First(&user)
	var forCheckPassword helper.User
	res := forCheckPassword.CheckPassword(reqTOken.Password, user.Password)

	if user.Email != reqTOken.Email {
		return errors.New("Akun tidak ditemukan")
	} else if res == false {
		return errors.New("Password salah!")
	}

	return nil
}
