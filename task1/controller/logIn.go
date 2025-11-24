package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github/http/copy/task1/models"
	"github/http/copy/task1/storage"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	AuthorID int `json:"author_id"`
	jwt.StandardClaims
}

// Обработчик для логина
func (c *Controler) LoginHandler(ctx *gin.Context) {

	var creds models.LoginRequest
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	log.Println("Попытка входа для пользователя:", creds.AuthorID)

	// Проверяем введённые данные
	if err := storage.LoginHandler(c.db, creds); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "srver error"})
		return
	}

	// Генерация токена при успешном входе
	tokenString, err := GenerateJWTToken(creds.AuthorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in generating token"})
		return
	}

	// Отправляем токен клиенту
	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Функция для генерации JWT токена
func GenerateJWTToken(author_id int) (string, error) {
	expirationTime := time.Now().Add(3600 * time.Minute)
	claims := &Claims{
		AuthorID: author_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Добавление описания только авторизованным пользователям
func (c *Controler) AddDescriptionHandler(ctx *gin.Context) {
	var creds models.Models

	tokenAuthorID, err := ExtractTokenAuthorID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "srver error"})
		return
	}

	fmt.Println("TokenId", tokenAuthorID)
	fmt.Println("id", creds.Author_id)

	ctx.Header("author_id_from_token", fmt.Sprintf("%d", tokenAuthorID))
}

// Обновление пароля
func (c *Controler) UpdateUserPass(ctx *gin.Context) {

	var user struct {
		AuthorID    int    `json:"author_id"`
		NewPassword string `json:"password"`
	}

	// Декодируем запрос
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	// Обновляем пароль пользователя
	if err := storage.UpdateUserPassword(c.db, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on updating"})
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "password succcessfully updated"})

}
