package config

type Config struct {
	HttpPort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	SSLMode          string
}

func Cfg() Config {
	cfg := Config{}

	cfg.HttpPort = ":8080"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "postgres"
	cfg.PostgresDatabase = "catalog_service"
	cfg.PostgresPassword = "postgres"
	cfg.PostgresPort = "5432"
	cfg.SSLMode = "disable"

	return cfg
}
