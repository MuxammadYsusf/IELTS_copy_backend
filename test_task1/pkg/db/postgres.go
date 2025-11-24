package db

import (
	"database/sql"
	"fmt"

	"github/http/copy/test_task1/config"

	_ "github.com/lib/pq"
)

func NewConnectPostgres(cfg *config.Config) (*sql.DB, error) {
	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
		cfg.SSLMode,
	)

	conn, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
