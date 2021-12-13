package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dots, directions, points := read()

	for i := range directions {
		for dot := range dots {
			var x, y int
			fmt.Sscanf(dot, "%d,%d", &x, &y)

			if directions[i] == 'y' && y >= points[i] {
				y = points[i] - (y - points[i])
				dots[fmt.Sprintf("%d,%d", x, y)] = true
				delete(dots, dot)
			} else if directions[i] == 'x' && x >= points[i] {
				x = points[i] - (x - points[i])
				dots[fmt.Sprintf("%d,%d", x, y)] = true
				delete(dots, dot)
			}
		}

		if i == 0 {
			fmt.Printf("Part 1: %d\n", len(dots))
		}
	}

	fmt.Println("Part 2:")

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if dots[fmt.Sprintf("%d,%d", x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func read() (map[string]bool, []rune, []int) {
	input, _ := os.Open("day-13/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	dots := make(map[string]bool) //A set of dots
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		dots[scanner.Text()] = true
	}

	var directions []rune
	var points []int

	for scanner.Scan() {
		var direction rune
		var point int
		fmt.Sscanf(scanner.Text(), "fold along %c=%d", &direction, &point)

		directions = append(directions, direction)
		points = append(points, point)
	}

	scanner.Scan()

	return dots, directions, points
}
