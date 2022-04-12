package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	timeLimit := flag.Int("limit", 30, "Time limit for the Quiz in seconds")

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var correctCount int

LOOP:
	for index, value := range problems {
		answerChannel := make(chan string)

		fmt.Printf("Question #%d: %s = ", index+1, value.question)

		go func() {
			var answer string

			_, err := fmt.Scanf("%s\n", &answer)
			if err != nil {
				exitWithMessage("Failed to scan answer!")
			}

			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break LOOP
		case answer := <-answerChannel:
			if answer == value.answer {
				correctCount++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctCount, len(problems))
}
