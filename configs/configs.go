package configs

import (
	"os"
)

type Config struct {
	// Service configs
	AppEnv             string `json:"APP_ENV" env:"APP_ENV"`
	HostName           string `json:"HOSTNAME" env:"HOSTNAME"`
	ServerExternalPort string `json:"SERVER_EXTERNAL_PORT" env:"SERVER_EXTERNAL_PORT"`
	// Postgres configs
	PostgresDBStr string `json:"POSTGRESS_DB_STR" env:"POSTGRESS_DB_STR"`
}

func NewConfig() *Config {
	return &Config{
		AppEnv:             getEnv("APP_ENV", ""),
		HostName:           getEnv("SERVER_TOKEN", ""),
		ServerExternalPort: getEnv("SERVER_EXTERNAL_PORT", ""),
		PostgresDBStr:      getEnv("POSTGRESS_DB_STR", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
