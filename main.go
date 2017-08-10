package main

import (
	"net/http"

	_ "github.com/demonshreder/tamil-reader/models"
	"github.com/demonshreder/tamil-reader/routers"
)

func main() {
	// r := chi.NewRouter()
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	http.ListenAndServe(":4000", routers.Router())
}
