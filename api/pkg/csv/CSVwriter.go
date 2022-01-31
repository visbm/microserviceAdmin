package csv

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

//MakeCSV make csv file and write data
func MakeCSV(data interface{}) error {

	file, err := os.Create("/api/pkg/csv/pkgfile.csv")
	fmt.Println(file)
	fmt.Println(file.Name())

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
