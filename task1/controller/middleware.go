package controller

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (c *Controler) Middleware(ctx *gin.Context) {

	// Выполнение логики middleware перед вызовом следующего обработчика
	c.AddDescriptionHandler(ctx)

	ctx.Next()

}

func ExtractTokenAuthorID(ctx *gin.Context) (int, error) {

	var err error

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "authorization header is missing!"})
		return 0, err
	}

	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString := strings.Split(authHeader, " ")
	fmt.Println("token string")
	if tokenString[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token format!"})
		return 0, err
	}

	fmt.Println("AuthHeader>>>> ", authHeader)
	fmt.Println("tokenString>>>> ", tokenString)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("|:(")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token!"})
		return 0, err

	}

	return claims.AuthorID, nil
}
