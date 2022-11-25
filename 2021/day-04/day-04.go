package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()
	numbers, boards := getData(lines)

	bingo, winIndex := game(numbers, boards)
	total := getTotal(bingo, boards[winIndex/5*5:winIndex/5*5+5])

	fmt.Printf("Part 1: %d\n", total*bingo[len(bingo)-1])

	totalLast := 0
	bingo = make([]int, 0)

	for _, n := range numbers {
		if len(boards) == 0 {
			break
		}
		bingo = append(bingo, n)
		if len(bingo) > 4 {
			for i, boardLine := range boards {
				if contains(bingo, boardLine[0]) && contains(bingo, boardLine[1]) && contains(bingo, boardLine[2]) && contains(bingo, boardLine[3]) && contains(bingo, boardLine[4]) {
					if len(boards) == 5 {
						totalLast = getTotal(bingo, boards)
						boards = make([][]int, 0)
					} else {
						boards = append(boards[:(i/5*5)], boards[(i/5*5+5):]...)
					}
					break
				}
			}
			length := len(boards)
			for i := 0; i < len(boards); i += 5 {
				if length < len(boards) {
					break
				}
				for j := 0; j < 5; j += 1 {
					if contains(bingo, boards[i][j]) && contains(bingo, boards[i+1][j]) && contains(bingo, boards[i+2][j]) && contains(bingo, boards[i+3][j]) && contains(bingo, boards[i+4][j]) {
						if len(boards) == 5 {
							totalLast = getTotal(bingo, boards)
							boards = make([][]int, 0)
						} else {
							boards = append(boards[:(i/5*5)], boards[(i/5*5+5):]...)
						}
						break
					}
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", totalLast*bingo[len(bingo)-1])

}

func getLines() []string {
	file, err := os.Open("day-04/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func getData(lines []string) ([]int, [][]int) {
	numbersLine := strings.Split(lines[0], ",")
	var numbers []int

	for i := range numbersLine {
		numbers = append(numbers, int(toInt(numbersLine[i])))
	}

	var boards [][]int

	for i := range lines {
		if i > 0 && len(lines[i]) > 0 {
			var line []int
			for _, s := range strings.Split(lines[i], " ") {
				if len(s) > 0 {
					strings.ReplaceAll(s, " ", "")
					line = append(line, int(toInt(s)))
				}
			}
			boards = append(boards, line)
		}
	}

	return numbers, boards
}

func game(numbers []int, boards [][]int) ([]int, int) {
	var bingo []int
	var winIndex int

	for _, n := range numbers {
		if winIndex > 0 {
			break
		}
		bingo = append(bingo, n)
		if len(bingo) > 4 {
			for i, boardLine := range boards {
				if contains(bingo, boardLine[0]) && contains(bingo, boardLine[1]) && contains(bingo, boardLine[2]) && contains(bingo, boardLine[3]) && contains(bingo, boardLine[4]) {
					winIndex = i
				}
			}
		}
	}

	return bingo, winIndex
}

func getTotal(bingo []int, boards [][]int) int {
	total := 0

	for _, board := range boards {
		for _, n := range board {
			if !contains(bingo, n) {
				total += n
			}
		}
	}

	return total
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
