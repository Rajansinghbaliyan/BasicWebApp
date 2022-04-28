package main

import (
	"BasicWebApp/src/router"
	"BasicWebApp/src/server"
)

const portNumber = ":8080"

func main() {
	router.Router()
	server.StartServer(portNumber)
}
