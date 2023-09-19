package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define a flag to specify the CSV file name.
	csvFile := flag.String("csv", "proble.csv", "a csv file format of a question")
	flag.Parse()

	// Open the CSV file.
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFile))
	}
	defer file.Close()

	// Parse the CSV File
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
