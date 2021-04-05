package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// https://golang.org/pkg/flag/#String
var csvFlag = flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
var limitFlag = flag.Int("limit", 30, "the time limit for the quiz in sec")

func main() {
	flag.Parse() // https://golang.org/pkg/flag/

	f, err := os.Open(*csvFlag)
	if err != nil {
		exit("Failed to open the CSV file: %s" + *csvFlag)
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		exit("Failed to read file. Err: " + err.Error())
	}
	problems := parseLines(lines)

	var correct int
	isCorrect := make(chan bool)
	timer := time.NewTimer(time.Duration(*limitFlag) * time.Second)
loop:
	for _, problem := range problems {
		go askQuestion(problem, isCorrect)

		select {
		case <-timer.C:
			break loop

		case correct_answer := <-isCorrect:
			if correct_answer {
				correct += 1
			}
		}
	}

	printResultsAndExit(correct, len(problems))
}

func askQuestion(p problem, answer chan bool) {
	var user_answer string
	fmt.Printf(p.question + "= ")
	fmt.Scanf("%s\n", &user_answer)
	isCorrect := user_answer == p.answer
	answer <- isCorrect
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for idx, line := range lines {
		ret[idx] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func printResultsAndExit(correct, total int) {
	fmt.Printf("\n%v correct answers out of %v.\n", correct, total)
	os.Exit(1)
}

type problem struct {
	question, answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
