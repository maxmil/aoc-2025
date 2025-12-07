package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	t1 := part1("input-test.txt")
	if t1 != 21 {
		log.Fatalf("part1 test failed with %d", t1)
	}
	fmt.Println(part1("input.txt"))

	t2 := part2("input-test.txt")
	if t2 != 40 {
		log.Fatalf("part2 test failed with %d", t1)
	}
	fmt.Println(part2("input.txt"))
}

func part1(filename string) int {
	diagram, start, height, width := parseInput(filename)

	beams := map[Point]struct{}{start: {}}
	splits := 0
	for y := 0; y < height; y++ {
		newBeams := map[Point]struct{}{}
		for b := range beams {
			if _, exists := diagram[Point{b.X, b.Y + 1}]; exists {
				if b.X > 0 {
					newBeams[Point{b.X - 1, b.Y + 1}] = struct{}{}
				}
				if b.X < width-1 {
					newBeams[Point{b.X + 1, b.Y + 1}] = struct{}{}
				}
				splits++
			} else {
				newBeams[Point{b.X, b.Y + 1}] = struct{}{}
			}
		}
		beams = newBeams
	}

	return splits
}

func part2(filename string) int {
	diagram, start, height, width := parseInput(filename)

	beams := map[Point]int{start: 1}
	for y := 0; y < height; y++ {
		newBeams := map[Point]int{}
		for b := range beams {
			if _, exists := diagram[Point{b.X, b.Y + 1}]; exists {
				if b.X > 0 {
					count, exists := newBeams[Point{b.X - 1, b.Y + 1}]
					if !exists {
						count = 0
					}
					newBeams[Point{b.X - 1, b.Y + 1}] = count + beams[b]
				}
				if b.X < width-1 {
					count, exists := newBeams[Point{b.X + 1, b.Y + 1}]
					if !exists {
						count = 0
					}
					newBeams[Point{b.X + 1, b.Y + 1}] = count + beams[b]
				}
			} else {
				count, exists := newBeams[Point{b.X, b.Y + 1}]
				if !exists {
					count = 0
				}
				newBeams[Point{b.X, b.Y + 1}] = count + beams[b]
			}
		}
		beams = newBeams
	}

	timelines := 0
	for _, count := range beams {
		timelines += count
	}
	return timelines
}

type Point struct {
	X, Y int
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

// 3080 too low
