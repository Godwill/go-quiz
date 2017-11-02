package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
)

func main() {
	filename := flag.String("csv", "quiz.csv", "a csv file in the format of 'question, answer'")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Failed to open csv file: %s ", *filename)
	}

	r := csv.NewReader(file)

	_ = r

	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the csv file ")
	}

	problems := parseLines(lines)
	correct := 0
	for i, p := range problems{
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a{
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
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

type Problem struct {
	q string
	a string
}