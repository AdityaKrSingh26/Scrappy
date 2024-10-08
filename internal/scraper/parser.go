package scraper

import (
	"github.com/gocolly/colly/v2"
)

var internships [][]string // Global slice to hold internship data

// ParseInternshipData extracts internship data from the HTML content
func ParseInternshipData(c *colly.Collector) {
	c.OnHTML("div.internship", func(e *colly.HTMLElement) {
		title := e.ChildText("h2.title")
		company := e.ChildText("p.company")
		location := e.ChildText("p.location")

		// Store the extracted data
		internship := []string{title, company, location}
		internships = append(internships, internship) // Store data in the slice
	})
}

// Call this function after scraping is done to save the data
func SaveData(format string, fileName string) error {
	if format == "csv" {
		return WriteToCSV(fileName, internships)
	} else if format == "json" {
		return WriteToJSON(fileName, internships)
	}
	return nil
}
