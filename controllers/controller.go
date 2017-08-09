package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
<<<<<<< HEAD
	"github.com/astaxie/beego/orm"

	"github.com/demonshreder/tamil-reader/models"
=======
>>>>>>> cb86ddb3ee3187f8690a5ebef792c882c2310ad4
)

type HomeController struct {
	beego.Controller
}

type NewController struct {
	beego.Controller
}

type NewBook struct {
	beego.Controller
}

func (c *HomeController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.html"
}

func (c *NewController) Get() {
	c.TplName = "new.html"
}
func (c *NewBook) Post() {
	bookName := c.GetString("book-name")
	author := c.GetString("author")
<<<<<<< HEAD
	// redir := c.URLFor("HomeController.Get")
	// bookdata, _, _ := c.GetFile("book")
	// bookdata
	o := orm.NewOrm()
	book := new(models.Book)
	book.Name = bookName
	book.Author = author
	page := new(models.Page)
	fmt.Println(o.Insert(book))
	fmt.Println("reader", o.Read(book))
	page.Book = book
	fmt.Println(o.Insert(page))
	c.SaveToFile("book", "pdf/"+bookName+".pdf")
	// fmt.Println(bookName, author, redir)
=======
	redir := c.URLFor("HomeController.Get")
	// bookdata, _, _ := c.GetFile("book")
	c.SaveToFile("book", "/home/demonshreder/super.pdf")
	fmt.Println(bookName, author, redir)
>>>>>>> cb86ddb3ee3187f8690a5ebef792c882c2310ad4
	c.Redirect("/", 200)
}
