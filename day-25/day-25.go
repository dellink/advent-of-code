package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, _ := os.Open("day-25/input.txt")
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	width := len(lines[0])
	height := len(lines)

	cucumbers := map[Point]rune{}

	for j, line := range lines {
		for i, c := range line {
			if c != '.' {
				cucumbers[Point{X: i, Y: j}] = c
			}
		}
	}

	steps := 0
	var newMap map[Point]rune

	for {
		moved := false
		newMap = map[Point]rune{}

		for p, c := range cucumbers {
			if c == '>' {
				destination := Point{X: p.X + 1, Y: p.Y}
				if destination.X >= width {
					destination.X = 0
				}
				if cucumbers[destination] == 0 {
					newMap[destination] = c
					moved = true
				} else {
					newMap[p] = c
				}
			} else {
				newMap[p] = c
			}
		}
		cucumbers = newMap

		newMap = make(map[Point]rune)

		for p, c := range cucumbers {
			if c == 'v' {
				destination := Point{X: p.X, Y: p.Y + 1}
				if destination.Y >= height {
					destination.Y = 0
				}
				if cucumbers[destination] == 0 {
					newMap[destination] = c
					moved = true
				} else {
					newMap[p] = c
				}
			} else {
				newMap[p] = c
			}
		}
		cucumbers = newMap
		steps++

		if moved == false {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", steps)
}
