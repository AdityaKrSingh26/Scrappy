package scraper

import (
	"log"

	"github.com/gocolly/colly/v2"
)

// ScrapeInternships starts the scraping process for internship opportunities
func ScrapeInternships(url string, format string, fileName string) error {
	c := colly.NewCollector()

	// Set up parsing logic
	ParseInternshipData(c)

	// Define the scraping process
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request failed: %s, error: %v", r.Request.URL, err)
	})

	// Start scraping the URL
	err := c.Visit(url)
	if err != nil {
		return err // Return the error to the caller
	}

	// After scraping, save the data
	return SaveData(format, fileName)
}
