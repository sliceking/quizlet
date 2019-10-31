package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// Define flags with defaults
	f := flag.String("file", "problems.csv", "the csv with quiz problems - problems.csv")
	d := flag.Duration("time", time.Second*5, "How long should the quiz be - 5 seconds")
	// Load flag values
	flag.Parse()

	// Open the file
	file, err := os.Open(*f)
	if err != nil {
		panic(err)
	}

	// Create a reader and read from the csv
	cr := csv.NewReader(file)
	for {
		r, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(r)
	}
}
