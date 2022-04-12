package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func exitWithMessage(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseRows(rows [][]string) []problem {
	result := make([]problem, len(rows))

	for index, row := range rows {
		result[index] = problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}

	return result
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "CSV file with format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exitWithMessage(fmt.Sprintf("Failed to open the CSV file: %s\n!", *csvFile))
	}

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		exitWithMessage("Failed to parse CSV file!")
	}

	problems := parseRows(rows)

	var correctCount int

	for index, value := range problems {
		var answer string

		fmt.Printf("Question #%d: %s = ", index+1, value.question)

		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			exitWithMessage("Failed to scan answer!")
		}

		if answer == value.answer {
			correctCount++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctCount, len(problems))
}
