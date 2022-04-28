package main

import (
	"BasicWebApp/src/pkg/config"
	"BasicWebApp/src/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
