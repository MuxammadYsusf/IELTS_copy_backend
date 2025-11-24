package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github/http/copy/task3/generated/session"
	"github/http/copy/task3/models"

	"golang.org/x/crypto/bcrypt"
)

type login struct {
	db *sql.DB
}

type LoginRepo interface {
	Register(ctx context.Context, req *session.RegisterRequest) (*session.RegisterResponse, error)
	Login(ctx context.Context, req *session.LoginRequest) (*session.LoginResponse, error)
}

func NewLogin(db *sql.DB) LoginRepo {
	return &login{
		db: db,
	}
}

func (l *login) Register(ctx context.Context, req *session.RegisterRequest) (*session.RegisterResponse, error) {

	tx, err := l.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	hashedPass, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users(name, phone_number, password) VALUES($1, $2, $3)`

	_, err = tx.Exec(
		query,
		req.Name,
		req.PhoneNumber,
		hashedPass,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &session.RegisterResponse{
		Message: "success",
	}, nil

}

func (l *login) Login(ctx context.Context, req *session.LoginRequest) (*session.LoginResponse, error) {

	var user models.User

	tx, err := l.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	err = tx.QueryRow(`SELECT id, name, password FROM users WHERE name = $1`, req.Name).Scan(&user.Id, &user.Name, &user.Password)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return nil, fmt.Errorf("invalid name or password")
	} else if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid name or password")
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
