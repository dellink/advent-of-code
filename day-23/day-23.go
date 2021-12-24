package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("day-23/input.txt")
	lines := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		lines = append(lines, line)
	}

	fmt.Printf("Part 1: %d\n", solve(setState(lines), 0, math.MaxInt))

	lines2 := append([]string{}, lines[:3]...)
	lines2 = append(lines2, []string{"  #D#C#B#A#", "  #D#B#A#C#", lines[3], lines[4]}...)

	fmt.Printf("Part 2: %d\n", solve(setState(lines2), 0, math.MaxInt))
}

func setState(lines []string) *State {
	s := State{amphipods: map[Coord]rune{}}
	for y, line := range lines {
		for x, r := range line {
			switch r {
			case 'A', 'B', 'C', 'D':
				s.amphipods[Coord{x, y}] = r
				if y > s.yMax {
					s.yMax = y
				}
			}
		}
	}
	return &s
}

type Coord struct {
	x, y int
}

func (c Coord) InRoom() bool {
	switch c.x {
	case 3, 5, 7, 9:
		return c.y != 1
	}
	return false
}

func (c Coord) IsCorrectRoom(r rune) bool {
	if !c.InRoom() {
		return false
	}
	switch r {
	case 'A':
		return c.x == 3
	case 'B':
		return c.x == 5
	case 'C':
		return c.x == 7
	case 'D':
		return c.x == 9
	}
	return false
}

func (a Coord) Distance(b Coord) int {
	xDist := a.x - b.x
	if xDist < 0 {
		xDist *= -1
	}
	return xDist + (b.y - 1) + (a.y - 1)
}

type Move struct {
	from, to Coord
}

type State struct {
	amphipods map[Coord]rune
	yMax      int
}

func (s *State) Move(m Move) *State {
	res := &State{amphipods: map[Coord]rune{}, yMax: s.yMax}
	r := s.amphipods[m.from]
	res.amphipods[m.to] = r
	for k, v := range s.amphipods {
		if k == m.from {
			continue
		}
		res.amphipods[k] = v
	}
	return res
}

func (s *State) ValidMoves(current, best int) map[Move]int {
	standardMoves := map[Move]int{}
	bestMoves := map[Move]int{}
	for c, r := range s.amphipods {
		valid, foundBest := s.validMovesInner(c, r, current, best)
		if foundBest {
			for k, v := range valid {
				bestMoves[k] = v
			}
			continue
		}
		for k, v := range valid {
			standardMoves[k] = v
		}
	}
	if len(bestMoves) != 0 {
		return bestMoves
	}
	return standardMoves
}

func (s *State) validMovesInner(c Coord, r rune, current, best int) (map[Move]int, bool) {
	if c.IsCorrectRoom(r) && (c.y == s.yMax || func() bool {
		for y := s.yMax; y > c.y; y-- {
			if s.amphipods[Coord{c.x, y}] != r {
				return false
			}
		}
		return true
	}()) {
		return nil, false
	}
	space, d := s.SpaceInRoom(r)
	if space {
		pathable, distance := s.Pathable(c, d)
		if pathable {
			return map[Move]int{{c, d}: int(math.Pow10(int(r-'A'))) * distance}, true
		}
	}
	if !c.InRoom() {
		return nil, false
	}
	Moves := make(map[Move]int)
	multiplier := int(math.Pow10(int(r - 'A')))
	for _, d := range []Coord{
		{1, 1},
		{2, 1},
		{4, 1},
		{6, 1},
		{8, 1},
		{10, 1},
		{11, 1},
	} {
		if c.Distance(d)*multiplier+current > best {
			continue
		}
		if pathable, distance := s.Pathable(c, d); pathable {
			Moves[Move{c, d}] = multiplier * distance
		}
	}
	return Moves, false
}

func (s *State) Pathable(a, b Coord) (bool, int) {
	if _, ok := s.amphipods[b]; ok {
		return false, -1
	}
	switch a.y {
	case 2, 3, 4, 5:
		switch b.y {
		case 2, 3, 4, 5:
			for y := a.y - 1; y >= 1; y-- {
				if _, ok := s.amphipods[Coord{a.x, y}]; ok {
					return false, -1
				}
			}
			if pathable, distance := s.Pathable(Coord{a.x, 1}, b); pathable {
				return true, distance + a.y - 1
			}
			return false, -1
		case 1:
			for y := a.y - 1; y >= 1; y-- {
				if _, ok := s.amphipods[Coord{a.x, y}]; ok {
					return false, -1
				}
			}
			start, end := a.x, b.x
			if end < start {
				end, start = start, end
			}
			for x := start + 1; x <= end; x++ {
				if _, ok := s.amphipods[Coord{x, 1}]; ok {
					return false, -1
				}
			}
			return true, a.y - 1 + end - start
		}
	case 1:
		start, end := a.x, b.x
		if end < start {
			end, start = start, end
		}
		for x := start; x <= end; x++ {
			if x == a.x || x == b.x {
				continue
			}
			if _, ok := s.amphipods[Coord{x, 1}]; ok {
				return false, -1
			}
		}
		return true, end - start + b.y - 1
	}
	return false, -1
}

func (s *State) SpaceInRoom(r rune) (bool, Coord) {
	var x int
	switch r {
	case 'A':
		x = 3
	case 'B':
		x = 5
	case 'C':
		x = 7
	case 'D':
		x = 9
	}
	for y := s.yMax; y > 1; y-- {
		t, has := s.amphipods[Coord{x, y}]
		if has && t != r {
			return false, Coord{}
		}
		if !has {
			return true, Coord{x, y}
		}
	}
	return true, Coord{x, s.yMax}
}

func (s *State) IsOrganized() bool {
	for y := s.yMax; y > 1; y-- {
		if s.amphipods[Coord{3, y}] != 'A' {
			return false
		}
		if s.amphipods[Coord{5, y}] != 'B' {
			return false
		}
		if s.amphipods[Coord{7, y}] != 'C' {
			return false
		}
		if s.amphipods[Coord{9, y}] != 'D' {
			return false
		}
	}
	return true
}

func solve(s *State, score int, bestSolution int) int {
	for move, addScore := range s.ValidMoves(score, bestSolution) {
		if (score + addScore) > bestSolution {
			continue
		}
		bestSolution = solve(s.Move(move), score+addScore, bestSolution)
	}
	if s.IsOrganized() {
		if score < bestSolution {
			bestSolution = score
		}
	}
	return bestSolution
}
