package main

import (
	"fmt"
	"net/http"

	"github.com/demonshreder/tamil-reader/models"
	"github.com/demonshreder/tamil-reader/routers"
)

func main() {
	models.ORM.AutoMigrate(&models.User{}, &models.Session{}, &models.Page{}, &models.Book{})
	// models.ORM.CreateTable(&models.Session{})
	// fmt.Println(time.MarshalText())
	fmt.Println("Tamil reader listening on http://127.0.0.1:4000")
	http.ListenAndServe(":4000", routers.Router())
	defer models.ORM.Close()

}
