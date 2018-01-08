package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
)


func parseZippedLogFile(zipFile string) {
	// Open a zip archive for reading.
	print("\nParsing zipFile: " + zipFile + "\n")
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			print("\nerror!\n")
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}



func main() {
	zipFile := flag.String("zipFile", "", "Use --zipfile=/path/to/irccloudlogs.zip")
	flag.Parse()
	parseZippedLogFile(*zipFile)
}
