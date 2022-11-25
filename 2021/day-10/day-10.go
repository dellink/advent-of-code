package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var closing = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
var errorPoints = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var incompletePoints = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

func main() {
	lines := read()

	var errorScore = 0
	var incompleteScores []int

	for _, line := range lines {
		lineScore, stack := parse(line)
		if lineScore == 0 {
			score := 0
			for _, symbol := range stack {
				score = score*5 + incompletePoints[symbol]
			}
			incompleteScores = append(incompleteScores, score)
		} else {
			errorScore += lineScore
		}
	}

	fmt.Printf("Part 1: %d\n", errorScore)

	sort.Ints(incompleteScores)

	fmt.Printf("Part 2: %d\n", incompleteScores[len(incompleteScores)/2])
}

func parse(line string) (int, []rune) {
	var stack []rune
	for _, symbol := range line {
		switch symbol {
		case '(', '[', '{', '<':
			stack = append([]rune{closing[symbol]}, stack...)
		case ')', ']', '}', '>':
			if len(stack) == 0 || stack[0] != symbol {
				return errorPoints[symbol], stack
			} else {
				stack = stack[1:]
			}
		}
	}

	return 0, stack
}

func read() []string {
	input, _ := os.Open("day-10/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
