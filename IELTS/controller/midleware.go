package controller

import (
	"errors"
	"github/http/copy/task3/config"
	"github/http/copy/task3/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (c *Controller) Midleware(ctx *gin.Context) {
	if ctx.Request.Method == "PRI" && ctx.Request.RequestURI == "*" {
		ctx.AbortWithStatus(http.StatusOK)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tokenStr := strings.Split(authHeader, " ")
	if tokenStr[0] != "Bearer" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
	}

	claims := security.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr[1], &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.Cfg().JWTsecretkey, nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.Set("userId", claims.UserId)
	ctx.Next()
}
