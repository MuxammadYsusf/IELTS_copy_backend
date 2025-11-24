package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github/http/copy/task2/generated/session"
	"github/http/copy/task2/models"

	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	db *sql.DB
}

type LoginRepo interface {
	Register(ctx context.Context, req *session.RegisterRequest) (*session.RegisterResponse, error)
	Login(ctx context.Context, req *session.LoginRequest) (*session.LoginResponse, error)
	UpdatePassword(ctx context.Context, req *session.UpdatePasswordRequest) (*session.UpdatePasswordResponse, error)
}

func NewSession(db *sql.DB) LoginRepo {
	return &Session{
		db: db,
	}
}

func (s *Session) Register(ctx context.Context, req *session.RegisterRequest) (*session.RegisterResponse, error) {

	hashedPass, err := HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := `INSERT INTO users(name, phonenumber, password) VALUES($1, $2, $3)`

	_, err = s.db.Exec(
		query,
		req.Name,
		req.PhoneNumber,
		hashedPass,
	)
	if err != nil {
		return nil, err
	}

	return &session.RegisterResponse{
		Message: "success",
	}, nil
}

func (s *Session) Login(ctx context.Context, req *session.LoginRequest) (*session.LoginResponse, error) {
	var (
		user models.User
	)

	err := s.db.QueryRow(`SELECT id, phonenumber, name, password FROM users WHERE name=$1`, req.Name).Scan(&user.Id, &user.PhoneNumber, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("user not found!")
		}
		return nil, err
	}

	fmt.Println(user)

	if !CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("passwords mismatch")
	}

	return &session.LoginResponse{
		Message: "success",
		UserId:  int32(user.Id),
	}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err == nil
	}

	return err == nil
}

func (s *Session) UpdatePassword(ctx context.Context, req *session.UpdatePasswordRequest) (*session.UpdatePasswordResponse, error) {

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	query := `UPDATE users SET password = $2 WHERE name = $1`
	_, err = s.db.Exec(query, req.Name, hashedPassword)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Пароль обновлен для user: %s с хэшированным паролем: %s\n", req.Name, hashedPassword)
	return &session.UpdatePasswordResponse{
		Message: "success",
	}, nil
}
