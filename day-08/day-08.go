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
	patterns, values := read()

	var instances int

	for _, display := range values {
		for _, digit := range display {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				instances++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", instances)

	sum := 0

	for i := range patterns {
		codes := decode(patterns[i])
		decoded := ""
		for _, digit := range values[i] {
			decoded += codes[sortSegment(digit)]
		}
		value, _ := strconv.Atoi(decoded)
		sum += value
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func decode(signal []string) map[string]string {
	codes := make(map[string]string)
	instances := make(map[rune]int)

	for _, digit := range signal {
		for _, segment := range digit {
			instances[segment]++
		}
	}

	for _, digit := range signal {
		switch len(digit) {
		case 2:
			codes[sortSegment(digit)] = "1"
		case 3:
			codes[sortSegment(digit)] = "7"
		case 4:
			codes[sortSegment(digit)] = "4"
		case 5:
			for _, segment := range digit {
				if instances[segment] == 4 { // 2 has a segment that has 4 instances in all digits (segment e)
					codes[sortSegment(digit)] = "2"
					break
				} else if instances[segment] == 6 { // 5 has a segment that has 6 instances in all digits (segment b)
					codes[sortSegment(digit)] = "5"
					break
				}
			}
			if _, twoOrFive := codes[sortSegment(digit)]; !twoOrFive {
				codes[sortSegment(digit)] = "3"
			}
		case 6:
			countSegmentsWithKeyInstances := make(map[int]int)
			for _, segment := range digit {
				countSegmentsWithKeyInstances[instances[segment]]++
			}
			if countSegmentsWithKeyInstances[4] == 0 { // only one digit have 6 segments missing the segment e
				codes[sortSegment(digit)] = "9"
			} else if countSegmentsWithKeyInstances[7] == 1 { // only one digit with 6 segments missing the segment d
				codes[sortSegment(digit)] = "0"
			} else {
				codes[sortSegment(digit)] = "6"
			}
		case 7:
			codes[sortSegment(digit)] = "8"
		}
	}
	return codes
}

func sortSegment(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func read() ([][]string, [][]string) {
	input, _ := os.Open("day-08/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var patterns [][]string
	var values [][]string

	for scanner.Scan() {
		signal := strings.Split(scanner.Text(), "|")
		patterns = append(patterns, strings.Fields(signal[0]))
		values = append(values, strings.Fields(signal[1]))
	}

	return patterns, values
}
