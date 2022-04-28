package router

import (
	"BasicWebApp/src/pages"
	"net/http"
)

func Router() {
	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/about", pages.About)
}
