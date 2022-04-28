package main

import (
	"BasicWebApp/src/internal/handlers"
	"BasicWebApp/src/internal/helper"
	"BasicWebApp/src/pkg/config"
	"log"
	"os"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog

	appConfig.UseCache = true

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)
	helper.NewHelper(&appConfig)

	//router.GetConfigForRouter(&appConfig)
	//router.Router()
	StartServer(portNumber, &appConfig)
}
