package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AdityaKrSingh26/Scrappy/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Websites      []WebsiteConfig    `mapstructure:"websites"`
	Scraping      ScrapingConfig     `mapstructure:"scraping"`
	Output        OutputConfig       `mapstructure:"output"`
	Logging       LoggingConfig      `mapstructure:"logging"`
	Database      DatabaseConfig     `mapstructure:"database"`
	Notifications NotificationConfig `mapstructure:"notifications"`
	Environment   string             `mapstructure:"environment"`
}

// Load configuration from the config.yaml file
func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into config struct: %s", err)
	}
}

func main() {
	// Load the configuration
	loadConfig()

	// Set up the root command
	var rootCmd = &cobra.Command{
		Use:   "scrappy",
		Short: "Scrappy is a command-line tool for scraping internship opportunities.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Starting the scraping process manually...")
			return cmd.RunScrapeCommand()
		},
	}

	// Add the cron command to root
	rootCmd.AddCommand(cmd.CronCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		os.Exit(1)
	}
}
