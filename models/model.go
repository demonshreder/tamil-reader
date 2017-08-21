package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Book stores the metadata
type Book struct {
	ID      int
	Name    string
	Author  string
	Year    string
	Image   bool
	OCR     bool
	RawName string
	Path    string
	Pages   []Page
	Total   int
}

// Page stores every single page linked to its book
type Page struct {
	ID        int
	ImagePath string
	PageNo    int
	Complete  int
	BookID    uint
	Text      string
}

// User stores every single page linked to its book
type User struct {
	ID        int
	ImagePath string
	PageNo    int
	Complete  int
	BookID    uint
	Text      string
}

// ORM is the global DB
var ORM, err = gorm.Open("postgres", "user=tamil dbname=tamil_reader sslmode=disable")

func main() {
	ORM.CreateTable(&Page{}, &Book{})
	fmt.Println(err)
	fmt.Println("done")
	defer ORM.Close()

}
