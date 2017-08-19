package views

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/demonshreder/tamil-reader/scripts"

	"github.com/demonshreder/tamil-reader/models"
)

var ORM = models.ORM
var workDir, _ = os.Getwd()
var templatePath = filepath.Join(workDir, "templates/")

func Home(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	templatePath := filepath.Join(workDir, "templates/")
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/home.html"))
	page := models.Page{}
	ORM.First(&page)
	pageID := strconv.Itoa(page.ID)
	p := map[string]string{
		"imageURL": strings.Replace(page.ImagePath, workDir, "", -1),
		"pageText": page.Text,
		"pageID":   pageID,
	}
	err := t.Execute(w, p)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
func New(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/new.html"))
	if r.Method == "POST" {
		book, bookH, _ := r.FormFile("book")
		bookName := r.FormValue("book-name")
		bookPath := workDir + "/raw/" + bookName
		os.Mkdir(bookPath, 0755)
		bookPath = bookPath + "/" + bookName + ".pdf"
		pdf, _ := os.OpenFile(bookPath, os.O_WRONLY|os.O_CREATE, 0755)
		io.Copy(pdf, book)
		bookR := models.Book{
			Name:    bookName,
			Author:  r.FormValue("author"),
			Image:   false,
			Path:    bookPath,
			OCR:     false,
			RawName: bookH.Filename,
			Total:   scripts.CountPages(bookPath),
			Year:    "2017",
		}
		ORM.NewRecord(bookR)
		ORM.Create(&bookR)
		defer book.Close()
		defer pdf.Close()
		go scripts.PdfToImages(bookR)
	}
	t.Execute(w, nil)

}

func SavePage(w http.ResponseWriter, r *http.Request) {
	pageID, _ := strconv.Atoi(r.FormValue("pageID"))
	pageComp := 0
	page := models.Page{
		ID:   pageID,
		Text: r.FormValue("pageText"),
	}
	if r.FormValue("pageComplete") == "true" {
		pageComp = 1
	}
	fmt.Println(pageComp)

	ORM.Model(&page).Update(page)
	http.Redirect(w, r, "/", 302)
}
