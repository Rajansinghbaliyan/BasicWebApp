package pages

import (
	"fmt"
	"log"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the about page")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("No of bytes written %d", n)
}
