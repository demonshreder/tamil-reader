package pdfToImages

import (
	"fmt"
	"log"
	"os/exec"
)

// PdfToImages converst name pdf to dest images
func PdfToImages(name, dest string) {
	cmd := exec.Command("gm", "convert", "-verbose", "-trim", "-density", "300", name, "+adjoin", dest)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func main() {
	fmt.Println("started")
	PdfToImages("")
}
