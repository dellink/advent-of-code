package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y, z int
}

type Instruction struct {
	state bool
	start Point
	end   Point
}

func (this Instruction) volume() int {
	return (abs(this.end.x-this.start.x) + 1) * (abs(this.end.y-this.start.y) + 1) * (abs(this.end.z-this.start.z) + 1)
}

func (this Instruction) isValid() bool {
	return this.start.x <= this.end.x && this.start.y <= this.end.y && this.start.z <= this.end.z
}

func (this Instruction) intersect(that Instruction, state bool) Instruction {
	return Instruction{
		state,
		Point{max(this.start.x, that.start.x), max(this.start.y, that.start.y), max(this.start.z, that.start.z)},
		Point{min(this.end.x, that.end.x), min(this.end.y, that.end.y), min(this.end.z, that.end.z)},
	}
}

func main() {
	instructions1, instructions2 := read()

	fmt.Printf("Part 1: %d\n", solve1(instructions1))
	fmt.Printf("Part 2: %d\n", solve2(instructions2))
}

func solve2(instructions []Instruction) int {
	list := []Instruction{}

	for _, c1 := range instructions {
		add := []Instruction{}
		if c1.state {
			add = append(add, c1)
		}
		for _, c2 := range list {
			if intersection := c1.intersect(c2, !c2.state); intersection.isValid() {
				add = append(add, intersection)
			}
		}
		list = append(list, add...)
	}

	volume := 0
	for _, c := range list {
		if c.state {
			volume += c.volume()
		} else {
			volume -= c.volume()
		}
	}
	return volume
}

func solve1(instructions []Instruction) int {
	cubes := map[Point]bool{}

	for _, instruction := range instructions {
		for i := instruction.start.x; i <= instruction.end.x; i += 1 {
			for j := instruction.start.y; j <= instruction.end.y; j += 1 {
				for k := instruction.start.z; k <= instruction.end.z; k += 1 {
					cubes[Point{i, j, k}] = instruction.state
				}
			}
		}
	}

	count := 0

	for _, on := range cubes {
		if on {
			count += 1
		}
	}

	return count
}

func read() ([]Instruction, []Instruction) {
	file, _ := os.Open("day-22/input.txt")
	scanner := bufio.NewScanner(file)

	instructions1 := []Instruction{}
	instructions2 := []Instruction{}

	var command string
	var x1, x2, y1, y2, z1, z2 int

	for scanner.Scan() {

		fmt.Sscanf(scanner.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &x1, &x2, &y1, &y2, &z1, &z2)

		state := false
		if command == "on" {
			state = true
		}

		instructions1 = append(instructions1, Instruction{state, Point{max(x1, -50), max(-50, y1), max(-50, z1)}, Point{min(x2, 50), min(50, y2), min(50, z2)}})
		instructions2 = append(instructions2, Instruction{state, Point{x1, y1, z1}, Point{x2, y2, z2}})
	}

	return instructions1, instructions2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
