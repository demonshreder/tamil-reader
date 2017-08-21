package scripts

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/demonshreder/tamil-reader/models"
)

//CountPages counts the total number of pages in the PDF file
func CountPages(path string) int {
	countCmd := exec.Command("bash", "-c", "gs -q -dNODISPLAY -c '("+path+") (r) file runpdfbegin pdfpagecount = quit';")
	countOut, err := countCmd.Output()
	fmt.Println(countOut, err)
	count, _ := strconv.Atoi(strings.TrimSpace(string(countOut)))
	return count
}

// PdfToImages converts name pdf to dest images
func PdfToImages(book models.Book) {
	fmt.Println("fuccking started")
	// gm convert -verbose -trim -density 300 akan_aanuuru.pdf +adjoin akan/akan_aanuuru-%03d.jpg
	for i := 1; i < book.Total; i++ {
		fullPath := strings.Trim(book.Path, ".pdf") + "-" + strconv.Itoa(i) + ".jpg"
		cmd := exec.Command("bash", "-c", "gm convert -verbose -trim -density 300 "+book.Path+"["+strconv.Itoa(i)+"] "+fullPath)
		fmt.Println(cmd)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		text := ImageToText(fullPath)
		page := models.Page{
			ImagePath: fullPath,
			PageNo:    i,
			Complete:  0,
			BookID:    uint(book.ID),
			Text:      text,
		}

		models.ORM.NewRecord(&page)
		models.ORM.Create(&page)
		fmt.Println("inserted", i)
	}
}

//ImageToText converts the images to text and returns the OCR text
func ImageToText(path string) string {
	// tesseract akan/akan_aanuuru-000.jpg  stdout --oem 1 -l tam
	// fmt.Println("tooser action")
	// cmd := exec.Command("tesseract", path, "stdout", "--oem", "1", "-l", "tam")
	cmd := exec.Command("bash", "-c", "tesseract "+path+" stdout --oem 1 -l tam")
	fmt.Println(cmd)
	out, err := cmd.Output()
	// fmt.Println(out)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("tesseracted")
	return string(out)
}
