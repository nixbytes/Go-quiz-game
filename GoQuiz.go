package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define a flag to specify the CSV file name.
	csvFile := flag.String("csv", "proble.csv", "a csv file format of a question")
	flag.Parse()

	// Open the CSV file.
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFileName))
	}
	defer file.Close()
}
