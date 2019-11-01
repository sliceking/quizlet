package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Define flags with defaults
	f := flag.String("csv", "problems.csv", "the csv with quiz problems - problems.csv")
	d := flag.Duration("time", time.Second*5, "How long should the quiz be - 5 seconds")
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

	problems := parseLines(lines)
	correct := 0
	t := time.NewTimer(*d)
	answerCh := make(chan int)
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, prob.q)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == prob.a {
				answerCh <- 1
			} else {
				answerCh <- 0
			}
		}()
		select {
		case <-t.C:
			exit(fmt.Sprintf("You got %d out of %d correct", correct, len(problems)))
		case message := <-answerCh:
			correct += message
		}
	}
	fmt.Printf("You got %d out of %d correct", correct, len(problems))
}

// takes in an array of lines fromt the csv and returns an array of problem
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
