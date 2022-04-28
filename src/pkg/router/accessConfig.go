package router

import "BasicWebApp/src/pkg/config"

var appConfig *config.AppConfig

// GetConfigForRouter gets the config object form the main
func GetConfigForRouter(a *config.AppConfig) {
	appConfig = a
}
