package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := readNumbers()
	var waiting1 [9]int
	var waiting2 [9]int

	for _, n := range numbers {
		waiting1[n]++
		waiting2[n]++
	}

	fmt.Printf("Part 1: %d\n", solve(waiting1, 80))
	fmt.Printf("Part 2: %d\n", solve(waiting2, 256))
}

func solve(waiting [9]int, days int) int {
	for i := 0; i < days; i++ {
		newFishes := waiting[0]

		for restDays := range waiting[:len(waiting)-1] {
			waiting[restDays] = waiting[restDays+1]
		}

		waiting[6] += newFishes
		waiting[8] = newFishes
	}

	fishes := 0

	for _, f := range waiting {
		fishes += f
	}

	return fishes
}

func readNumbers() []int {
	input, _ := os.Open("day-06/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var numbers []int
	for _, n := range strings.Split(lines[0], ",") {
		res, _ := strconv.Atoi(n)
		numbers = append(numbers, res)
	}

	return numbers
}
