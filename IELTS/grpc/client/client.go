package client

import (
	"github/http/copy/task3/generated/session"
	"github/http/copy/task3/generated/test"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	LoginService() session.AuthServiceClient
	TestService() test.TestServiceClient
}

type GRPCClient struct {
	NewAuthService session.AuthServiceClient
	NewTestService test.TestServiceClient
}

func (g *GRPCClient) LoginService() session.AuthServiceClient {
	return g.NewAuthService
}

func (g *GRPCClient) TestService() test.TestServiceClient {
	return g.NewTestService
}

func NewGRPCClient() (ServiceManager, error) {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &GRPCClient{
			NewAuthService: session.NewAuthServiceClient(conn),
			NewTestService: test.NewTestServiceClient(conn),
		},
		err
}
