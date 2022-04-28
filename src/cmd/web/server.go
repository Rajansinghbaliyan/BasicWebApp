package main

import (
	"BasicWebApp/src/pkg/config"
	"log"
	"net/http"
	"time"
)

//StartServer starts and listen to the port
func StartServer(portNumber string, a *config.AppConfig) {
	srv := &http.Server{
		Handler: Router(a),
		Addr:    portNumber,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
