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

	res.WriteHeader(404)

	data := make(map[string]string)
	data["status"] = strconv.Itoa(404)
	data["data"] = "Not Found"

	json_data, _ := json.Marshal(data)

	res.Write(json_data)
	return
}

func ApiResponseType(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}

func Api() http.Handler {
	api := chi.NewRouter()
	api.Use(render.SetContentType(render.ContentTypeJSON))
	api.Use(ApiResponseType)

	api.Route("/v1", func(api_v1 chi.Router) {

		api_v1.Get("/ping", func(res http.ResponseWriter, req *http.Request) {
			data := make(map[string]string)
			data["status"] = strconv.Itoa(200)
			data["data"] = "pong"
			json_data, _ := json.Marshal(data)
			res.Write(json_data)
		})

		api_v1.Route("/books", func(api_v1_books chi.Router) {
			api_v1_books.Get("/", func(res http.ResponseWriter, req *http.Request) {
				var books []Book
				books = []Book{Book{Title: "Hello Word", Author: []string{"Fathoni Adi Kurniawan"}},
					Book{Title: "Sisop", Author: []string{"Thiar Hasbiya"}}}
				json_data, _ := json.Marshal(books)
				res.Write(json_data)
			})
		})
	})

	api.NotFound(CustomNotFound)

	api.Get("/", controllers_v1.URLDecodeGet)
	return api
}
