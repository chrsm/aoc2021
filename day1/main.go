package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("testdata/real.input")
	if err != nil {
		log.Fatalf("failed to open real.input: %s", err)
	}

	r, err := parse(input)
	if err != nil {
		log.Fatalf("failed to parse real.input: %s", err)
	}

	log.Printf("count: %d", countInc(r))
}

func parse(buf []byte) ([]int, error) {
	var (
		records []int
		split   = bytes.Split(buf, []byte{'\n'})
	)

	for i := range split {
		if len(split[i]) == 0 {
			continue
		}

		v, err := strconv.Atoi(string(split[i]))
		if err != nil {
			return nil, err
		}

		records = append(records, v)
	}

	return records, nil
}

func countInc(rec []int) int {
	var prev, count int

	for i := range rec {
		if i == 0 {
			prev = rec[i]
			continue
		}

		if prev < rec[i] {
			count++
		}

		prev = rec[i]
	}

	return count
}
