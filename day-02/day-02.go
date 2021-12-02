package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var commands []string

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", solveFirst(commands))
	fmt.Printf("Part 2: %d\n", solveSecond(commands))
}

func solveFirst(commands []string) int {
	position := 0
	depth := 0

	for _, command := range commands {
		c := strings.Split(command, " ")

		v, err := strconv.Atoi(c[1])
		if err != nil {
			panic(err)
		}

		switch c[0] {
		case "forward":
			position += v
		case "down":
			depth += v
		case "up":
			depth -= v
		}

	}

	return position * depth
}

func solveSecond(commands []string) int {
	position := 0
	depth := 0
	aim := 0

	for _, command := range commands {
		c := strings.Split(command, " ")

		v, err := strconv.Atoi(c[1])
		if err != nil {
			panic(err)
		}

		switch c[0] {
		case "forward":
			position += v
			depth += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		}

	}

	return position * depth
}
