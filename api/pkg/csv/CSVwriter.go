package csv

import (
	"encoding/csv"
	"os"
)

func MakeCSV() error {
	records := [][]string{
		{"Name", "surname", "occupation"},
		{"ffff", "ffff", "ffff"},
		{"aaaa", "surname", "aaaa"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}

	f, err := os.Create("users.csv")
	defer f.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}

	}
	return nil
}
