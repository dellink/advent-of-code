package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	x1, x2, y1, y2 := read()

	highest, shots := solve(x1, x2, y1, y2)

	fmt.Printf("Part 1: %d\n", highest)
	fmt.Printf("Part 2: %d\n", shots)
}

func solve(x1, x2, y1, y2 int) (int, int) {
	highestPeak := 0
	hits := 0

	for vx := 0; vx <= x2; vx++ {
		for vy := y1; vy <= -y1; vy++ {
			peak := check(vx, vy, x1, x2, y1, y2)
			if peak != math.MinInt {
				hits++
			}
			if peak > highestPeak {
				highestPeak = peak
			}
		}
	}

	return highestPeak, hits
}

func check(vx, vy, x1, x2, y1, y2 int) int {
	x, y := 0, 0
	peak := y
	for {
		if y > peak {
			peak = y
		}
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			return peak
		}

		x += vx
		y += vy

		if vx < 0 {
			vx++
		} else if vx > 0 {
			vx--
		}
		vy--

		if vy < 0 && y < y1 || vx == 0 && (x < x1 || x > x2) {
			return math.MinInt
		}
	}
}

func read() (int, int, int, int) {
	input, _ := os.Open("day-17/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Scan()

	var x1, x2, y1, y2 int
	fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	return x1, x2, y1, y2
}
