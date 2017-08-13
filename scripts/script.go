package scripts

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/demonshreder/tamil-reader/models"
)

//CountPages counts the total number of pages in the PDF file
func CountPages(path string) int {
	countCmd := exec.Command("bash", "-c", "gs -q -dNODISPLAY -c '("+path+") (r) file runpdfbegin pdfpagecount = quit';")
	countOut, _ := countCmd.Output()
	count, _ := strconv.Atoi(strings.TrimSpace(string(countOut)))
	return count
}

// PdfToImages converts name pdf to dest images
func PdfToImages(src, dest string, book *models.Book) {
	// gs -q -dNODISPLAY -c "(akan_aanuuru.pdf) (r) file runpdfbegin pdfpagecount = quit";
	// gm convert -verbose -trim -density 300 akan_aanuuru.pdf +adjoin akan/akan_aanuuru-%03d.jpg
	cmd := exec.Command("gm", "convert", "-verbose", "-trim", "-density", "300", src, "+adjoin", dest)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	// log.Printf("Command finished with error: %v", err)
	ImageToText()
}

//ImageToText converts the images to text and populates the database with it
func ImageToText() {

}
