package main

import (
	"fmt"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
)

func main() {
	envConfig := configs.GetCommonEnvConfigs()

	fmt.Println(envConfig.HostName)
}