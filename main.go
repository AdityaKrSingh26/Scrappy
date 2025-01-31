package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AdityaKrSingh26/Scrappy/core"
	"github.com/AdityaKrSingh26/Scrappy/internal/platforms/glassdoor"
	"github.com/AdityaKrSingh26/Scrappy/internal/platforms/wellfound"
)

var (
	outputFormat = flag.String("format", "json", "Output format (json/csv)")
	outputFile   = flag.String("output", "internships", "Output file name")
)

func main() {
	flag.Parse()

	activeScrapers := []core.Scraper{
		&wellfound.WellfoundScraper{},
		&glassdoor.GlassdoorScraper{},
		// &linkedin.LinkedinScraper{},
	}

	var allJobs []core.Job

	for _, scraper := range activeScrapers {
		fmt.Printf("Scraping %s...\n", scraper.Name())
		jobs, err := scraper.Scrape()
		if err != nil {
			log.Printf("Error scraping %s: %v", scraper.Name(), err)
			continue
		}
		allJobs = append(allJobs, jobs...)
	}

	filename := *outputFile + "." + *outputFormat
	err := core.SaveToFile(*outputFormat, filename, allJobs)
	if err != nil {
		log.Fatalf("Error saving results: %v", err)
	}

	fmt.Printf("Saved %d jobs to %s\n", len(allJobs), filename)
}
