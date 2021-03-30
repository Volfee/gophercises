package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func main() {
	// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
	file := "problems.csv"

	f, err := os.Open(file)
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
