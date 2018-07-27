package routes

import (
	controllers_v1 "github.com/fathoniadi/go-urlshort/controllers/v1"

	"github.com/go-chi/chi"
	"net/http"
)

func Api() http.Handler {
	api := chi.NewRouter()
	api.Get("/", controllers_v1.URLDecodeGet)
	return api
}
