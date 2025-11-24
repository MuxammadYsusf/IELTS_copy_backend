package controller

import (
	"errors"
	"github/http/copy/task2/config"
	"github/http/copy/task2/pkg/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (c *Handler) AddDescHandler(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == " " {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error!"})
		return
	}

	tokenStr := strings.Split(authHeader, " ")
	if tokenStr[0] != "Bearer" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error!"})
		return
	}

	claims := security.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr[1], &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.Cfg().JWTsecretkey, nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token!"})
		return
	}

	ctx.Set("user_id", claims.UserId)
	ctx.Next()

}
