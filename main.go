package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags with defaults
	f := flag.String("csv", "problems.csv", "the csv with quiz problems - problems.csv")
	// d := flag.Duration("time", time.Second*5, "How long should the quiz be - 5 seconds")
	// Load flag values
	flag.Parse()

	// Open the file
	file, err := os.Open(*f)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV: %s\n", *f))
	}

	// Create a reader and read from the csv
	cr := csv.NewReader(file)
	lines, err := cr.ReadAll()
	if err != nil {
		exit("Failed to parse CSV")
	}
	fmt.Println(lines)
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
