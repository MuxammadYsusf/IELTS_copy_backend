package security

import (
	"time"

	"github/http/copy/task3/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userId int) (string, error) {

	expiredTime := time.Now().Add(time.Hour * 24)

	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(config.Cfg().JWTsecretkey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
