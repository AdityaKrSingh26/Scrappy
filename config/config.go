package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds the application configuration
type Config struct {
	Websites      []WebsiteConfig     `yaml:"websites"`
	Scraping      ScrapingConfig      `yaml:"scraping"`
	Output        OutputConfig        `yaml:"output"`
	Logging       LoggingConfig       `yaml:"logging"`
	Database      DatabaseConfig      `yaml:"database"`
	Notifications NotificationsConfig `yaml:"notifications"`
	Environment   string              `yaml:"environment"`
}

// WebsiteConfig defines the structure for each website configuration
type WebsiteConfig struct {
	URL       string            `yaml:"url"`
	Selectors map[string]string `yaml:"selectors"`
}

// ScrapingConfig holds scraping configuration
type ScrapingConfig struct {
	Schedule    string `yaml:"schedule"`
	Concurrency int    `yaml:"concurrency"`
	Timeout     int    `yaml:"timeout"`
}

// OutputConfig defines the output settings
type OutputConfig struct {
	Format          string `yaml:"format"`
	Directory       string `yaml:"directory"`
	FilenamePattern string `yaml:"filename_pattern"`
}

// LoggingConfig defines the logging settings
type LoggingConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// DatabaseConfig defines the database settings
type DatabaseConfig struct {
	Enabled          bool   `yaml:"enabled"`
	Type             string `yaml:"type"`
	ConnectionString string `yaml:"connection_string"`
	CollectionName   string `yaml:"collection_name"`
}

// NotificationsConfig defines the notification settings
type NotificationsConfig struct {
	Enabled bool        `yaml:"enabled"`
	Email   EmailConfig `yaml:"email"`
}

// EmailConfig holds email notification settings
type EmailConfig struct {
	SMTPServer     string `yaml:"smtp_server"`
	Port           int    `yaml:"port"`
	SenderEmail    string `yaml:"sender_email"`
	RecipientEmail string `yaml:"recipient_email"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
}

// Global config variable
var cfg *Config

// GetConfig returns the application configuration
func GetConfig() (*Config, error) {
	if cfg == nil {
		if err := LoadConfig(); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

// LoadConfig reads the config.yaml file and returns the parsed config
func LoadConfig() error {
	file, err := os.Open("config.yaml")
	if err != nil {
		return err // Return the error instead of logging and exiting
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return err // Return the error instead of logging and exiting
	}

	cfg = config // Set the global config variable
	return nil   // Return nil on success
}
