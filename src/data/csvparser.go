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

func (p *CSVParser) Top(rows ...int) ([][]string, error) {

	records, err := p.ParseCSV()

	if err != nil {
		log.Fatal("Unable to load the csv file", err)
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	numOfRows := 5

	if len(rows) > 0 {
		numOfRows = rows[0]
	}

	if numOfRows > len(records) {
		numOfRows = len(records)
	}

	responseHead := records[:numOfRows+1]

	return responseHead, nil

}

func (p *CSVParser) Bottom(rows ...int) ([][]string, error) {

	records, err := p.ParseCSV()
	if err != nil {
		log.Fatal("Unable to load the csv file ", err)
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	header := records[0]

	numsOfRows := 5

	if len(rows) > 0 {
		numsOfRows = rows[0]
	}

	if numsOfRows > len(records) {
		numsOfRows = len(records)
	}

	responseBottom := records[len(records)-numsOfRows:]

	result := [][]string{header}
	result = append(result, responseBottom...)
	return result, nil
}
