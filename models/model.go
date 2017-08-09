package models

import (
	"github.com/astaxie/beego/orm"
)

// Book stores the metadata
type Book struct {
	ID     int
	Name   string
	Author string
	Year   string
	Image  bool
	Path   string
	// Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

// Page stores every single page linked to its book
type Page struct {
	ID        int
	ImagePath string
	PageNo    int
	Complete  int
	Book      *Book `orm:"rel(fk)"` // Reverse relationship (optional)
}

func init() {
	// o := orm.NewOrm()
	// o.Using("default")
	// Need to register model in init
	orm.RegisterModel(new(Book), new(Page))
}
