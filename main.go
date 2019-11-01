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
	d := flag.Int("time", 5, "How long should the quiz be - 5 seconds")
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
	// Create a timer so we know when the time has run out
	t := time.NewTimer(time.Second * time.Duration(*d))
	// Create an answer channel so we know if there was a correct answer
	answerCh := make(chan int)
	correct := 0
	// Loop through the problems
	for i, prob := range problems {
		// Ask the question
		fmt.Printf("Problem #%d: %s\n", i+1, prob.q)
		// Create a goroutine that communicates answers through the answerCh
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
		// if the timer finishes, end the game
		case <-t.C:
			exit(fmt.Sprintf("You got %d out of %d correct", correct, len(problems)))
		// if we get an answer back, update the amount of correct we have
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
