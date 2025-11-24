package grpc

import (
	"github/http/copy/task2/generated/session"
	"github/http/copy/task2/generated/test"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	LoginService() session.AuthServiceClient
	QuestionService() test.QuestionServiceClient
}

type grpcClient struct {
	NewQuestionService test.QuestionServiceClient
	NewLoginService    session.AuthServiceClient
}

func (g *grpcClient) QuestionService() test.QuestionServiceClient {
	return g.NewQuestionService
}

func (g *grpcClient) LoginService() session.AuthServiceClient {
	return g.NewLoginService
}

func NewGRPCClient() (ServiceManager, error) {

	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &grpcClient{
			NewQuestionService: test.NewQuestionServiceClient(conn),
			NewLoginService:    session.NewAuthServiceClient(conn)},
		err
}
