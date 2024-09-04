package csvparser

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVParser struct {
	Filepath string
}

func (p *CSVParser) ParseCSV() ([][]string, error) {

	file, err := os.Open(p.Filepath)

	if err != nil {
		log.Fatal("Unable to read the csv file "+p.Filepath, err)
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse the csv file "+p.Filepath, err)
		return nil, err
	}

	return records, nil
}
