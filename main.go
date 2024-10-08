package main

import (
	"fmt"
	"log"
	"os"

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

// WebsiteConfig holds the configuration for each website
type WebsiteConfig struct {
	URL       string            `mapstructure:"url"`
	Selectors map[string]string `mapstructure:"selectors"`
}

// ScrapingConfig holds scraping-related settings
type ScrapingConfig struct {
	Schedule    string `mapstructure:"schedule"`
	Concurrency int    `mapstructure:"concurrency"`
	Timeout     int    `mapstructure:"timeout"`
}

// OutputConfig holds output-related settings
type OutputConfig struct {
	Format          string `mapstructure:"format"`
	Directory       string `mapstructure:"directory"`
	FilenamePattern string `mapstructure:"filename_pattern"`
}

// LoggingConfig holds logging-related settings
type LoggingConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

// DatabaseConfig holds database-related settings
type DatabaseConfig struct {
	Enabled          bool   `mapstructure:"enabled"`
	Type             string `mapstructure:"type"`
	ConnectionString string `mapstructure:"connection_string"`
	CollectionName   string `mapstructure:"collection_name"`
}

// NotificationConfig holds notification-related settings
type NotificationConfig struct {
	Enabled bool        `mapstructure:"enabled"`
	Email   EmailConfig `mapstructure:"email"`
}

// EmailConfig holds email notification settings
type EmailConfig struct {
	SMTPServer     string `mapstructure:"smtp_server"`
	Port           int    `mapstructure:"port"`
	SenderEmail    string `mapstructure:"sender_email"`
	RecipientEmail string `mapstructure:"recipient_email"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
}

var config Config

// loadConfig loads the configuration from config.yaml
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
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting the scraping process...")
			// Here, you would start the scraping logic
			startScraping()
		},
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		os.Exit(1)
	}
}

// startScraping implements the scraping logic
func startScraping() {
	for _, website := range config.Websites {
		fmt.Printf("Scraping website: %s\n", website.URL)
		// Here you would implement the scraping logic using Colly
	}
}
