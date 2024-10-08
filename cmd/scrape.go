package cmd

import (
	"log"

	"github.com/AdityaKrSingh26/Scrappy/config"
	"github.com/AdityaKrSingh26/Scrappy/internal/scraper"
)

// RunScrapeCommand runs the scraping process using the configuration
func RunScrapeCommand() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err // Return the error if loading configuration fails
	}

	// Loop through the websites and scrape each one
	for _, website := range cfg.Websites {
		log.Printf("Scraping website: %s\n", website.URL)
		err := scraper.ScrapeInternships(website.URL, cfg.Scraping.Format, cfg.Output.FilenamePattern)
		if err != nil {
			return err // Return the error if scraping fails
		}
		log.Printf("Successfully scraped: %s\n", website.URL)
	}

	return nil // Return nil on success
}
