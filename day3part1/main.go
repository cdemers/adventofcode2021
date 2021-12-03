package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func asInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counter [2][64]int
	var i int64

	for scanner.Scan() {
		if i, err = strconv.ParseInt(scanner.Text(), 2, 64); err != nil {
			log.Fatal(err)
		}

		for q := 0; q < 64; q++ {
			counter[1&(i>>q)][q]++
		}
	}

	var gamma, epsilon int64

	for q := 0; q < 12; q++ {
		gamma = gamma | asInt64(counter[1][q] > counter[0][q])<<q
		epsilon = gamma ^ 0xFFF
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("gamma=%d, epsilon=%d, power consumption=%d\n", gamma, epsilon, gamma*epsilon)
}
