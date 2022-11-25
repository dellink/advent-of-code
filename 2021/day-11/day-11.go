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

	flashes := 0

	step := 0
	stepFlashes := 0

	for stepFlashes != 100 {
		for i := range numbers {
			for j := range numbers[i] {
				numbers[i][j] += 1
			}
		}

		stepFlashes = 0
		wasFlashes := true

		for wasFlashes {
			wasFlashes = false
			for i := range numbers {
				for j := range numbers[i] {
					if numbers[i][j] > 9 {
						numbers[i][j] = 0
						stepFlashes += 1
						wasFlashes = true
						if i > 0 {
							if j > 0 && numbers[i-1][j-1] != 0 {
								numbers[i-1][j-1] += 1 // left top
							}
							if numbers[i-1][j] != 0 {
								numbers[i-1][j] += 1 // top
							}

							if j < len(numbers[i])-1 && numbers[i-1][j+1] != 0 { // right top
								numbers[i-1][j+1] += 1
							}
						}
						if j < len(numbers[i])-1 {
							if numbers[i][j+1] != 0 {
								numbers[i][j+1] += 1 // right
							}
							if i < len(numbers)-1 && numbers[i+1][j+1] != 0 {
								numbers[i+1][j+1] += 1 // right bottom
							}
						}
						if i < len(numbers)-1 {
							if numbers[i+1][j] != 0 {
								numbers[i+1][j] += 1 // bottom
							}
							if j > 0 && numbers[i+1][j-1] != 0 {
								numbers[i+1][j-1] += 1 // left bottom
							}
						}
						if j > 0 && numbers[i][j-1] != 0 {
							numbers[i][j-1] += 1 // left
						}
					}
				}
			}
		}
		if step < 100 {
			flashes += stepFlashes
		}
		step += 1
	}

	fmt.Printf("Part 1: %d\n", flashes)
	fmt.Printf("Part 2: %d\n", step)
}

func read() [][]int {
	input, _ := os.Open("day-11/input.txt")
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
