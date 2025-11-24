package server

import (
	"database/sql"

	"github/http/copy/task3/generated/session"
	"github/http/copy/task3/generated/test"

	client "github/http/copy/task3/grpc/client"
	service "github/http/copy/task3/grpc/service"

	"google.golang.org/grpc"
)

func NewServer(db *sql.DB, serviceManager client.ServiceManager) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	test.RegisterTestServiceServer(grpcServer, service.NewTestService(db, serviceManager))
	session.RegisterAuthServiceServer(grpcServer, service.NewLoginService(db, serviceManager))

	return
}
