package main

import (
	"github.com/fathoniadi/go-urlshort/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	app := chi.NewRouter()

	app.Use(middleware.RequestID)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.URLFormat)
	app.Use(middleware.RealIP)

	api := routes.Api()
	web := routes.Web()

	app.Mount("/api", api)
	app.Mount("/", web)

	http.ListenAndServe(":3000", app)
}
