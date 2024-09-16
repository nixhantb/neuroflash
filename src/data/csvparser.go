package csvparser

import (
	"encoding/csv"
	"errors"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type CSVParser struct {
	Filepath string
	Data     [][]string
}

type Statistics struct {
	Count        int
	Sum          float64
	Mean         float64
	Min          float64
	Max          float64
	Median       float64
	Variance     float64
	StdDeviation float64
	Q1           float64
	Q2           float64
	Q3           float64
}

func (p *CSVParser) ParseCSV() error {

	file, err := os.Open(p.Filepath)

	if err != nil {
		log.Fatal("Unable to read the csv file "+p.Filepath, err)
		return err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse the csv file "+p.Filepath, err)
		return err
	}

	p.Data = records
	return nil
}

func (p *CSVParser) Top(rows ...int) *CSVParser {

	if len(p.Data) == 0 {
		return p
	}

	numOfRows := 5

	if len(rows) > 0 {
		numOfRows = rows[0]
	}

	if numOfRows > len(p.Data) {
		numOfRows = len(p.Data)
	}

	p.Data = p.Data[:numOfRows+1]

	return p

}

func (p *CSVParser) Bottom(rows ...int) *CSVParser {
	if len(p.Data) == 0 {
		return p
	}

	numOfRows := 5
	if len(rows) > 0 {
		numOfRows = rows[0]
	}
	if numOfRows > len(p.Data) {
		numOfRows = len(p.Data)
	}
	p.Data = append([][]string{p.Data[0]}, p.Data[len(p.Data)-numOfRows:]...)
	return p
}

func (p *CSVParser) Describe() ([][]string, error) {

	if len(p.Data) < 2 {
		return nil, errors.New(`insufficient csv record`)
	}

	headers := p.Data[0]

	response := [][]string{
		{"Column", "Count", "Sum", "Mean", "Min", "Max", "Q1(25%)", "Q2(50%)", "Q3(75%)", "Variance", "Std Deviation"},
	}

	for colIndex, header := range headers {

		var values []float64

		for i := 1; i < len(p.Data); i++ {

			value, err := strconv.ParseFloat(p.Data[i][colIndex], 64)

			if err == nil {
				values = append(values, value)
			}
		}

		if len(values) == 0 {
			continue
		}

		stats := calculateStatistics(values)
		response = append(response, []string{
			header,
			strconv.Itoa(stats.Count),
			strconv.FormatFloat(stats.Sum, 'f', 2, 64),
			strconv.FormatFloat(stats.Mean, 'f', 2, 64),
			strconv.FormatFloat(stats.Min, 'f', 2, 64),
			strconv.FormatFloat(stats.Max, 'f', 2, 64),
			strconv.FormatFloat(stats.Q1, 'f', 2, 64),
			strconv.FormatFloat(stats.Q2, 'f', 2, 64),
			strconv.FormatFloat(stats.Q3, 'f', 2, 64),
			strconv.FormatFloat(stats.Variance, 'f', 2, 64),
			strconv.FormatFloat(stats.StdDeviation, 'f', 2, 64),
		})

	}
	return response, nil

}

func calculateStatistics(values []float64) Statistics {
	count := len(values)
	sum := 0.0

	sort.Float64s(values)
	min := values[0]
	max := values[count-1]

	for _, value := range values {
		sum += value
	}

	mean := sum / float64(count)

	q2 := 0.0
	if count%2 == 0 {
		q2 = (values[count/2-1] + values[count/2]) / 2
	} else {
		q2 = values[count/2]
	}

	variance := 0.0
	for _, value := range values {
		variance += math.Pow(value-mean, 2)
	}
	variance /= float64(count)

	stdDeviation := math.Sqrt(variance)

	q1_index := int64((float64(count) + 1) / 4)
	q3_index := 3 * q1_index
	q1 := values[q1_index]
	q3 := values[q3_index]

	return Statistics{
		Count:        count,
		Sum:          sum,
		Mean:         mean,
		Min:          min,
		Max:          max,
		Q2:           q2,
		Variance:     variance,
		StdDeviation: stdDeviation,
		Q1:           q1,
		Q3:           q3,
	}
}
