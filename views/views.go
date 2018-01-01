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
	"time"

	"golang.org/x/crypto/bcrypt"

	"golang.org/x/crypto/blake2b"

	"github.com/demonshreder/tamil-reader/scripts"

	"github.com/demonshreder/tamil-reader/models"
)

var ORM = models.ORM
var workDir, _ = os.Getwd()
var templatePath = filepath.Join(workDir, "templates/")

// Home sweet home for the web app
func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/home.html"))
	cookie, err := r.Cookie("username")
	if err == nil {
		username := strings.Split(cookie.Value, ":")
		p := map[string][]string{
			"books":    []string{"cooool", "kekek"},
			"username": []string{username[0]},
		}
		t.Execute(w, p)
	} else {
		//p := {}
		p := map[string][]string{
			"books":    []string{"cooool", "kekek"},
			"username": []string{},
		}
		t.Execute(w, p)
	}
}

// UserLogin checks username and passwords and logs the user in
func UserLogin(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/userPage.html"))
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		user := models.User{Username: username}
		if username == "" {
			http.Redirect(w, r, "/user/", 302)
		} else if password == "" {
			http.Redirect(w, r, "/user/", 302)
		} else {
			models.ORM.Where("Username = ?", username).First(&user)
			if models.ORM.NewRecord(&user) {
				http.Redirect(w, r, "/user/", 302)
			} else {
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
				if err != nil {
					http.Redirect(w, r, "/user/", 302)
				}
				mc, _ := scripts.HashMAC(username, scripts.CookieHMACSecret)
				cookieExpiryTime := time.Now().Add(1209600 * time.Second)
				// fmt.Println(fmt.Sprint(cookieExpiry))
				cookieExpiryByte, _ := cookieExpiryTime.MarshalText()
				cookie := http.Cookie{
					Name:     "username",
					Value:    username + ":" + mc,
					Path:     "/",
					Expires:  cookieExpiryTime,
					HttpOnly: true,
				}
				session := models.Session{
					UserID: uint(user.ID),
					Hash:   string(mc),
					End:    string(cookieExpiryByte),
					Expiry: cookieExpiryTime,
				}
				models.ORM.Create(&session)
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/", 302)
			}
		}
		// fmt.Println(http.)
	}
	// http.Redirect(w, r, "/", 302)
}

// UserPage just renders the template to show user login and register page
func UserPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/userLogReg.html"))
	t.Execute(w, nil)
}

// UserRegister accepts POST data and registers the user
func UserRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		username := r.FormValue("username")
		pass1 := r.FormValue("password")
		pass2 := r.FormValue("password2")
		email := r.FormValue("email")

		user := models.User{}

		if username == "" {
			http.Redirect(w, r, "/user/", 302)
		} else if pass1 == "" || pass2 == "" {
			http.Redirect(w, r, "/user/", 302)
		} else if pass1 != pass2 {
			http.Redirect(w, r, "/user/", 302)
		} else if email == "" {
			http.Redirect(w, r, "/user/", 302)
		} else {
			models.ORM.Where("Username = ?", username).First(&user)
			if !models.ORM.NewRecord(&user) {
				http.Redirect(w, r, "/user/", 302)
			} else {
				hash, _ := bcrypt.GenerateFromPassword([]byte(pass1), 7)
				user = models.User{
					Username: username,
					Password: string(hash),
					Email:    email,
				}
				fmt.Println("success")
				models.ORM.NewRecord(&user)
				models.ORM.Save(&user)
			}
		}
	}
	http.Redirect(w, r, "/", 302)

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
	cookie, _ := r.Cookie("username")
	username := strings.Split(cookie.Value, ":")
	p := map[string][]string{
		"username": []string{username[0]},
	}

	t.Execute(w, p)

}

func PageEdit(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(templatePath+"/base.html", templatePath+"/edit.html"))
	page := models.Page{}
	ORM.Where("Complete = ?", 0).First(&page)
	book := models.Book{ID: int(page.BookID)}
	ORM.Find(&book)
	cookie, _ := r.Cookie("username")
	username := strings.Split(cookie.Value, ":")
	p := map[string]string{
		"imageURL":  strings.Replace(page.ImagePath, workDir, "", -1),
		"pageText":  page.Text,
		"pageID":    strconv.Itoa(page.ID),
		"bookName":  book.Name,
		"current":   strconv.Itoa(page.PageNo),
		"totalPage": strconv.Itoa(book.Total),
		"username":  username[0],
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
	http.Redirect(w, r, "/page/edit/", 302)
}
