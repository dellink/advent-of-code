package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := read()

	version, _, value := parse(input, 0)

	fmt.Printf("Part 1: %d\n", version)
	fmt.Printf("Part 2: %d\n", value)
}

func parse(data string, start int) (int, int, int) {
	oldStart := start
	version := toInt(data[start : start+3])
	start += 3
	typeId := toInt(data[start : start+3])
	start += 3
	value := 0
	if typeId == 4 {
		binary := ""
		last := 1
		for last != 0 {
			last = toInt(data[start : start+1])
			start += 1
			binary += data[start : start+4]
			start += 4
		}
		value = toInt(binary)
	} else {
		lengthType := toInt(data[start : start+1])
		start += 1
		subValues := make([]int, 0)
		if lengthType == 0 {
			totalLength := toInt(data[start : start+15])
			start += 15
			for totalLength > 0 {
				subVersion, subStart, subValue := parse(data, start)
				start += subStart
				totalLength -= subStart
				version += subVersion
				subValues = append(subValues, subValue)
			}
		} else {
			nSubPackages := toInt(data[start : start+11])
			start += 11
			for i := 0; i < nSubPackages; i++ {
				subVersion, subStart, subValue := parse(data, start)
				start += subStart
				version += subVersion
				subValues = append(subValues, subValue)
			}
		}

		switch typeId {
		case 0:
			for _, v := range subValues {
				value += v
			}
		case 1:
			value = 1
			for _, v := range subValues {
				value *= v
			}
		case 2:
			sort.Ints(subValues)
			value = subValues[0]
		case 3:
			sort.Ints(subValues)
			value = subValues[len(subValues)-1]
		case 5:
			if subValues[0] > subValues[1] {
				value = 1
			}
		case 6:
			if subValues[0] < subValues[1] {
				value = 1
			}
		case 7:
			if subValues[0] == subValues[1] {
				value = 1
			}
		}

	}
	return version, start - oldStart, value
}

func toInt(data string) int {
	version, _ := strconv.ParseInt(data, 2, 64)
	return int(version)
}

func read() string {
	input, _ := os.Open("day-16/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()

	representation := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	decimal := ""
	for _, s := range scanner.Text() {
		decimal += representation[string(s)]
	}

	return decimal
}
