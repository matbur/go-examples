package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"flag"
	"io"
)

var (
	fn = flag.String("fn", "problems.csv", "Name of CSV file with questions")
	n  = flag.Int("n", 4, "Number of questions to ask")
)

func main() {
	flag.Parse()

	f, err := os.Open(*fn)
	if err != nil {
		panic(err)
	}

	if err := quiz(csv.NewReader(f)); err != nil {
		panic(err)
	}
}

func quiz(r *csv.Reader) error {
	var good, all int

	for i := 0; i < *n; i++ {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		ok, err := getAnswer(i, rec[0], rec[1])
		if err != nil {
			return err
		}
		if ok {
			good++
		}
		all++
	}

	fmt.Printf("\nYou scored %d out of %d\n", good, all)
	return nil
}

func getAnswer(idx int, question, answer string) (bool, error) {
	var resp string

	fmt.Printf("Problem #%d: %s = ", idx+1, question)
	if _, err := fmt.Scan(&resp); err != nil {
		return false, err
	}
	return resp == answer, nil
}
