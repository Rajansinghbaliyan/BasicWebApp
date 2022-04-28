package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func renderPage(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./src/templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
