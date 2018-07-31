package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/hero0926/docx"
)

func main() {
	// Read from docx file
	r, err := docx.ReadDocxFile("./test.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
	if err != nil {
		panic(err)
	}

	docx1 := r.Editable()
	example := docx1.GetText()

	// Make a Regex to say we only want
	reg, err := regexp.Compile("<[^>]*>")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(example, "")

	fmt.Printf(processedString)
	r.Close()
}
