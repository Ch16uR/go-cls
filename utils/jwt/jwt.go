package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secret = "frontolUp"
)

func CreateToken(uuid string) string {

	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["uuid"] = uuid
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {

	}
	return tokenString
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil && token.Valid {
		if token.Claims["uuid"] != nil {
			return token.Claims["uuid"].(string), nil
		} else {
			return "", errors.New("Error!")
		}

	} else {
		return "", errors.New("Error!")
	}

}
