package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv filr in thr format (question,answer)")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	problems := parseLines(lines)

	points := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.q)
		var answer string

		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			points++
			fmt.Println("Correct!")
		}

	}
	fmt.Printf("You got %d out of %d points\n", points, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
