package core

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/jszwec/csvutil"
)

func SaveToFile(format string, filename string, jobs []Job) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "json":
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		return encoder.Encode(jobs)
	case "csv":
		writer := csv.NewWriter(file)
		defer writer.Flush()

		enc := csvutil.NewEncoder(writer)
		return enc.Encode(jobs)
	default:
		return os.WriteFile(filename, []byte("Unsupported format"), 0644)
	}
}
