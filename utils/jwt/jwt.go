package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secret = "faKXAsQqpwiVmtkgHNsJ7poQiuDY2dckweZeGovsWCpG1AOE9u4qlUYPU90ZVGNX"
)

func CreateToken(id uint, companyID int) string {

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["ID"] = id
	token.Claims["company_id"] = companyID
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {

	}
	return tokenString
}

func ParseToken(tokenString string) (int, int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil && token.Valid {
		if token.Claims["ID"] != nil {
			return int(token.Claims["ID"].(float64)), int(token.Claims["company_id"].(float64)), nil
		} else {
			return -1, -1, errors.New("Bad Token!")
		}

	} else {
		return -1, -1, errors.New("Bad Token!")
	}

}
