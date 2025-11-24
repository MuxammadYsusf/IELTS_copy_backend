package service_test

import (
	"github/http/copy/task3/config"
	"github/http/copy/task3/grpc/client"
	"github/http/copy/task3/grpc/service"
	"github/http/copy/task3/pkg/db"
	"os"
	"testing"
)

var (
	loginService *service.LoginService
)

func TestMain(m *testing.M) {
	cfg := config.Cfg()

	conn, err := db.Postgres(&cfg)
	if err != nil {
		os.Exit(1)
	}

	clients, err := client.NewGRPCClient()
	if err != nil {
		os.Exit(1)
	}

	loginService = service.NewLoginService(conn, clients)

	code := m.Run()
	os.Exit(code)
}
