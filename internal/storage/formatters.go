package storage

import (
	"encoding/json"
	"fmt"
)

// FormatCSV formats data into a CSV-compatible structure
func FormatCSV(data [][]string) string {
	var csvData string

	for _, record := range data {
		csvData += fmt.Sprintf("%s\n", csvJoin(record))
	}
	return csvData
}

// FormatJSON formats data into JSON
func FormatJSON(data interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// Helper function to join a slice of strings into a CSV-compatible string
func csvJoin(record []string) string {
	return "\"" + record[0] + "\",\"" + record[1] + "\",\"" + record[2] + "\"" // Example for three fields
}
