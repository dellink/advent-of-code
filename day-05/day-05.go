package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines()

	fmt.Printf("Part 1: %d\n", countPoints(fillLines(lines, false)))
	fmt.Printf("Part 2: %d\n", countPoints(fillLines(lines, true)))
}

func fillLines(lines []string, fillDiagonal bool) map[string]int {
	linesInPoint := make(map[string]int)

	for _, l := range lines {
		var start, end [2]int
		fmt.Sscanf(l, "%d,%d -> %d,%d", &(start[0]), &(start[1]), &(end[0]), &(end[1]))

		if start[0] == end[0] || start[1] == end[1] || fillDiagonal {
			linesInPoint[fmt.Sprintf("%d,%d", start[0], start[1])] += 1

			for start[0] != end[0] || start[1] != end[1] {
				if start[0] > end[0] {
					start[0]--
				} else if start[0] < end[0] {
					start[0]++
				}
				if start[1] > end[1] {
					start[1]--
				} else if start[1] < end[1] {
					start[1]++
				}
				linesInPoint[fmt.Sprintf("%d,%d", start[0], start[1])]++
			}

		}
	}

	return linesInPoint
}

func readLines() []string {
	input, _ := os.Open("day-05/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func countPoints(linesInPoint map[string]int) int {
	countPoints := 0
	for _, v := range linesInPoint {
		if v > 1 {
			countPoints++
		}
	}
	return countPoints
}
