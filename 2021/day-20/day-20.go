package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	i, j int
}

func main() {
	algorithm, image := read()

	fmt.Printf("Part 1: %d\n", enhanceImage(image, algorithm, 2))
	fmt.Printf("Part 2: %d\n", enhanceImage(image, algorithm, 50))
}

func enhanceImage(image map[Position]string, algorithm string, steps int) int {
	for step := 0; step < steps; step++ {
		pixel := "."
		if step%2 != 0 {
			pixel = string(algorithm[0])
		}
		image = enhance(add(image, pixel), algorithm, pixel)
	}
	counter := 0
	for _, v := range image {
		if v == "#" {
			counter += 1
		}
	}
	return counter
}

func add(image map[Position]string, pixel string) map[Position]string {
	newImage := map[Position]string{}

	for i, p := range image {
		newImage[i] = p
	}

	minI, maxI, minJ, maxJ := 0, 0, 0, 0

	for t := range newImage {
		minI = min(minI, t.i)
		minJ = min(minJ, t.j)
		maxI = max(maxI, t.i)
		maxJ = max(maxJ, t.j)
	}

	for j := minJ; j <= maxJ; j += 1 {
		newImage[Position{minI - 1, j}] = pixel
		newImage[Position{maxI + 1, j}] = pixel
	}

	for i := minI; i <= maxI; i += 1 {
		newImage[Position{i, minJ - 1}] = pixel
		newImage[Position{i, maxJ + 1}] = pixel
	}

	newImage[Position{minI - 1, minJ - 1}] = pixel
	newImage[Position{minI - 1, maxJ + 1}] = pixel
	newImage[Position{maxI + 1, minJ - 1}] = pixel
	newImage[Position{maxI + 1, maxJ + 1}] = pixel
	return newImage
}

func enhance(image map[Position]string, algorithm, pixel string) map[Position]string {
	enhancedImage := map[Position]string{}

	keys := []Position{}
	for k := range image {
		keys = append(keys, k)
	}

	for _, k := range keys {
		arr := generateRow(pixel, k, image, []Position{{-1, -1}, {-1, 0}, {-1, 1}})
		arr = append(arr, generateRow(pixel, k, image, []Position{{0, -1}, {0, 0}, {0, 1}})...)
		arr = append(arr, generateRow(pixel, k, image, []Position{{1, -1}, {1, 0}, {1, 1}})...)

		bits := ""
		for _, s := range arr {
			if s == "#" {
				bits += "1"
			} else {
				bits += "0"
			}
		}

		index, _ := strconv.ParseInt(bits, 2, 64)
		enhancedImage[k] = string(algorithm[index])
	}
	return enhancedImage
}

func generateRow(pixel string, t Position, image map[Position]string, directions []Position) []string {
	row := []string{}
	for _, dir := range directions {
		newP := Position{
			i: t.i + dir.i,
			j: t.j + dir.j,
		}
		if _, ok := image[newP]; !ok {
			row = append(row, pixel)
		} else {
			row = append(row, image[newP])
		}
	}
	return row
}

func read() (string, map[Position]string) {
	file, _ := os.Open("day-20/input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	algorithm := scanner.Text()

	image := map[Position]string{}

	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for y, s := range line {
			image[Position{x, y}] = string(s)
		}
		x += 1

	}
	return algorithm, image
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
