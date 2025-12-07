package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	timelines, splits := run("input.txt")

	fmt.Println("part 1", splits)

	count := 0
	for _, c := range timelines {
		count += c
	}
	fmt.Println("part 2", count)
}

type Point struct {
	X, Y int
}

func run(filename string) (timelines map[Point]int, splits int) {
	diagram, start, height, width := parseInput(filename)
	timelines = map[Point]int{start: 1}
	splits = 0
	for y := 0; y < height; y++ {
		next := map[Point]int{}
		for b := range timelines {
			if _, exists := diagram[Point{b.X, b.Y + 1}]; exists {
				splits++
				if b.X > 0 {
					next[Point{b.X - 1, b.Y + 1}] += timelines[b]
				}
				if b.X < width-1 {
					next[Point{b.X + 1, b.Y + 1}] += timelines[b]
				}
			} else {
				next[Point{b.X, b.Y + 1}] += timelines[b]
			}
		}
		timelines = next
	}
	return
}

func parseInput(filename string) (diagram map[Point]struct{}, start Point, height int, width int) {
	diagram = make(map[Point]struct{})
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for x, ch := range line {
			switch ch {
			case '^':
				diagram[Point{x, y}] = struct{}{}
			case 'S':
				start = Point{x, y}
			}
		}
		y++
	}
	height = y
	return
}
