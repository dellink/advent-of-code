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

func solveFirstPart(records []int) int {
	counter := 0
	for i := 1; i < len(records); i++ {
		if records[i] > records[i-1] {
			counter += 1
		}
	}
	return counter
}

func solveSecondPart(records []int) int {
	counter := 0
	for i := 3; i < len(records); i++ {
		if records[i] > records[i-3] {
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

	fmt.Printf("Part 1: %d\n", solveFirstPart(records))
	fmt.Printf("Part 2: %d\n", solveSecondPart(records))
}
