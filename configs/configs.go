package configs

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type CommonEnvConfigs struct {
	// Service configs
	AppEnv             string `json:"APP_ENV" env:"APP_ENV"`
	GinMode            string `json:"GIN_MODE" env:"GIN_MODE"`
	HostName           string `json:"HOSTNAME" env:"HOSTNAME"`
	ServerExternalPort string `json:"SERVER_EXTERNAL_PORT" env:"SERVER_EXTERNAL_PORT" envDefault:"8080"`
	// Postgres configs
	PostgresDBStr  string `json:"POSTGRESS_DB_STR" env:"POSTGRESS_DB_STR" envDefault:"postgres://postgres:postgres@localhost?sslmode=disable"`
	PostgresUser   string `json:"POSTGRESS_USER" env:"POSTGRESS_USER"`
	PostgresPass   string `json:"POSTGRESS_PASS" env:"POSTGRESS_PASS"`
	PostgresDBName string `json:"POSTGRESS_DB_NAME" env:"POSTGRESS_DB_NAME"`
	PostgresHost   string `json:"POSTGRESS_HOST" env:"POSTGRESS_HOST"`
	PostgresPort   string `json:"POSTGRESS_PORT" env:"POSTGRESS_PORT"`
}

func GetCommonEnvConfigs() CommonEnvConfigs {
	envConfig := CommonEnvConfigs{}
	err := env.Parse(&envConfig)
	if err != nil {
		log.Panicf("Error parse env config:%s", err)
		return envConfig
	}
	return envConfig
}

func (e CommonEnvConfigs) TestMode() bool {
	return e.AppEnv != "production"
}
