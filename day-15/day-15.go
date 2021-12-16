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
	numbers := read()
	size := len(numbers)
	numbers = generate(numbers)

	visited := [][]bool{}

	for i := range numbers {
		line := []bool{}
		for range numbers[i] {
			line = append(line, false)
		}
		visited = append(visited, line)
	}

	queue := append([][2][2]int{}, [2][2]int{{0, 0}, {0}})

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		if point[0][0] == size-1 && point[0][1] == size-1 {
			fmt.Printf("Part 1: %d\n", point[1][0])
		}

		if point[0][0] == len(numbers)-1 && point[0][1] == len(numbers)-1 {
			fmt.Printf("Part 2: %d\n", point[1][0])
			break
		}

		for _, direction := range directions {
			ii := point[0][0] + direction[0]
			jj := point[0][1] + direction[1]

			if ii < 0 || jj < 0 || len(numbers)-1 < ii || len(numbers)-1 < jj || visited[ii][jj] {
				continue
			}

			visited[ii][jj] = true
			queue = append(queue, [2][2]int{{ii, jj}, {point[1][0] + numbers[ii][jj]}})
		}

		sort.Slice(queue, func(i, j int) bool {
			return queue[i][1][0] < queue[j][1][0]
		})
	}
}

func read() [][]int {
	input, _ := os.Open("day-15/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	numbers := [][]int{}

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")

		line := []int{}

		for _, n := range text {
			full, _ := strconv.Atoi(n)
			line = append(line, full)
		}

		numbers = append(numbers, line)
	}

	return numbers
}

func generate(numbers [][]int) [][]int {
	size := len(numbers)

	for lineIndex := range numbers {
		newLine := append([]int{}, numbers[lineIndex]...)
		for i := 1; i < 5; i += 1 {
			for j := range numbers[lineIndex] {
				v := numbers[lineIndex][j] + i
				if v > 9 {
					v -= 9
				}
				newLine = append(newLine, v)
			}
		}
		numbers[lineIndex] = newLine
	}

	for i := 1; i < 5; i += 1 {

		for lineIndex, line := range numbers {
			if lineIndex > size-1 {
				break
			}
			newLine := []int{}
			for j := range line {
				v := line[j] + i
				if v > 9 {
					v -= 9
				}
				newLine = append(newLine, v)
			}
			numbers = append(numbers, newLine)
		}
	}

	return numbers
}
