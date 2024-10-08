#!/bin/bash

# scripts/setup.sh: A script to set up the Scrappy project environment.

# Exit immediately if a command exits with a non-zero status.
set -e

echo "Starting setup for Scrappy project..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go before running this script."
    exit 1
fi

# Set up Go modules and download dependencies
echo "Initializing Go modules and downloading dependencies..."
go mod tidy

# Create a .env file if it does not exist and add environment variables
ENV_FILE=".env"
if [ ! -f "$ENV_FILE" ]; then
    echo "Creating .env file..."
    touch "$ENV_FILE"
    echo "SCRAPER_ENV=development" >> "$ENV_FILE"
    echo "LOG_LEVEL=info" >> "$ENV_FILE"
    echo "OUTPUT_PATH=./output" >> "$ENV_FILE"
    echo "Environment variables set in .env file."
else
    echo ".env file already exists. Skipping creation."
fi

# Create output directory for storing scraped data
OUTPUT_DIR="./output"
if [ ! -d "$OUTPUT_DIR" ]; then
    echo "Creating output directory..."
    mkdir -p "$OUTPUT_DIR"
    echo "Output directory created at $OUTPUT_DIR."
else
    echo "Output directory already exists. Skipping creation."
fi

# Compile the project to ensure everything is working fine
echo "Compiling the project to ensure setup is correct..."
go build -o bin/scrappy main.go
echo "Project compiled successfully."

echo "Setup complete! You can now run the Scrappy CLI using the following command:"
echo "  ./bin/scrappy scrape --output json"

