package main

import (
	"fmt"
	"maxmil/aoc2025/util"
	"os"
	"strconv"
	"strings"
)

func main() {
	// util.Run(part1, "input-test.txt", 50)
	// util.Run(part1, "input.txt", 4750297200)
	// print(part1("input.txt"))

	util.Run(part2, "input-test.txt", 24)
	util.Run(part2, "input.txt", 1578115935)
	// fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string) int {
	tiles := parseInput(filename)
	maxSize := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			r := NewRect(tiles[i], tiles[j])
			s := r.size()
			if s > maxSize {
				maxSize = s
			}
		}
	}
	return maxSize
}

func part2(filename string) int {
	tiles := parseInput(filename)

	perimeter := getPermiter(tiles)

	maxSize := 0
	for i := 0; i < len(tiles); i++ {
		fmt.Println("tile ", i)
		for j := i + 1; j < len(tiles); j++ {
			r := NewRect(tiles[i], tiles[j])
			s := r.size()
			if s > maxSize {
				invalid := false
				for x := r.start.X + 1; x < r.end.X; x++ {
					if _, exists := perimeter[Tile{x, r.start.Y + 1}]; exists {
						invalid = true
						break
					}
					if _, exists := perimeter[Tile{x, r.end.Y - 1}]; exists {
						invalid = true
						break
					}
				}
				if !invalid {
					for y := r.start.Y + 1; y < r.end.Y; y++ {
						if _, exists := perimeter[Tile{r.start.X + 1, y}]; exists {
							invalid = true
							break
						}
						if _, exists := perimeter[Tile{r.end.X - 1, y}]; exists {
							invalid = true
							break
						}
					}
					if !invalid {
						fmt.Println(i, j, r, s)
						maxSize = s
					}
				}
			}
		}
	}

	return maxSize
}

func getPermiter(tiles []Tile) map[Tile]struct{} {
	perimiter := make(map[Tile]struct{})
	for i := 0; i < len(tiles); i++ {
		for _, t := range tilesInLine(tiles[i], tiles[(i+1)%len(tiles)]) {
			perimiter[t] = struct{}{}
		}
	}
	return perimiter
}

func tilesInLine(t1, t2 Tile) []Tile {
	tiles := []Tile{}
	if t1.Y == t2.Y {
		for x := 0; x <= max(t1.X, t2.X)-min(t1.X, t2.X); x++ {
			tiles = append(tiles, Tile{min(t1.X, t2.X) + x, t1.Y})
		}
	} else {
		for y := 0; y <= max(t1.Y, t2.Y)-min(t1.Y, t2.Y); y++ {
			tiles = append(tiles, Tile{t1.X, min(t1.Y, t2.Y) + y})
		}
	}

	return tiles
}

func NewRect(t1, t2 Tile) Rect {
	return Rect{
		start: Tile{X: min(t1.X, t2.X), Y: min(t1.Y, t2.Y)},
		end:   Tile{X: max(t1.X, t2.X), Y: max(t1.Y, t2.Y)},
	}
}

func (r Rect) size() int {
	return (r.end.X - r.start.X + 1) * (r.end.Y - r.start.Y + 1)
}

func parseInput(filename string) []Tile {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	tile := make([]Tile, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tile = append(tile, Tile{X: x, Y: y})
	}
	return tile
}

type Rect struct {
	start Tile
	end   Tile
}

type Tile struct {
	X int
	Y int
}

func draw(tiles map[Tile]struct{}) {
	for y := 0; y <= 12; y++ {
		for x := 0; x <= 12; x++ {
			if _, exists := tiles[Tile{x, y}]; exists {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// 438530586 too low
// 1578115935