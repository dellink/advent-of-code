package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func solve(records []int, index int) int {
	counter := 0
	for i := index; i < len(records); i++ {
		if records[i] > records[i-index] {
			counter += 1
		}
	}
	return counter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records []int

	for scanner.Scan() {
		records = append(records, toInt(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", solve(records, 1))
	fmt.Printf("Part 2: %d\n", solve(records, 3))
}
