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

	// Check if there was an error opening the CSV file.
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFile))
	}
	defer file.Close()

	// Parse the CSV File
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	// Check if there was an error parsing the CSV file.
	if err != nil {
		exit("Failed to parse CSV file")
	}

	// Convert the CSV data into a slice of problems.
	problems := parseLines(lines)

	correct := 0 // Initialize a variable to keep track of the number of correct answers.

	// Iterate through the problems and present them to the user for input.
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer) // Read user input for the answer.

		// Check if the user's answer matches the correct answer.
		if answer == p.answer {
			correct++
		}
	}

	// Print the user's score and the total number of questions.
	fmt.Printf("Score: %d Total Questions: %d", correct, len(problems))
}

// parseLines converts a 2D slice of strings (CSV data) into a slice of problem structs.
func parseLines(lines [][]string) []problem {
	returnval := make([]problem, len(lines))

	// Iterate through the CSV data and create problem structs for each question-answer pair.
	for i, line := range lines {
		returnval[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return returnval
}

// problem is a struct that represents a question and its correct answer.
type problem struct {
	question string
	answer   string
}

// exit prints an error message and exits the program with an error code.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
