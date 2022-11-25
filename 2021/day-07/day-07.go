package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := readNumbers()
	sort.Ints(numbers)

	fmt.Printf("Part 1: %d\n", solveFirstPart(numbers))
	fmt.Printf("Part 2: %d\n", solveSecondPart(numbers))
}

func solveFirstPart(numbers []int) int {
	mNumber := len(numbers) / 2
	median := numbers[mNumber]

	if len(numbers)%2 == 0 {
		median = (numbers[mNumber-1] + numbers[mNumber]) / 2
	}

	total := 0

	for _, n := range numbers {
		if median > n {
			total += median - n
		} else {
			total += n - median
		}
	}

	return total
}

func solveSecondPart(numbers []int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	average := total / len(numbers)

	total2 := 0

	for _, n := range numbers {
		steps := n - average
		if average > n {
			steps = average - n
		}

		total2 += (steps * (steps + 1)) / 2
	}

	return total2
}

func readNumbers() []int {
	input, _ := os.Open("day-07/input.txt")
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
