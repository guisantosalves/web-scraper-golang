package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

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
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	// c.Visit will visit page and runs OnHTML function
	for i := 0; i < 108; i++ {
		fmt.Printf("scrapping page: %d\n", i)

		//converting i to string with Itoa
		c.Visit("https://internshala.com/internships/work-from-home-internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping complete")
	log.Println(c)
}
