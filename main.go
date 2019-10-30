package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

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
