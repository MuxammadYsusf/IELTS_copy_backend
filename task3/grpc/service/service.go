package service

import (
	"database/sql"
	"github/http/copy/task3/generated/session"
	"github/http/copy/task3/generated/test"
	"github/http/copy/task3/grpc/client"
	"github/http/copy/task3/postgres"
)

type TestService struct {
	testPostgres postgres.NewPostgres
	service      client.ServiceManager
	test.UnimplementedTestServiceServer
}

type LoginService struct {
	loginPostgres postgres.NewPostgres
	service       client.ServiceManager
	session.UnimplementedAuthServiceServer
}

func NewTestService(db *sql.DB, service client.ServiceManager) *TestService {
	return &TestService{
		testPostgres: postgres.NewPostgresI(db),
		service:      service,
	}
}

func NewLoginService(db *sql.DB, service client.ServiceManager) *LoginService {
	return &LoginService{
		loginPostgres: postgres.NewPostgresI(db),
		service:       service,
	}
}
