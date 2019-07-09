package main

import (
	"github.com/GuilhermeFirmiano/mock-google-maps/api/server"
	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/settings"
)

func main() {
	appSettings := settings.FromDotEnv()

	api := server.Configure(appSettings)

	api.RegisterDependencies()
	api.RegisterControllers()
	api.Run()
}
