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

	// Intialize var for tracking quiz results
	//	correct := 0
	//	total := len(lines)

	/*
	   	for _, line := range lines {
	   		question := line[0]
	   		expectedAnswer := strings.TrimSpace(line[1])

	   		fmt.Printf("Question: %s\nYour answer: ", question)

	   		var userAnswer string
	   		_, err := fmt.Scan(&userAnswer)
	   		if err != nil {
	   			exit("Error reading user input.")
	   		}

	   		if userAnswer == expectedAnswer {
	   			correct++
	   		}
	   	}

	   // Print the quiz results.
	   fmt.Printf("You got %d out of %d questions correct.\n", correct, total)
	*/
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
