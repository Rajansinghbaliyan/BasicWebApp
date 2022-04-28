package main

import (
	"BasicWebApp/src/pkg/config"
	"BasicWebApp/src/pkg/handlers"
	"BasicWebApp/src/pkg/router"
)

const portNumber = ":8080"

func main() {
	var appConfig config.AppConfig

	appConfig.UseCache = true

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	router.GetConfigForRouter(&appConfig)
	//router.Router()
	StartServer(portNumber, &appConfig)
}
