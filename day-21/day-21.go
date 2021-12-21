package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", solve1())
	fmt.Printf("Part 2: %d\n", solve2())
}

type Player struct {
	position, score int
}

func (p Player) move(dice int) Player {
	p.position = (p.position + dice) % 10
	if p.position == 0 {
		p.position = 10
	}
	p.score += p.position
	return p
}

func solve1() int {
	positions := read()

	players := []Player{{positions[0], 0}, {positions[1], 0}}

	rolledCount := 0

	for players[0].score < 1000 && players[1].score < 1000 {
		for i := range players {
			dice := 0

			for j := 0; j < 3; j += 1 {
				rolledCount += 1
				if rolledCount%100 == 0 {
					dice += 100
				} else {
					dice += (rolledCount % 100)
				}
			}

			players[i] = players[i].move(dice)

			if players[i].score >= 1000 {
				break
			}
		}
	}

	return rolledCount * min(players[0].score, players[1].score)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func solve2() int {
	positions := read()

	player1, player2 := game(Player{positions[0], 0}, Player{positions[1], 0}, true, 1)

	if player1 > player2 {
		return player1
	}
	return player2
}

func game(current, other Player, turn bool, universes int) (p1 int, p2 int) {
	if other.score > 20 {
		if turn {
			return 0, universes
		}
		return universes, 0
	}

	for roll, additional := range map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1} {
		p1wins, p2wins := game(other, current.move(roll), !turn, universes*additional)
		p1 += p1wins
		p2 += p2wins
	}

	return p1, p2
}

func read() []int {
	file, _ := os.Open("day-21/input.txt")
	scanner := bufio.NewScanner(file)

	positions := []int{}

	for scanner.Scan() {
		var pos int
		fmt.Sscanf(scanner.Text(), "Player %d starting position: %d", &pos, &pos)
		positions = append(positions, pos)
	}

	return positions
}
