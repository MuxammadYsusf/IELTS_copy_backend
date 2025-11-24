package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github/http/copy/task1/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db *sql.DB, creds models.LoginRequest) error {
	var storedPassword string

	// Получаем хэшированный пароль из базы данных
	err := db.QueryRow(`SELECT password FROM authors WHERE author_id = $1`, creds.AuthorID).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Пользователь с данным AuthorID не найден")
			return fmt.Errorf("user not found")
		}
		fmt.Println("Ошибка при выполнении запроса:", err)
		return err
	}

	fmt.Println("Извлеченный хэшированный пароль:", storedPassword)

	// Проверяем введённый пароль против хэшированного пароля
	if !CheckPasswordHash(creds.Password, storedPassword) {
		fmt.Println("Пароль не совпадает")
		return fmt.Errorf("invalid password")
	}

	return nil
}

// Хэширование пароля для хранения в базе данных
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Проверка соответствия введённого пароля и хэшированного пароля
func CheckPasswordHash(password, hash string) bool {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PASS:>>> ", string(bytes), hash)
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Функция для обновления пароля пользователя
func UpdateUserPassword(db *sql.DB, user struct {
	AuthorID    int    `json:"author_id"`
	NewPassword string `json:"password"`
}) error {

	// Хэшируем новый пароль перед сохранением в базу данных
	hashedPassword, err := HashPassword(user.NewPassword)
	if err != nil {
		return err
	}

	//$2a$10$ThPOD546FxZamCy1Uv6QZuiBqFgtwCTyLuSS3E4BGvY3nmjcJl2JS - hashed
	//$2a$10$yDUVNicryVjPHs6PlnJWcOvpWXLBOWihGDnVtsWh8VZ6SF7YSVUE2

	query := `UPDATE authors SET password = $2 WHERE author_id = $1`
	_, err = db.Exec(query, user.AuthorID, hashedPassword)
	if err != nil {
		return err
	}

	fmt.Printf("Пароль обновлен для AuthorID: %d с хэшированным паролем: %s\n", user.AuthorID, hashedPassword)
	return nil
}
