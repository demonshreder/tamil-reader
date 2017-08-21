package views

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/crypto/blake2b"

	"github.com/demonshreder/tamil-reader/scripts"

	"github.com/demonshreder/tamil-reader/models"
)

var ORM = models.ORM
var workDir, _ = os.Getwd()
var templatePath = filepath.Join(workDir, "templates/")

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/home.html"))
	t.Execute(w, nil)

}
func New(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/new.html"))
	if r.Method == "POST" {
		book, bookH, _ := r.FormFile("book")
		defer book.Close()
		bookName := r.FormValue("book-name")
		bookByte, _ := ioutil.ReadAll(book)
		book.Seek(0, 0)
		blakeSum := blake2b.Sum512(bookByte)
		blakeStr := base64.StdEncoding.EncodeToString([]byte(blakeSum[:6]))
		fmt.Println(blakeStr)
		bookPath := workDir + "/raw/" + blakeStr
		os.Mkdir(bookPath, 0755)
		bookPath = bookPath + "/" + bookH.Filename
		pdf, _ := os.OpenFile(bookPath, os.O_WRONLY|os.O_CREATE, 0755)
		defer pdf.Close()
		io.Copy(pdf, book)
		bookR := models.Book{
			Name:    bookName,
			Author:  r.FormValue("author"),
			Image:   false,
			Path:    bookPath,
			OCR:     false,
			RawName: bookH.Filename,
			Total:   scripts.CountPages(bookPath),
			Year:    r.FormValue("year"),
		}
		ORM.NewRecord(bookR)
		ORM.Create(&bookR)
		go scripts.PdfToImages(bookR)
	}
	t.Execute(w, nil)

}

func PageEdit(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/edit.html"))
	page := models.Page{}
	ORM.Where("Complete = ?", 0).First(&page)
	book := models.Book{ID: int(page.BookID)}
	ORM.Find(&book)
	p := map[string]string{
		"imageURL":  strings.Replace(page.ImagePath, workDir, "", -1),
		"pageText":  page.Text,
		"pageID":    strconv.Itoa(page.ID),
		"bookName":  book.Name,
		"current":   strconv.Itoa(page.PageNo),
		"totalPage": strconv.Itoa(book.Total),
	}
	err := t.Execute(w, p)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
func PageSave(w http.ResponseWriter, r *http.Request) {
	pageID, _ := strconv.Atoi(r.FormValue("pageID"))
	pageComp := 0
	if r.FormValue("pageComplete") == "true" {
		pageComp = 1
	}
	page := models.Page{
		ID:       pageID,
		Text:     r.FormValue("pageText"),
		Complete: pageComp,
	}
	ORM.Model(&page).Update(page)
	http.Redirect(w, r, "/page/edit", 302)
}
