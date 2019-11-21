package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	re := flag.String("re", "", "regular expression for value in PDF page content")
	in := flag.String("in", "", "input PDF")
	out := flag.String("out", "", "directory for outputing PDFs")
	skip := flag.Bool("skip", false, "skip unmatched pages")
	debug := flag.Bool("debug", false, "output extracted text for each page")
	flag.Parse()

	//check -re
	if *re == "" {
		fmt.Println("-re must be set")
		return
	}
	matchRegexp, err := regexp.Compile(*re)
	if err != nil {
		fmt.Println("Invalid regexp:", err)
		return
	}

	//check -in
	if *in == "" {
		fmt.Println("Must specify -in file")
		return
	}

	//check -out
	if *out == "" {
		fmt.Println("Must specify -out directory")
		return
	}

	//open file
	f, err := os.Open(*in)
	if err != nil {
		log.Fatalln("Unable to open input PDF:", err)
	}

	//defer close file
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalln("Unable to close input PDF:", err)
		}
	}()

	//create PDF reader
	pdf, err := model.NewPdfReader(f)
	if err != nil {
		log.Fatalln("Unable to create PDF reader:", err)
	}

	var count int

	//loop through each page
	for i, p := range pdf.PageList {
		ex, err := extractor.New(p)
		if err != nil {
			log.Fatalf("Unable to create PDF page %d extractor: %v\n", i, err)
		}

		//extract text
		text, err := ex.ExtractText()
		if err != nil {
			log.Fatalf("Unable to extract PDF page %d text: %v\n", i, err)
		}

		if *debug {
			fmt.Printf("Page %d text:\n", i+1)
			fmt.Println(text)
		}

		//find regexp
		matches := matchRegexp.FindStringSubmatch(text)
		if len(matches) != 2 {
			if *skip {
				log.Println("Skipping page", i)
				continue
			} else {
				log.Fatalln("Unable to locate identifier in PDF text")
			}
		}

		username := matches[1]

		//create PDF writer for page
		w := model.NewPdfWriter()
		if err = w.AddPage(p); err != nil {
			log.Fatalln("Unable to add page to writer:", err)
		}

		fn := path.Join(*out, fmt.Sprintf("%s.pdf", username))

		//open output file
		wf, err := os.Create(fn)
		if err != nil {
			log.Fatalln("Unable to open new PDF file", fn, "for writing:", err)
		}

		log.Println("Writing", fn)

		//write PDF page
		if err = w.Write(wf); err != nil {
			log.Fatalf("Unable to write PDF file %s: %v\n", fn, err)
		}

		count = i + 1
	}

	log.Println("Wrote", count, "pages.")
}
