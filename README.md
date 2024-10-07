# Scrappy

Scrappy is a command-line application for scraping internship opportunities from various websites using Golang. It offers fast execution with Cobra, efficient data extraction with Colly, and automated scheduling using Go-cron, storing results in CSV/JSON formats.

## Project Structure

The folder structure is organized as follows:

```
Scrappy/
│
├── cmd/                       # Command definitions for the CLI application
│   ├── root.go                # Main entry point for the CLI commands
│   ├── scrape.go              # Command for initiating the scraping process
│   └── cron.go                # Command for setting up and managing cron jobs
│
├── internal/                  # Internal packages for modular functionality
│   ├── scraper/               # Scraping logic and parsing mechanisms
│   │   ├── scraper.go         # Core scraping logic using Colly
│   │   ├── parser.go          # Functions for parsing scraped data
│   │   └── helpers.go         # Utility functions used in scraping
│   │
│   ├── storage/               # Storage mechanisms and data formatting
│   │   ├── storage.go         # Logic for saving data to CSV/JSON files
│   │   └── formatters.go      # Formatting functions for output data
│   │
│   └── scheduler/             # Scheduler for automating scraping tasks
│       └── cron.go            # Functions for setting up and managing Go-cron jobs
│
├── config/                    # Configuration files for the application
│   └── config.yaml            # Application configurations (e.g., scraping frequency, URLs)
│
├── scripts/                   # Scripts for setting up and managing the project
│   └── setup.sh               # Script for setting up the environment (dependencies, etc.)
│
├── data/                      # Output directory for scraped data
│   ├── output.csv             # Example CSV file containing scraped data
│   └── output.json            # Example JSON file containing scraped data
│
├── logs/                      # Logs generated during scraping
│   └── scraper.log            # Log file for tracking scraping events and errors
│
├── .env                       # Environment variables for API keys, URLs, etc.
├── .gitignore                 # Files and directories to be ignored by git
├── README.md                  # Project documentation (this file)
└── main.go                    # Entry point for running the Scrappy CLI application
```

## Getting Started

1. **Clone the Repository**: 
   ```
   git clone <repository_url>
   cd Scrappy
   ```

2. **Setup**:
   Run the setup script to install dependencies:
   ```
   ./scripts/setup.sh
   ```

3. **Configure**:
   Update `config/config.yaml` with the websites you want to scrape and other settings.

4. **Run the Scraper**:
   Start scraping using the command:
   ```
   go run main.go scrape
   ```

5. **Automate with Cron**:
   Schedule automated scraping using:
   ```
   go run main.go cron
   ```

## Description of Key Files

- **cmd/root.go**: Entry point for the CLI, sets up the base structure of the command.
- **internal/scraper/scraper.go**: Contains the logic for scraping different websites using Colly.
- **internal/scraper/parser.go**: Handles parsing the raw HTML or JSON data into usable formats.
- **internal/storage/storage.go**: Manages the storage of scraped data into CSV or JSON.
- **internal/scheduler/cron.go**: Contains the logic to run the scraper at regular intervals using Go-cron.
- **config/config.yaml**: Stores configuration such as scraping intervals and target URLs.
- **logs/scraper.log**: Keeps track of scraping events, errors, and execution logs.

## Technologies Used

- **Golang**: Core programming language used for developing the CLI.
- **Cobra**: Library for creating the command-line interface.
- **Colly**: Powerful Golang library for web scraping.
- **Go-cron**: Scheduler for automating scraping tasks.
- **CSV/JSON**: Data formats for exporting scraped data.
