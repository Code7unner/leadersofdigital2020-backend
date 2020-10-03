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
	PostgresDBStr     string `json:"POSTGRESS_DB_STR" env:"POSTGRESS_DB_STR"`
	PostgresTestDBStr string `json:"POSTGRESS_TEST_DB_STR" env:"POSTGRESS_DB_STR"`
	// Nexmo configs (mobile notification service)
	NexmoApiKey    string `json:"NEXMO_API_KEY"`
	NexmoApiSecret string `json:"NEXMO_API_SECRET"`
	// Secret token
	TokenSecret string `json:"TOKEN_SECRET"`
}

func NewConfig() *Config {
	return &Config{
		AppEnv:             getEnv("APP_ENV", "test"),
		HostName:           getEnv("SERVER_TOKEN", ""),
		ServerExternalPort: getEnv("SERVER_EXTERNAL_PORT", ""),
		PostgresDBStr:      getEnv("POSTGRESS_DB_STR", ""),
		PostgresTestDBStr:  getEnv("POSTGRESS_TEST_DB_STR", ""),
		NexmoApiKey:        getEnv("NEXMO_API_KEY", ""),
		NexmoApiSecret:     getEnv("NEXMO_API_SECRET", ""),
		TokenSecret:        getEnv("TOKEN_SECRET", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
