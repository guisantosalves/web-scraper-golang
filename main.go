package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"

	file, err := os.Create(fName)
	if err != nil {
		// using %q to set quotes ""
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	// run after all executuin in this function
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// preparing domain to collect data
	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{})
	})
}
