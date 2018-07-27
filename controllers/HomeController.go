package controllers

import (
	pongo "github.com/flosch/pongo2"
	"github.com/go-chi/chi"
	"net/http"
)

type HomeController struct {
}

type Page struct {
	Title string
	Exist bool
}

func HomeGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome"))
	return
}

func HomeDecodeGet(res http.ResponseWriter, req *http.Request) {

	if URLParam := chi.URLParam(req, "url-params"); URLParam == "" {
		res.WriteHeader(400)
		res.Write([]byte("Gagal"))
		return
	}

	template := "views/home/index.html"

	user := []string{"Fathoni", "Bagong"}

	book := []Page{Page{Title: "Hello", Exist: true}, Page{Title: "C", Exist: false}}

	tmpl := pongo.Must(pongo.FromFile(template))

	err := tmpl.ExecuteWriter(pongo.Context{"title": "Index", "greating": "Hallo", "users": user, "books": book}, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	return
}
