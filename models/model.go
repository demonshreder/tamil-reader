package models

import (
	"fmt"
	"time"

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
	Total   int
	Pages   []Page
	UserID  uint
}

// Page stores every single page linked to its book
type Page struct {
	ID        int
	ImagePath string
	PageNo    int
	Complete  int
	Text      string
	BookID    uint
	UserID    uint
}

// User stores every single page linked to its book
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Books    []Book
	Pages    []Page
	Sessions []Session
}

// Session stores session for every user
type Session struct {
	ID     int
	UserID uint
	Hash   string
	End    string
	Expiry time.Time
}

// ORM is the global DB
var ORM, err = gorm.Open("postgres", "user=tamil dbname=tamil_reader sslmode=disable password=tamilrocks")

func main() {
	ORM.CreateTable(&Page{}, &Book{}, &User{})
	fmt.Println(err)
	fmt.Println("done")
	defer ORM.Close()

}
