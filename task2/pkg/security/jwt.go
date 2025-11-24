package security

import (
	"github/http/copy/task2/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJWT(userID int) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.Cfg().JWTsecretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
