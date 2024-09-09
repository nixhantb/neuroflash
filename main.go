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

	stat, err := parser.Describe()
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

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

	nullFlags, err := parser.IsNull()
	if err != nil {
		log.Fatal("Error finding null flags:", err)
	}

	// Print the nullFlags matrix
	fmt.Println("Null Flags Matrix:")
	for i := range nullFlags {
		for j := range nullFlags[i] {
			fmt.Printf("%v ", nullFlags[i][j])
		}
		fmt.Println()
	}

	elapsed := time.Since(start)

	fmt.Printf("Time taken: %.6f seconds\n", elapsed.Seconds())

}
