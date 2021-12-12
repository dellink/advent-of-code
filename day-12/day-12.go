package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	lines := read()

	connections := make(map[string][]string)

	for _, line := range lines {
		cave := strings.Split(line, "-")
		if cave[1] != "start" {
			if _, ok := connections[cave[0]]; ok {
				connections[cave[0]] = append(connections[cave[0]], cave[1])
			} else {
				connections[cave[0]] = []string{cave[1]}
			}
		}
		if cave[0] != "start" {
			if _, ok := connections[cave[1]]; ok {
				connections[cave[1]] = append(connections[cave[1]], cave[0])
			} else {
				connections[cave[1]] = []string{cave[0]}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", solve(connections, false))
	fmt.Printf("Part 2: %d\n", solve(connections, true))
}

func read() []string {
	input, _ := os.Open("day-12/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func solve(connections map[string][]string, allowDouble bool) int {
	pathes := 0
	stack := [][]string{}

	for _, v := range connections["start"] {
		stack = append(stack, []string{v})
	}

	for len(stack) > 0 {
		i := len(stack) - 1
		path := stack[i]
		stack = stack[:i]

		for _, dest := range connections[path[len(path)-1]] {
			if dest == "end" {
				pathes += 1
			} else {
				if allowDouble {
					if !isLower(dest) || !containsDouble(path) || !contains(path, dest) {
						copy := append([]string{}, path...)
						stack = append(stack, append(copy, dest))
					}
				} else {
					if !isLower(dest) || !contains(path, dest) {
						copy := append([]string{}, path...)
						stack = append(stack, append(copy, dest))
					}
				}
			}
		}
	}

	return pathes
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func containsDouble(s []string) bool {
	contains := make(map[string]int)
	for _, a := range s {
		contains[a] += 1
	}
	hasDouble := false
	for k, v := range contains {
		if isLower(k) && v > 1 {
			hasDouble = true
			break
		}
	}
	return hasDouble
}
