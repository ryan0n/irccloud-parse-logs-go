package main

import (
	"archive/zip"
	"fmt"
	"log"
	"flag"
	"bufio"
	"strings"
)

func parseZippedLogFile(zipFile string, searchPhrase string) {
	print("\nParsing zipFile: " + zipFile + "\n")

	// Open the zip file for reading
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Loop through the files in the archive
	for _, f := range r.File {
		rc, _ := f.Open()
		fileScanner := bufio.NewScanner(rc)
		for fileScanner.Scan() {
			// Print the current line out if it matches the searchPhrase
			if strings.Contains(fileScanner.Text(), searchPhrase) {
				fmt.Println(fileScanner.Text())
			}
		}
		rc.Close()
	}
}

func main() {
	zipFile := flag.String("zipFile", "", "Use --zipfile=/path/to/irccloudlogs.zip --searchPhrase=SomeSearchPhrase")
	searchPhrase := flag.String("searchPhrase", "", "Use --zipfile=/path/to/irccloudlogs.zip --searchPhrase=SomeSearchPhrase")
	flag.Parse()
	parseZippedLogFile(*zipFile, *searchPhrase)
}
