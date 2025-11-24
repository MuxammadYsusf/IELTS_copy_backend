package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
	SSLMode          string
}

func Load() Config {
	cfg := Config{}
	cfg.HTTPPort = ":8080"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "postgres"
	cfg.PostgresDatabase = "catalog_service"
	cfg.PostgresPassword = "postgres"
	cfg.PostgresPort = "5432"
	cfg.SSLMode = "disable"

	return cfg
}
