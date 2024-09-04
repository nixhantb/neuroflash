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
	csvparser "neuroflash/src/data"
)

func main() {
	parser := csvparser.CSVParser{Filepath: "temp.csv"}
	records, err := parser.ParseCSV()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)

}

[[ID Name Age Country Score] 
[1 Alice 25 USA 88.5] 
[2 Bob  UK 92.0]
[3 Charlie 35 Canada ] 
[4 David 28 Germany 80.0] 
[5 Eve 22 Australia 85.0]]
```
