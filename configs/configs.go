package configs

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type CommonEnvConfigs struct {
	AppEnv   string `json:"APP_ENV" env:"APP_ENV" envDefault:"test"`
	GinMode  string `json:"GIN_MODE" env:"GIN_MODE" envDefault:"test"`
	HostName string `json:"HOSTNAME" env:"HOSTNAME" envDefault:"TESTHOSTNAME"`
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
