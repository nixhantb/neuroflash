package csvparser

func (p *CSVParser) IsNull() ([][]bool, error) {

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

	return nullFlags, nil

}
