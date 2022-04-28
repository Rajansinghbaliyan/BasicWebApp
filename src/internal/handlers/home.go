package handlers

import (
	"log"
	"net/http"
)

//Home is about the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "This is the home page")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Printf("No of bytes written %d", n)

	log.Println(m.App.UseCache)
	renderPage(w, "home.page.tmpl")
}
