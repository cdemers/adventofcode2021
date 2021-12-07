package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var bitWidth = 12
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputs = make([]string, 0, 1000)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var skew int
	var bitToKeep byte

	filteredInputs := make([]string, 0, len(inputs))
	prevFilteredInputs := inputs

	var generatorRating int64

	for q := 0; q < bitWidth; q++ {
		skew = 0
		for _, input := range prevFilteredInputs {
			if input[q] == byte('1') {
				skew++
			}
			if input[q] == byte('0') {
				skew--
			}
		}

		if skew >= 0 {
			bitToKeep = byte('1')
		} else {
			bitToKeep = byte('0')
		}

		filteredInputs = filteredInputs[:0]

		for _, v := range prevFilteredInputs {
			if v[q] == bitToKeep {
				filteredInputs = append(filteredInputs, v)
			}
		}

		prevFilteredInputs = filteredInputs
		if len(filteredInputs) == 1 {
			break
		}
	}
	if generatorRating, err = strconv.ParseInt(filteredInputs[0], 2, 64); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Oxygen Generator Rating is %d (%s)\n", generatorRating, filteredInputs[0])

	filteredInputs = make([]string, 0, len(inputs))
	prevFilteredInputs = inputs

	var scrubberRating int64

	for q := 0; q < bitWidth; q++ {
		skew = 0
		for _, input := range prevFilteredInputs {
			if input[q] == byte('1') {
				skew++
			}
			if input[q] == byte('0') {
				skew--
			}
		}

		if skew >= 0 {
			bitToKeep = byte('0')
		} else {
			bitToKeep = byte('1')
		}

		filteredInputs = filteredInputs[:0]

		for _, v := range prevFilteredInputs {
			if v[q] == bitToKeep {
				filteredInputs = append(filteredInputs, v)
			}
		}

		prevFilteredInputs = filteredInputs
		if len(filteredInputs) == 1 {
			break
		}
	}

	if scrubberRating, err = strconv.ParseInt(filteredInputs[0], 2, 64); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Carbon Dioxyde Scrubber Rating is %d (%s)\n", scrubberRating, filteredInputs[0])

	fmt.Printf("Life Support Rating: %d\n", generatorRating*scrubberRating)
}
