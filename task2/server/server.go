package server

import (
	"database/sql"
	"github/http/copy/task2/generated/session"
	"github/http/copy/task2/generated/test"

	client "github/http/copy/task2/grpc/client"
	service "github/http/copy/task2/grpc/service"

	"google.golang.org/grpc"
)

func SetUpServer(db *sql.DB, ServiceManager client.ServiceManager) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	test.RegisterQuestionServiceServer(grpcServer, service.NewQuestionService(db, ServiceManager))
	session.RegisterAuthServiceServer(grpcServer, service.NewLoginService(db, ServiceManager))

	return
}
