package main

import (
	"fmt"
	"log"
	csvparser "neuroflash/src/data"
)

func main() {
	parser := csvparser.CSVParser{Filepath: "test_data.csv"}
	records, err := parser.ParseCSV()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}
