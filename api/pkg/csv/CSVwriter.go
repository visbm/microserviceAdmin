package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

//MakeCSV make csv file and write data
func MakeCSV(data interface{}) error {	

	
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '\t'
		return gocsv.NewSafeCSVWriter(writer)
	})

	file, err := os.Create("/api/pkg/csv/pkgfile.csv")

	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return err
	}
	fmt.Println("form MakeCSV  ")
	return nil
}
