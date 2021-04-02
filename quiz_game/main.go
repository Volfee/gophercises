package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// https://golang.org/pkg/flag/#String
var csvFlag = flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")

// var limitFlag = flag.Int("limit", 30, "the time limit for the quiz in sec")

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
	correct, wrong := askQuestions(problems)
	fmt.Printf("Correct answers: %v. Wrong Answers %v.\n", correct, wrong)
}

func askQuestions(problems []problem) (correct, wrong int) {
	var user_answer string
	for _, prob := range problems {
		fmt.Printf(prob.question + "= ")
		fmt.Scanf("%s\n", &user_answer)

		if user_answer == prob.answer {
			correct += 1
		} else {
			wrong += 1
		}
	}
	return
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

type problem struct {
	question, answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
