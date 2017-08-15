package main

import (
	"fmt"
	"net/http"

	"github.com/demonshreder/tamil-reader/models"
	"github.com/demonshreder/tamil-reader/routers"
)

func main() {
	// fmt.Println("starting tesseraction")
	// path := "/home/demonshreder/Projects/tamilreader/tamil/tamilvu/literature/akan/akan_aanuuru-000.jpg"
	// fmt.Println(scripts.ImageToText(path))
	fmt.Println("Tamil reader listening on http://127.0.0.1:4000")
	http.ListenAndServe(":4000", routers.Router())
	defer models.ORM.Close()

}
