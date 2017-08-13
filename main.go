package main

import (
	"fmt"
	"net/http"

	"github.com/demonshreder/tamil-reader/models"
	"github.com/demonshreder/tamil-reader/routers"
)

func main() {

	fmt.Println("Tamil reader listening on http://127.0.0.1:4000")
	http.ListenAndServe(":4000", routers.Router())
	defer models.ORM.Close()

}
