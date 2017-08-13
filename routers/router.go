package routers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/demonshreder/tamil-reader/views"
	"github.com/go-chi/chi"
)

func Router() chi.Router {

	r := chi.NewRouter()
	r.Get("/", views.Home)
	r.Get("/new/", views.New)
	r.Post("/new/", views.New)
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))
	return r
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

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
