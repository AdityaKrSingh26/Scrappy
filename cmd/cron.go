package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/AdityaKrSingh26/Scrappy/config"
	"github.com/AdityaKrSingh26/Scrappy/internal/scraper"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Starts the scraping process on a scheduled basis",
	Long:  `This command sets up a cron job to execute the scraping process at regular intervals defined in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		startCron()
	},
}

// startCron initializes and starts the cron job based on the configuration
func startCron() {
	// Load configuration
	cfg, err := config.GetConfig() // Ensure it returns the config and an error
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	// Create a new cron instance
	c := cron.New()

	// Add the scraping job with the schedule from the config file
	_, err = c.AddFunc(cfg.Scraping.Schedule, func() {
		fmt.Printf("Starting scraping job at %s...\n", time.Now().Format(time.RFC3339))
		// Use the new FileName field for output
		if err := scraper.ScrapeInternships(cfg.Scraping.URL, cfg.Scraping.Format, cfg.Output.FilenamePattern); err != nil {
			log.Printf("Error during scraping: %s\n", err)
		} else {
			fmt.Printf("Scraping job completed at %s\n", time.Now().Format(time.RFC3339))
		}
	})
	if err != nil {
		log.Fatalf("Failed to add scraping job to cron: %s", err)
	}

	// Inform the user that the cron job has started
	fmt.Println("Cron job started with schedule:", cfg.Scraping.Schedule)
	c.Start()

	// Keep the cron job running indefinitely
	select {}
}

func init() {
	// Register the cron command with the root command
	rootCmd.AddCommand(cronCmd)
}
