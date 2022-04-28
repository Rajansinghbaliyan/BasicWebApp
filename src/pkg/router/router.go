package router

import (
	handlers2 "BasicWebApp/src/pkg/handlers"
	"net/http"
)

//Router handles the routing
func Router() {
	http.HandleFunc("/", handlers2.Repo.Home)
	http.HandleFunc("/about", handlers2.Repo.About)
}
