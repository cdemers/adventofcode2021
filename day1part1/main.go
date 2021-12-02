package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	incrementCounter := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	previousNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if number > previousNumber {
			incrementCounter++
		}
		previousNumber = number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(incrementCounter)
}
