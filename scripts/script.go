package scripts

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/crypto/blake2b"

	"github.com/demonshreder/tamil-reader/models"
)

// CookieHMACSecret is what the name says
var CookieHMACSecret = []byte("YaathumOoreYaavarumKelirTheethumNandrumPirantharaVaaraSaathalum")

//CountPages counts the total number of pages in the PDF file using ghostscript via bash
func CountPages(path string) int {
	countCmd := exec.Command("bash", "-c", "gs -q -dNODISPLAY -c '("+path+") (r) file runpdfbegin pdfpagecount = quit';")
	countOut, err := countCmd.Output()
	fmt.Println(countOut, err)
	count, _ := strconv.Atoi(strings.TrimSpace(string(countOut)))
	return count
}

// PdfToImages converts name pdf to dest images via graphicsmagick using bash
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

//ImageToText converts the images to text and returns the OCR text using
// Tesseract v4.0 via bash
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

// HashMAC computes the blake2b HMAC based on input string,
// and [64]secretKey, if nil is randomly generated via crypt/rand
// HashMAC returns hex-encoded HMAC and secret.
func HashMAC(message string, secretKey []byte) (hash, secret string) {
	// HMAC = hash(key + hash(key + message))
	key := make([]byte, 64)
	if secretKey == nil {
		rand.Read(key)
	} else if len(secretKey) > 64 {
		key = secretKey[:64]
	} else {
		key = secretKey[:len(secretKey)]
	}
	mac, _ := blake2b.New512(key)
	kmac, _ := blake2b.New512(key)
	mac.Write([]byte(message))
	kmac.Write(mac.Sum(nil))
	return hex.EncodeToString(kmac.Sum(nil)), hex.EncodeToString(key)
}
