package security

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super-secret-key")

func CreateToken(username string, phone string, email string) (string, error) {
	log.Println("[INFO]: Creating token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"phone":    phone,
			"email":    email,
			"exp_date": time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("[ERROR]: Error to generate token")
		return "", err
	}
	log.Println("[INFO]: Token created successfully")
	return tokenString, nil
}
