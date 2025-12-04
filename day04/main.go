package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	fmt.Println("part 1", part1("input.txt"))
	fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string) int {
	rolls := parseInput(filename)
	_, removed := removeRolls(rolls)
	return len(removed)
}

func part2(filename string) int {
	rolls := parseInput(filename)
	count := 0
	for rolls, removed := removeRolls(rolls); len(removed) != 0; rolls, removed = removeRolls(rolls) {
		count += len(removed)
	}

	return count
}

func parseInput(filename string) map[Point]struct{} {
	rolls := make(map[Point]struct{})
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, ch := range scanner.Text() {
			if ch == '@' {
				rolls[Point{x, y}] = struct{}{}
			}
		}
		y++
	}
	return rolls
}

func removeRolls(rolls map[Point]struct{}) (remaining map[Point]struct{}, removed map[Point]struct{}) {
	remaining = make(map[Point]struct{})
	removed = make(map[Point]struct{})

	for p := range rolls {
		count := 0
		for _, n := range p.adjacentRolls(rolls) {
			if _, exists := rolls[n]; exists {
				count++
			}
		}

		if count < 4 {
			removed[p] = struct{}{}
		} else {
			remaining[p] = struct{}{}
		}
	}

	return
}

func (p Point) adjacentRolls(rolls map[Point]struct{}) []Point {
	adj := make([]Point, 0, 8)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx != 0 || dy != 0 {
				p := Point{p.X + dx, p.Y + dy}
				if _, exists := rolls[p]; exists {
					adj = append(adj, p)
				}
			}
		}
	}
	return adj
}
