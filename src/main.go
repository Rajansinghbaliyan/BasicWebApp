package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello from go")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("No of bytes written: ", n)
	})

	_ = http.ListenAndServe("localhost:8080", nil)
}
