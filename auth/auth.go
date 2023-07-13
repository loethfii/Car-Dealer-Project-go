package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"time"
)

func jwtKeyEnv() string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Gagal membaca file konfigurasi: %v", err)
	}

	jwtKey := viper.GetString("JWT_KEY")

	return jwtKey
}

var keyEnv = jwtKeyEnv()

var jwtKey = []byte(keyEnv)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func GenereteJWT(email, username string, role int) (string, error) {
	expTime := time.Now().Add(time.Hour * 24 * 365)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ValidateToken(signedToken string) (email string, role int, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)

	//jika claim gagal
	if !ok {
		err = errors.New("Cloud parse claims for token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}

	role = claims.Role
	email = claims.Email

	return
}
