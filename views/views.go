package views

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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
	// fmt.Println(err.Error())
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
func NewBook(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	templatePath := filepath.Join(workDir, "templates/")
	fmt.Println(templatePath)
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/new.html"))
	book, _, _ := r.FormFile("book")
	defer book.Close()
	bookPath := workDir + "/pdf/" + r.FormValue("book-name") + ".pdf"
	// book.Close()
	pdf, _ := os.OpenFile(bookPath, os.O_WRONLY|os.O_CREATE, 0666)
	defer pdf.Close()
	io.Copy(pdf, book)
	// fmt.Println(err.Error())
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
