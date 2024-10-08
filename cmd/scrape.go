package cmd

import (
	"fmt"
	"log"

	"github.com/AdityaKrSingh26/Scrappy/internal/scraper"
	"github.com/spf13/cobra"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape internship opportunities from various websites",
	Long: `The scrape command fetches internship opportunities from a predefined set of
websites and stores the results in specified output formats (CSV/JSON).`,
	Run: func(cmd *cobra.Command, args []string) {
		// Extract any flags or arguments if needed
		outputFormat, _ := cmd.Flags().GetString("output")
		if outputFormat == "" {
			outputFormat = "json" // default format
		}

		fmt.Printf("Starting the scraping process. Output format: %s\n", outputFormat)

		// Call the scraping function from the scraper package
		err := scraper.ScrapeInternships(outputFormat)
		if err != nil {
			log.Fatalf("Failed to complete scraping: %v", err)
		}

		fmt.Println("Scraping completed successfully!")
	},
}

func init() {
	// Add the scrape command to the root command
	rootCmd.AddCommand(scrapeCmd)

	// Add any flags specific to the scrape command
	scrapeCmd.Flags().StringP("output", "o", "json", "Output format for the scraped data (options: json, csv)")
}
