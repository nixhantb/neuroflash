# NeuroFlash

A library for machine learning , data science and utilities, including algorithms for regression, classification, clustering, and more. This library is designed to be modular and extendable for various machine learning tasks.

![logo](/docs/assets/neuroflash.jpg)

## Features

- **Algorithms**: Implementations of popular ML algorithms including linear regression, logistic regression, decision trees, KNN, k-means, and PCA.
- **Preprocessing**: Tools for data preprocessing including scaling, normalization, and handling missing values.
- **Evaluation**: Metrics for evaluating model performance such as accuracy, precision, recall, and F1-score.
- **Utilities**: Mathematical and matrix operations essential for ML tasks.

## Data Manipulation with NeuroFlash

NeuroFlash offers efficient data manipulation functions to clean and preprocess your data before feeding it into machine learning models. By using asynchronous operations, you can achieve faster data processing, especially when dealing with large datasets.


```go
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


CSV Description:
Column         Count          Sum            Mean           Min            Max            Q1(25%)        Q2(50%)        Q3(75%)        Variance       Std Deviation  
ID             100            5050.00        50.50          1.00           100.00         26.00          50.50          76.00          833.25         28.87          
Age            100            2931.00        29.31          22.00          35.00          27.00          29.00          32.00          10.51          3.24           
Score          100            8532.20        85.32          75.50          92.50          82.30          85.75          89.00          19.04          4.36      


Top 10 rows:
[ID Name Age Country Score]
[1 Frank 29 Australia 90.2]
[2 Grace 34 Canada 88.0]
[3 Hannah 26 USA 75.5]
[4 Ivy 31 UK 82.3]
[5 Jack 22 Germany 85.0]
[6 Kevin 28 India 79.5]
[7 Liam 30 Australia 92.1]
[8 Mia 27 Canada 84.0]
[9 Nina 32 USA 87.5]
[10 Owen 24 UK 81.0]

Bottom 5 rows:
[ID Name Age Country Score]
[6 Kevin 28 India 79.5]
[7 Liam 30 Australia 92.1]
[8 Mia 27 Canada 84.0]
[9 Nina 32 USA 87.5]
[10 Owen 24 UK 81.0]

Data after dropping column:
[ID Name Age Score]
[6 Kevin 28 79.5]
[7 Liam 30 92.1]
[8 Mia 27 84.0]
[9 Nina 32 87.5]
[10 Owen 24 81.0]

Data after filling missing values:
[ID Name Age Score]
[6 Kevin 28 79.5]
[7 Liam 30 92.1]
[8 Mia 27 84.0]
[9 Nina 32 87.5]
[10 Owen 24 81.0]

Data after deleting null rows:
[ID Name Age Score]
[6 Kevin 28 79.5]
[7 Liam 30 92.1]
[8 Mia 27 84.0]
[9 Nina 32 87.5]
[10 Owen 24 81.0]

Number of null values: 0

Time taken: 0.000685 seconds


```

### Contributing

Contributions are welcome! If you have suggestions or find issues, please submit a pull request or create an issue on our GitHub repository.

### License

This project is licensed under the MIT License 
