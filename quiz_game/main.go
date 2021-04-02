package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// https://golang.org/pkg/flag/#String
var csvFlag = flag.String("csv", "problems.csv",
	"a csv file in format of 'question,answer'")

// var limitFlag = flag.Int("limit", 30, "the time limit for the quiz in sec")

func main() {
	flag.Parse() // https://golang.org/pkg/flag/

	f, err := os.Open(*csvFlag)
	if err != nil {
		panic(err)
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var question string
	var user_answer string
	var correct_answer string
	var correct_answers int
	var wrong_answers int
	for _, line := range lines {
		question, correct_answer = line[0], line[1]
		fmt.Printf(question + "= ")
		fmt.Scanln(&user_answer)

		if user_answer == correct_answer {
			correct_answers += 1
		} else {
			wrong_answers += 1
		}
	}
	fmt.Printf("Correct answers: %v. Wrong Answers %v.\n", correct_answers, wrong_answers)
}
