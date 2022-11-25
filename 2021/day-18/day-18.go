package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Snail struct {
	val   int
	depth int
}

func main() {
	numbers := read()

	number := append([]Snail{}, numbers[0]...)
	for i := 1; i < len(numbers); i++ {
		number = reduce(add(number, numbers[i]))
	}
	fmt.Printf("Part 1: %d\n", magnitude(number))

	max := -1
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			val := magnitude(reduce(add(append([]Snail{}, numbers[i]...), append([]Snail{}, numbers[j]...))))
			if val > max {
				max = val
			}
		}
	}
	fmt.Printf("Part 2: %d\n", max)
}

func read() [][]Snail {
	file, _ := os.Open("day-18/input.txt")
	scanner := bufio.NewScanner(file)
	numbers := [][]Snail{}

	for scanner.Scan() {
		number := []Snail{}
		depth := 0
		for _, char := range scanner.Text() {
			switch char {
			case '[':
				depth++
				continue
			case ']':
				depth--
				continue
			case ',':
				continue
			default:
				num, _ := strconv.Atoi(string(char))
				number = append(number, Snail{num, depth})
			}
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func reduce(number []Snail) []Snail {
	reduced := true
	for reduced {
		reduced = false
		for index := range number {
			if number[index].depth >= 5 {
				if index > 0 {
					number[index-1].val += number[index].val
				}
				if index+2 < len(number) {
					number[index+2].val += number[index+1].val
				}

				number[index].depth--
				number[index].val = 0

				number = append(number[:index+1], number[index+2:]...)
				reduced = true

				break
			}
		}

		if reduced {
			continue
		}

		for index := range number {
			if number[index].val > 9 {

				left := number[index].val / 2
				right := number[index].val - left

				snailLeft := Snail{left, number[index].depth + 1}
				snailRight := Snail{right, number[index].depth + 1}

				number = append(number[:index], append([]Snail{snailLeft, snailRight}, number[index+1:]...)...)
				reduced = true

				break
			}
		}
	}

	return number
}

func add(left []Snail, right []Snail) []Snail {
	newNumber := append(left, right...)

	for i := range newNumber {
		newNumber[i].depth++
	}

	return newNumber
}

func magnitude(number []Snail) int {
	for depth := 4; depth > 0; depth -= 1 {
		newNumber := make([]Snail, 0)
		for i := 0; i < len(number); i++ {
			if number[i].depth == depth {
				newNumber = append(newNumber, Snail{3*number[i].val + 2*number[i+1].val, depth - 1})
				i += 1
			} else {
				newNumber = append(newNumber, number[i])
			}
		}
		number = newNumber
	}
	return number[0].val
}
