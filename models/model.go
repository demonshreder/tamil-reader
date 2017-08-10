package models

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
