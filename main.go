package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"time"
)

type Problem struct {
	q string
	a string
}

func main() {
	lines := readFile()
	problems := parseLines(lines)
	quiz(problems)
}

func quiz(pr []Problem) {
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	timer := time.NewTicker(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range pr{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		go func() {
			for _ = range timer.C {
				fmt.Print("_")
			}
		}()

		answerChan := answer()
		elapsed := <-timer.C

		select {
		case <-timer.C:
			fmt.Printf("Your %v time limit has passed. You took %s seconds to answer. \n",
				time.Duration(*timeLimit) * time.Second, elapsed)
			score(correct, pr)
			return
		case answer := <- answerChan:
			if answer == p.a{
				correct++
			}
		}
	}

	score(correct, pr)
}

func answer() chan string{
	answerChan := make(chan string)

	go func() {
		var answer string
		fmt.Scanf("%s\n", &answer)
		answerChan <- answer
	}()

	return answerChan
}

func parseLines(lines [][]string) []Problem{
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i]= Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}

	}

	return ret
}

func readFile() [][]string {
	filename := flag.String("csv", "quiz.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Failed to open csv file: %s ", *filename)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the csv file ")
	}

	return lines
}

func score(c int, p []Problem) {
	fmt.Printf("You scored %d out of %d.\n", c, len(p))
}
