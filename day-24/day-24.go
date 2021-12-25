package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	a, b, c int
}

func main() {
	lines := read()

	fmt.Printf("Part 1: %d\n", solve(lines, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Printf("Part 2: %d\n", solve(lines, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func solve(instructions [14]Instruction, order []int) int {
	res := [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	stack := []int{}
	functions := [14](func(int) int){}
	indexes := []int{}
	for i, instruction := range instructions {
		switch instruction.a {
		case 1:
			c := instruction.c
			fn := func(n int) int { return n + c }
			functions[i] = fn
			stack = append(stack, i)
		case 26:
			fnIdx := stack[len(stack)-1]
			currentIdx := i
			b := instruction.b
			setFunc := func(n int) int {
				res[fnIdx] = n
				res[currentIdx] = functions[fnIdx](n) + b
				return 0
			}
			indexes = append(indexes, currentIdx)
			stack = stack[:len(stack)-1]
			functions[i] = setFunc
		}
	}

	for _, index := range indexes {
		for _, n := range order {
			functions[index](n)
			if isValid(res) {
				break
			}
		}
	}

	return toInt(res)
}

func isValid(data [14]int) bool {
	for _, v := range data {
		if v < 1 || v > 9 {
			return false
		}
	}
	return true
}

func toInt(data [14]int) int {
	res := 0
	for _, v := range data {
		res *= 10
		res += v
	}
	return res
}

func read() [14]Instruction {
	file, _ := os.Open("day-24/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := [14]Instruction{}

	var a, b string
	var c int

	for i := 0; i < 14; i++ {
		instruction := Instruction{}
		for j := 0; j < 18; j++ {
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%s %s %d", &a, &b, &c)
			switch j {
			case 4:
				instruction.a = c
			case 5:
				instruction.b = c
			case 15:
				instruction.c = c
			}
		}
		instructions[i] = instruction
	}
	return instructions
}
