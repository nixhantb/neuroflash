package main

import (
	"fmt"
	"log"
	"time"

	csvparser "neuroflash/src/data"
)

func main() {
	start := time.Now()

	parser := csvparser.CSVParser{Filepath: "test_data.csv"}

	err := parser.ParseCSV()
	if err != nil {
		log.Fatal("Error parsing CSV file:", err)
	}

	stat, err := parser.Describe()
	if err != nil {
		log.Fatal("Error describing CSV file:", err)
	}

	fmt.Println("CSV Description:")
	for _, header := range stat[0] {
		fmt.Printf("%-15s", header)
	}
	fmt.Println()

	for i := 1; i < len(stat); i++ {
		for _, value := range stat[i] {
			fmt.Printf("%-15s", value)
		}
		fmt.Println()
	}

	parser.Top(10)
	fmt.Println("\nTop 10 rows:")
	for _, row := range parser.Data {
		fmt.Println(row)
	}

	parser.Bottom(5)
	fmt.Println("\nBottom 5 rows:")
	for _, row := range parser.Data {
		fmt.Println(row)
	}

	parser.DropCol([]string{"Country"})
	fmt.Println("\nData after dropping column:")
	for _, row := range parser.Data {
		fmt.Println(row)
	}

	records := parser.FillMissing("default_value")
	fmt.Println("\nData after filling missing values:")
	for _, row := range records {
		fmt.Println(row)
	}
	records = parser.DeleteNull("row")

	fmt.Println("\nData after deleting null rows:")
	for _, row := range records {
		fmt.Println(row)
	}

	nullFlags, err := parser.IsNull()
	if err != nil {
		log.Fatal("Error finding null flags:", err)
	}
	nullCount := nullFlags.Sum()
	fmt.Printf("\nNumber of null values: %d\n", nullCount)

	elapsed := time.Since(start)

	fmt.Printf("\nTime taken: %.6f seconds\n", elapsed.Seconds())
}
