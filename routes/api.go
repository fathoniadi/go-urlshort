package routes

import (
	controllers_v1 "github.com/fathoniadi/go-urlshort/controllers/v1"

	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type Book struct {
	Title  string
	Author []string
}

func CustomNotFound(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(404)

	data := make(map[string]string)
	data["status"] = strconv.Itoa(404)
	data["data"] = "Not Found"

	json_data, _ := json.Marshal(data)

	res.Write(json_data)
	return
}

func Api() http.Handler {
	api := chi.NewRouter()
	api.Use(render.SetContentType(render.ContentTypeJSON))

	api.Route("/v1", func(api_v1 chi.Router) {
		var books []Book

		books = []Book{Book{"Hello Word", []string{"Fathoni Adi Kurniawan"}}}

		api_v1.Get("/ping", func(res http.ResponseWriter, req *http.Request) {
			json_data, _ := json.Marshal(books)
			res.Header().Set("Content-Type", "application/json")
			res.Write(json_data)
		})
	})

	api.NotFound(CustomNotFound)

	api.Get("/", controllers_v1.URLDecodeGet)
	return api
}
