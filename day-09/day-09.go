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
	points := read()

	risk := 0
	var lowPoints [][2]int

	for i := range points {
		for j := range points[i] {
			if (j == 0 || points[i][j] < points[i][j-1]) && (i == 0 || points[i][j] < points[i-1][j]) && (j == (len(points[i])-1) || points[i][j] < points[i][j+1]) && (i == (len(points)-1) || points[i][j] < points[i+1][j]) {
				risk += points[i][j] + 1
				coordinates := [2]int{i, j}
				lowPoints = append(lowPoints, coordinates)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", risk)

	var basinPoints []int

	for _, coordinates := range lowPoints {
		copy := append([][]int{}, points...)
		_, count := explore(copy, 0, coordinates[0], coordinates[1])
		basinPoints = append(basinPoints, count)
	}

	sort.Ints(basinPoints)
	fmt.Printf("Part 2: %d\n", basinPoints[len(basinPoints)-3]*basinPoints[len(basinPoints)-2]*basinPoints[len(basinPoints)-1])
}

func explore(points [][]int, count, x, y int) ([][]int, int) {
	if points[x][y] != 9 {
		count++
		points[x][y] = 9
		if x != len(points)-1 {
			points, count = explore(points, count, x+1, y)
		}
		if x != 0 {
			points, count = explore(points, count, x-1, y)
		}
		if y != len(points[x])-1 {
			points, count = explore(points, count, x, y+1)
		}
		if y != 0 {
			points, count = explore(points, count, x, y-1)
		}
	}
	return points, count
}

func read() [][]int {
	input, _ := os.Open("day-09/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var numbers [][]int

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")

		var line []int

		for _, n := range text {
			res, _ := strconv.Atoi(n)
			line = append(line, res)
		}

		numbers = append(numbers, line)
	}

	return numbers
}
