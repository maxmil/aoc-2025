package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	// t := part1("input-test.txt")
	// if t != 13 {
	// 	log.Fatalf("Failed part 1 test %d != 13", t)
	// }

	t2 := part2("input-test.txt")
	if t2 != 43 {
		log.Fatalf("Failed part 2 test %d != 42", t2)
	}
	// fmt.Println("part 1", part1("input.txt"))
	fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string) int {
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

	accessible := []Point{}
	for p := range rolls {
		adj := []Point{}
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for x := p.X - 1; x <= p.X+1; x++ {
				if x != p.X || y != p.Y {
					if _, isRoll := rolls[Point{x, y}]; isRoll {
						adj = append(adj, Point{x, y})
					}
				}
			}
		}
		// fmt.Println(p, rolls)
		if len(adj) < 4 {
			accessible = append(accessible, p)
			fmt.Println(p, adj)
		}
	}

	return len(accessible)
}

func part2(filename string) int {
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

	removed := 0
	for r := removeRolls(rolls); r != 0; r = removeRolls(rolls) {
		removed += r
		fmt.Printf("Removed %d rolls\n", r)
		printGrid(rolls, 10, 10)
		fmt.Println()
	}

	return removed
}

func removeRolls(rolls map[Point]struct{}) int {
	removed := 0
	for p := range rolls {
		if _, present := rolls[p]; !present {
			continue
		}

		adj := []Point{}
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for x := p.X - 1; x <= p.X+1; x++ {
				if x != p.X || y != p.Y {
					if _, isRoll := rolls[Point{x, y}]; isRoll {
						adj = append(adj, Point{x, y})
					}
				}
			}
		}
		// fmt.Println(p, rolls)
		if len(adj) < 4 {
			delete(rolls, p)
			removed++
		}
	}

	return removed
}

func printGrid(rolls map[Point]struct{}, width, height int) {
      for y := 0; y < height; y++ {
          for x := 0; x < width; x++ {
              if _, exists := rolls[Point{x, y}]; exists {
                  fmt.Print("@")
              } else {
                  fmt.Print(".")
              }
          }
          fmt.Println()
      }
  }

// 8436 too low
