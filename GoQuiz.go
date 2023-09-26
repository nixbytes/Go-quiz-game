package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define a flag to specify the CSV file name.
	csvFile := flag.String("csv", "problem.csv", "a csv file format of a question")
	flag.Parse()

	// Open the CSV file, get the value not the pointer
	file, err := os.Open(*csvFile)

	// using exit function for CSV can't Open
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFile))
	}
	defer file.Close()

	// Parse the CSV File
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse CSV file")
	}

	problems := parseLines(lines)
	fmt.Println(problems)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scan("%s\n", &anwser)
		if anwser == p.a {
			fmt.PrintLn("Correct Anwser!")

		}
	}

}

func parseLines(lines [][]string) []problem {

	returnval := make([]problem, len(lines))

	for i, line := range lines {
		returnval[i] = problem{
			question: line[0],
			answer:   line[1],
		}

	}
	return returnval
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
