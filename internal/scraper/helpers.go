package scraper

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

// WriteToCSV writes the scraped data to a CSV file
func WriteToCSV(fileName string, data [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

// WriteToJSON writes the scraped data to a JSON file
func WriteToJSON(fileName string, data interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
