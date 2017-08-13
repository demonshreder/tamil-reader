package views

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/demonshreder/tamil-reader/scripts"

	"github.com/demonshreder/tamil-reader/models"
)

var ORM = models.ORM

func Home(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	templatePath := filepath.Join(workDir, "templates/")
	fmt.Println(templatePath)
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/home.html"))
	// fmt.Println(err.Error())
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
func New(w http.ResponseWriter, r *http.Request) {

	workDir, _ := os.Getwd()
	templatePath := filepath.Join(workDir, "templates/")
	fmt.Println(templatePath)
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/new.html"))

	if r.Method == "POST" {
		book, _, _ := r.FormFile("book")

		bookName := r.FormValue("book-name")
		bookPath := workDir + "/raw/" + bookName
		os.Mkdir(bookPath, 0755)
		fmt.Println(bookPath)
		bookPath = bookPath + "/" + bookName + ".pdf"
		fmt.Println(bookPath)
		// book.Close()
		pdf, err := os.OpenFile(bookPath, os.O_WRONLY|os.O_CREATE, 0755)
		fmt.Println(err)
		io.Copy(pdf, book)
		bookR := models.Book{
			Name:    bookName,
			Author:  r.FormValue("author"),
			Image:   false,
			Path:    bookPath,
			OCR:     false,
			RawName: bookName + ".pdf",
			Total:   scripts.CountPages(bookPath),
			Year:    "2017",
		}
		ORM.NewRecord(bookR)
		ORM.Create(&bookR)
		defer book.Close()

		defer pdf.Close()
	}
	t.Execute(w, nil)
	// go scripts.PdfToImages(bookPath)

}
