package helper

import (
	"crypto/sha1"
	"fmt"
	"tugas_akhir_course_net/models"
)

type User struct {
	user models.User
}

func (user *User) HashPassword(password string) string {
	var sha = sha1.New()
	sha.Write([]byte(password))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	user.user.Password = encryptedString

	return user.user.Password
}

func (user *User) CheckPassword(passwordString, passwordEncrypt string) bool {
	var sha = sha1.New()
	sha.Write([]byte(passwordString))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	return encryptedString == passwordEncrypt
}
