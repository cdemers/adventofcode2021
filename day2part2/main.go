package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func split(s string) (o []string) {
	if len(s) == 0 {
		return []string{}
	}
	p := 0
	for q := 0; q < len(s); q++ {
		if s[q] == ' ' {
			o = append(o, s[p:q])
			p = q + 1
		}
	}
	return append(o, s[p:])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depth int
	var horizontalPosition int
	var aim int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tokens := split(scanner.Text())
		positions, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}

		switch tokens[0] {

		case "forward":
			horizontalPosition += positions
			depth += aim * positions
		case "up":
			aim -= positions
		case "down":
			aim += positions
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontalPosition * depth)
}
