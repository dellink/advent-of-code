package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day-03/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []string

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", solveFirst(numbers))
	fmt.Printf("Part 2: %d\n", solveSecond(numbers))
}

func solveFirst(numbers []string) int64 {
	var zeroes [12]int

	for _, number := range numbers {
		for i, bit := range number {
			if bit == '0' {
				zeroes[i] += 1
			}
		}
	}

	gammaString := ""
	epsilonString := ""

	for _, count := range zeroes {
		if count > len(numbers)/2 {
			gammaString += "0"
			epsilonString += "1"
		} else {
			gammaString += "1"
			epsilonString += "0"
		}
	}
	return toInt(gammaString) * toInt(epsilonString)
}

func solveSecond(numbers []string) int64 {
	oxygenNumbers := numbers
	oxygenIndex := 0

	for len(oxygenNumbers) != 1 {
		common := getMostCommon(oxygenNumbers, oxygenIndex)
		oxygenNumbers = filter(oxygenNumbers, func(val string) bool {
			return string(val[oxygenIndex]) == common
		})
		oxygenIndex += 1
	}

	coNumbers := numbers
	coIndex := 0

	for len(coNumbers) != 1 {
		common := getMostCommon(coNumbers, coIndex)
		coNumbers = filter(coNumbers, func(val string) bool {
			return string(val[coIndex]) != common
		})
		coIndex += 1
	}

	return toInt(oxygenNumbers[0]) * toInt(coNumbers[0])
}

func filter(arr []string, cond func(string) bool) []string {
	var result []string
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func getMostCommon(arr []string, index int) string {
	zeroes := 0
	for i := range arr {
		if string(arr[i][index]) == "0" {
			zeroes += 1
		}
	}

	if (len(arr) / 2) >= zeroes {
		return "1"
	} else {
		return "0"
	}
}

func toInt(data string) int64 {
	value, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		panic(err)
	}
	return value
}
