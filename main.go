package main

import (
	"authen-system/cmd"
	"authen-system/internal/config"
	"authen-system/pkg/configreader"
	"os"
)

func main() {
	r := cmd.CreateGin()
	appConfig, err := configreader.Init[config.App](os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(err)
	}
	cmd.RegisterHandler(r, appConfig)
	if err := r.Run(appConfig.Server.Port); err != nil {
		panic(err)
	}
}
