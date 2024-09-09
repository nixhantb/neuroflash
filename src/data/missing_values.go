package csvparser

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
