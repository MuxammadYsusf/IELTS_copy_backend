package service

import (
	"database/sql"
	"github/http/copy/task2/generated/session"
	"github/http/copy/task2/generated/test"
	client "github/http/copy/task2/grpc/client"
	"github/http/copy/task2/storage"
)

type QuestionService struct {
	Questionstorage storage.NewStorage
	Service         client.ServiceManager
	test.UnimplementedQuestionServiceServer
}

type LoginService struct {
	Loginstorage storage.NewStorage
	Service      client.ServiceManager
	session.UnimplementedAuthServiceServer
}

func NewQuestionService(db *sql.DB, Service client.ServiceManager) *QuestionService {
	return &QuestionService{
		Questionstorage: storage.StorageI(db),
		Service:         Service,
	}
}

func NewLoginService(db *sql.DB, Service client.ServiceManager) *LoginService {
	return &LoginService{
		Loginstorage: storage.StorageI(db),
		Service:      Service,
	}
}
