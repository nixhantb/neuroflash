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
	nullCount := nullFlags.Sum()
	fmt.Printf("Number of null values: %d\n", nullCount)

	elapsed := time.Since(start)

	fmt.Printf("Time taken: %.6f seconds\n", elapsed.Seconds())

}

Column         Count          Sum            Mean           Min            Max            Q1(25%)        Q2(50%)        Q3(75%)        Variance       Std Deviation  
ID             100            5050.00        50.50          1.00           100.00         26.00          50.50          76.00          833.25         28.87          
Age            100            2931.00        29.31          22.00          35.00          27.00          29.00          32.00          10.51          3.24           
Score          100            8532.20        85.32          75.50          92.50          82.30          85.75          89.00          19.04          4.36           
Number of null values: 0
Time taken: 0.000511 seconds
```
