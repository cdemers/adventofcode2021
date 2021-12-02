package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type slidingWindow struct {
	window [3]int
	ptr    int
}

func (w *slidingWindow) Push(i int) int {
	w.window[w.ptr] = i
	if w.ptr >= 2 {
		w.ptr = 0
	} else {
		w.ptr++
	}
	return w.window[0] + w.window[1] + w.window[2]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	previousSum := 0
	window := slidingWindow{}

	scanner := bufio.NewScanner(file)

	for q := 0; q < 3; q++ {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		previousSum = window.Push(number)
	}

	incrementCounter := 0
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum := window.Push(number)

		if sum > previousSum {
			incrementCounter++
		}
		previousSum = sum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(incrementCounter)
}
