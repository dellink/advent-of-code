package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := read()
	size := len(numbers)
	numbers = generate(numbers)

	full := [][]int{}

	for i := range numbers {
		full = append(full, append([]int{}, numbers[i]...))
	}

	for i := range full {
		for j := range full[i] {
			full[i][j] = 0
		}
	}

	full[0][0] = numbers[0][0]

	for i := range numbers {
		for j := range numbers[i] {
			if j < len(numbers[i])-1 {
				if full[i][j+1] > 0 {
					if full[i][j+1] > full[i][j]+numbers[i][j+1] {
						full[i][j+1] = full[i][j] + numbers[i][j+1]
					}
				} else {
					full[i][j+1] = full[i][j] + numbers[i][j+1]
				}
			}
			if i < len(numbers[i])-1 {
				if full[i+1][j] > 0 {
					if full[i+1][j] > full[i][j]+numbers[i+1][j] {
						full[i+1][j] = full[i][j] + numbers[i+1][j]
					}
				} else {
					full[i+1][j] = full[i][j] + numbers[i+1][j]
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", full[size-1][size-1]-full[0][0])
	fmt.Printf("Part 2: %d\n", full[len(full)-1][len(full[0])-1]-full[0][0])
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
