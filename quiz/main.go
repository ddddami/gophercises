package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
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
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
