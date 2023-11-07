package scurity

import (
	"api/model/domain"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func PassEncrypt(pass domain.Register) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass.Password), 20)
	if err != nil {
		return "", errors.New("failed to create hash")
	}
	return string(hash), nil
}

func ClaimsJwt(logData domain.Login, userData string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(logData.Password), []byte(userData))
	if err != nil {
		return "", errors.New("Password salah")
	}

	// Membuat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": logData.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte("asc22332afg0061729940qqr"))
	if err != nil {
		return "", errors.New("Gagal membuat token JWT")
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
