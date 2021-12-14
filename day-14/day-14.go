package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	input, _ := os.Open("day-14/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var template string

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		template = scanner.Text()
	}

	rules := map[string]string{}

	for scanner.Scan() {
		var pair string
		var insertion string
		fmt.Sscanf(scanner.Text(), "%s -> %s", &pair, &insertion)

		rules[pair] = insertion
	}

	fmt.Printf("Part 1: %d\n", solve(template, rules, 10))
	fmt.Printf("Part 2: %d\n", solve(template, rules, 40))
}

func solve(template string, rules map[string]string, steps int) int {
	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 0; i < steps; i++ {
		newPairs := map[string]int{}
		for pair := range pairs {
			newPairs[string(pair[0])+rules[pair]] += pairs[pair]
			newPairs[rules[pair]+string(pair[1])] += pairs[pair]
		}
		pairs = newPairs
	}

	counter := make(map[string]int)

	for i := range pairs {
		counter[string(i[0])] += pairs[i]
	}
	counter[string(template[len(template)-1])] += 1

	min := math.MaxInt
	max := 0

	for _, v := range counter {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}
