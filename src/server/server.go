package server

import (
	"log"
	"net/http"
)

func StartServer(portNumber string) {
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Started server on port: %v", portNumber)
}
