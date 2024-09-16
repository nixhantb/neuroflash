package csvparser

import "errors"

type NullFlags struct {
	flags [][]bool
}

func (p *CSVParser) IsNull() (*NullFlags, error) {

	records, err := p.ParseCSV()
	if err != nil {
		return nil, err
	}

	nullFlags := make([][]bool, len(records))

	for i := range nullFlags {
		nullFlags[i] = make([]bool, len(records[i]))
	}

	for i := 1; i < len(records); i++ {

		for j := 1; j < len(records[0]); j++ {

			if records[i][j] == "" {
				nullFlags[i][j] = true
			} else {
				nullFlags[i][j] = false
			}
		}
	}

	return &NullFlags{flags: nullFlags}, nil

}

func (nf *NullFlags) Sum() int {
	count := 0

	for i := range nf.flags {
		for j := range nf.flags[i] {
			if nf.flags[i][j] {
				count++
			}
		}
	}
	return count
}

func (p *CSVParser) FillMissing(defaultValue string) ([][]string, error) {

	records, err := p.ParseCSV()

	if err != nil {
		return nil, err
	}

	nullFlags, err := p.IsNull()

	if err != nil {
		return nil, err
	}

	for i := 1; i < len(records); i++ {

		for j := 1; i < len(records[i]); j++ {

			if nullFlags.flags[i][j] {
				records[i][j] = defaultValue
			}
		}
	}

	return records, nil

}

func (p *CSVParser) DeleteNull(axis string) ([][]string, error) {

	records, err := p.ParseCSV()

	if len(records) == 0 {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

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
			return nil, nil
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

	return filteredRecords, nil

}

func (p *CSVParser) DropCol(colsToDrop []string) ([][]string, error) {

	records, err := p.ParseCSV()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.New("No records found")
	}

	header := records[0]
	dropIndices := map[int]bool{}

	for _, col := range colsToDrop {
		for idx, headerCol := range header {

			if col == headerCol {
				dropIndices[idx] = true
			}
		}
	}

	filteredRecords := [][]string{}

	for _, record := range records {

		newRecord := []string{}

		for idx, value := range record {

			if !dropIndices[idx] {
				newRecord = append(newRecord, value)

			}
		}
		filteredRecords = append(filteredRecords, newRecord)
	}
	return filteredRecords, nil
}
