package main

import (
	"maxmil/aoc2025/util"
	"os"
	"strconv"
	"strings"
)

func main() {
	util.Run(part1, "input.txt")
	util.Run(part2, "input.txt")
}

func part1(filename string) int {
	tiles := parseInput(filename)
	maxSize := 0
	for i := range len(tiles) {
		for j := i + 1; j < len(tiles); j++ {
			r := newRect(tiles[i], tiles[j])
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

	perimeter := make([]Line, 0, len(tiles))
	for t := range len(tiles) {
		perimeter = append(perimeter, newLine(tiles[t], tiles[(t+1)%len(tiles)]))
	}

	maxSize := 0
	for i := range len(tiles) {
		for j := i + 1; j < len(tiles); j++ {
			rect := newRect(tiles[i], tiles[j])
			size := rect.size()
			if size > maxSize && !rect.crossesPerimeter(perimeter) {
				maxSize = size
			}
		}
	}

	return maxSize
}

func parseInput(filename string) []Tile {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	tiles := make([]Tile, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, Tile{X: x, Y: y})
	}
	return tiles
}

func newRect(t1, t2 Tile) Rect {
	return Rect{
		start: Tile{X: min(t1.X, t2.X), Y: min(t1.Y, t2.Y)},
		end:   Tile{X: max(t1.X, t2.X), Y: max(t1.Y, t2.Y)},
	}
}

func (r Rect) size() int {
	return (r.end.X - r.start.X + 1) * (r.end.Y - r.start.Y + 1)
}

func (r Rect) crossesPerimeter(perimeter []Line) bool {
	innerPerimeter := []Line{
		{Tile{r.start.X + 1, r.start.Y + 1}, Tile{r.end.X - 1, r.start.Y + 1}},
		{Tile{r.end.X - 1, r.start.Y + 1}, Tile{r.end.X - 1, r.end.Y - 1}},
		{Tile{r.start.X + 1, r.end.Y - 1}, Tile{r.end.X - 1, r.end.Y - 1}},
		{Tile{r.start.X + 1, r.start.Y + 1}, Tile{r.start.X + 1, r.end.Y - 1}},
	}
	for _, p := range perimeter {
		for _, l := range innerPerimeter {
			if l.intersects(p) {
				return true
			}
		}
	}
	return false
}

func newLine(t1, t2 Tile) Line {
	if t1.X < t2.X || (t1.X == t2.X && t1.Y < t2.Y) {
		return Line{t1, t2}
	}
	return Line{t2, t1}
}

func (l Line) intersects(other Line) bool {
	if l.start.Y == l.end.Y && other.start.Y == other.end.Y {
		if l.start.Y != other.start.Y {
			return false
		}
		return l.start.X <= other.end.X && other.start.X <= l.end.X
	}

	if l.start.X == l.end.X && other.start.X == other.end.X {
		if l.start.X != other.start.X {
			return false
		}
		return l.start.Y <= other.end.Y && other.start.Y <= l.end.Y
	}

	var h, v Line
	if l.start.Y == l.end.Y {
		h, v = l, other
	} else {
		h, v = other, l
	}

	return v.start.X >= h.start.X && v.start.X <= h.end.X && h.start.Y >= v.start.Y && h.start.Y <= v.end.Y
}

type Tile struct {
	X, Y int
}

type Line struct {
	start, end Tile
}

type Rect struct {
	start, end Tile
}
