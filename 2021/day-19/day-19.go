package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	x, y, z int
}

func main() {
	lines := read()

	found := map[int][]Coordinate{0: lines[0]}

	skip := map[string]bool{}
	beacons := map[Coordinate]bool{}
	scanners := []Coordinate{}

	for _, p := range lines[0] {
		beacons[p] = true
	}

	for len(found) < len(lines) {
		for i, scan := range lines {
			if _, ok := found[i]; ok {
				continue
			}

			for j := range found {
				if _, ok := skip[strconv.Itoa(i)+"."+strconv.Itoa(j)]; ok {
					continue
				}

				if ok, positions, scanner := findMatch(found[j], scan); ok {
					scanners = append(scanners, scanner)
					found[i] = positions
					for _, m := range positions {
						beacons[m] = true
					}
					break
				}
				skip[strconv.Itoa(i)+"."+strconv.Itoa(j)] = true
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(beacons))

	max := 0.0
	for _, scan1 := range scanners {
		for _, scan2 := range scanners {
			distance := math.Abs(float64(scan1.x-scan2.x)) + math.Abs(float64(scan1.y-scan2.y)) + math.Abs(float64(scan1.z-scan2.z))
			if distance > max {
				max = distance
			}
		}
	}
	fmt.Printf("Part 2: %d\n", int(max))
}

func findMatch(scanner1, scanner2 []Coordinate) (bool, []Coordinate, Coordinate) {
	inScanner1 := map[Coordinate]bool{}
	for _, x := range scanner1 {
		inScanner1[x] = true
	}

	position := []struct{ x, y, z int }{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}

	invert := []struct{ x, y, z int }{
		{1, 1, 1},
		{1, 1, -1},
		{1, -1, 1},
		{1, -1, -1},
		{-1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{-1, -1, -1},
	}

	for _, position := range position {
		for _, invert := range invert {

			scanner2NewCoordinates := make([]Coordinate, len(scanner2))
			for i := 0; i < len(scanner2); i++ {
				coords := []int{scanner2[i].x, scanner2[i].y, scanner2[i].z}
				scanner2NewCoordinates[i] = Coordinate{
					x: coords[position.x] * invert.x,
					y: coords[position.y] * invert.y,
					z: coords[position.z] * invert.z,
				}
			}

			for _, cScanner1 := range scanner1 {
				for _, cScanner2 := range scanner2NewCoordinates {
					scanner2Coordinate := Coordinate{
						x: cScanner2.x - cScanner1.x,
						y: cScanner2.y - cScanner1.y,
						z: cScanner2.z - cScanner1.z,
					}

					matches := 0
					relativePositions := make([]Coordinate, 0)
					for _, otherCoordScanner2 := range scanner2NewCoordinates {
						relativeToScanner1 := Coordinate{
							x: otherCoordScanner2.x - scanner2Coordinate.x,
							y: otherCoordScanner2.y - scanner2Coordinate.y,
							z: otherCoordScanner2.z - scanner2Coordinate.z,
						}
						if _, ok := inScanner1[relativeToScanner1]; ok {
							matches++
						}
						relativePositions = append(relativePositions, relativeToScanner1)
					}

					if matches >= 12 {
						return true, relativePositions, scanner2Coordinate
					}
				}
			}
		}
	}
	return false, nil, Coordinate{}
}

func read() [][]Coordinate {
	file, _ := os.Open("day-19/input.txt")
	scanner := bufio.NewScanner(file)

	scans := [][]Coordinate{}
	scan := []Coordinate{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			scans = append(scans, scan)
			scan = []Coordinate{}
			continue
		}
		var x, y, z int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		if err == nil {
			scan = append(scan, Coordinate{x, y, z})
		}
	}
	scans = append(scans, scan)

	return scans
}
