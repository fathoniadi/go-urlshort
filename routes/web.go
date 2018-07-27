package routes

import (
	controllers "github.com/fathoniadi/go-urlshort/controllers"

	"github.com/go-chi/chi"
	"net/http"
)

func Web() http.Handler {
	web := chi.NewRouter()
	web.Get("/", controllers.HomeGet)
	web.Get("/{url-params}", controllers.HomeDecodeGet)
	return web
}
