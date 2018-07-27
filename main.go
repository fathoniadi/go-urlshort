package main

import (
	"github.com/fathoniadi/go-urlshort/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(app, "/static", http.Dir(filesDir))

	http.ListenAndServe(":3000", app)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if info, err := os.Stat(req.URL.Path[1:]); err == nil && info.IsDir() {
			res.WriteHeader(403)
			res.Write([]byte("403 Forbidden Access"))
			return
		}

		fs.ServeHTTP(res, req)
	}))
}
