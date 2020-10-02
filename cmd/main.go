package main

import (
	"flag"
	"fmt"
	"github.com/code7unner/leadersofdigital2020-backend/configs"
)

var (
	envPathFlag string
)

func init() {
	flag.StringVar(&envPathFlag, "env", ".env", "env file path")
	flag.Parse()
}

func main() {
	envConfig := configs.GetCommonEnvConfigs()

	fmt.Println(envConfig.HostName)
}
