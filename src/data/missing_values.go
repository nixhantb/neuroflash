package csvparser

import (
	"errors"
)

type NullFlags struct {
	Flags [][]bool
}

func (p *CSVParser) IsNull() (*NullFlags, error) {
	if len(p.Data) == 0 {
		return nil, errors.New("No records to process")
	}
	nullFlags := make([][]bool, len(p.Data))

	for i := range nullFlags {
		nullFlags[i] = make([]bool, len(p.Data[i]))
	}

	for i := 1; i < len(p.Data); i++ {
		for j := 1; j < len(p.Data[0]); j++ {
			nullFlags[i][j] = p.Data[i][j] == ""
		}
	}
	return &NullFlags{Flags: nullFlags}, nil

}

func (nf *NullFlags) Sum() int {
	count := 0
	for i := range nf.Flags {
		for j := range nf.Flags[i] {
			if nf.Flags[i][j] {
				count++
			}
		}
	}
	return count
}

func (p *CSVParser) FillMissing(defaultValue string) [][]string {
	records := p.Data

	for i := 1; i < len(records); i++ {
		for j := 1; j < len(records[i]); j++ {
			if records[i][j] == "" {
				records[i][j] = defaultValue
			}
		}
	}

	return records
}

func (p *CSVParser) DeleteNull(axis string) [][]string {
	records := p.Data

	filteredRecords := [][]string{}
	if axis == "row" {
		for i := 0; i < len(records); i++ {
			nonEmptyRow := []string{}
			for j := 0; j < len(records[i]); j++ {
				if records[i][j] != "" {
					nonEmptyRow = append(nonEmptyRow, records[i][j])
				}
			}
			if len(nonEmptyRow) > 0 {
				filteredRecords = append(filteredRecords, nonEmptyRow)
			}
		}
	}

	if axis == "column" {
		if len(records) == 0 {
			return nil
		}

		numColumns := len(records[0])

		for i := 0; i < len(records); i++ {
			filteredRecords = append(filteredRecords, []string{})
		}

		for j := 0; j < numColumns; j++ {
			nonEmptyColumn := false

			for i := 0; i < len(records); i++ {
				if records[i][j] != "" {
					nonEmptyColumn = true
					break
				}
			}
			if nonEmptyColumn {
				for i := 0; i < len(records); i++ {
					filteredRecords[i] = append(filteredRecords[i], records[i][j])
				}
			}
		}
	}

	return filteredRecords
}

func (p *CSVParser) DropCol(colsToDrop []string) *CSVParser {
	if len(p.Data) == 0 {
		return p
	}

	header := p.Data[0]
	dropIndices := map[int]bool{}

	for _, col := range colsToDrop {
		for i, h := range header {
			if h == col {
				dropIndices[i] = true
				break
			}
		}
	}

	if len(dropIndices) == 0 {
		return p
	}

	newData := make([][]string, len(p.Data))
	for i, row := range p.Data {
		newRow := []string{}
		for j, cell := range row {
			if !dropIndices[j] {
				newRow = append(newRow, cell)
			}
		}
		newData[i] = newRow
	}

	p.Data = newData
	return p
}
