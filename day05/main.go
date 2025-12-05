package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// if part1("input-test.txt") != 3 {
	// 	panic("part 1 test failed")
	// }
	// fmt.Println("part 1", part1("input.txt"))

	t2 := part2("input-test.txt")
	if t2 != 14 {
		panic("part 2 test failed " + strconv.Itoa(t2))
	}
	fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string) int {
	ranges, ingredients := parseInput(filename)
	fmt.Println(ranges, ingredients)
	count := 0
	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r.Start && ingredient <= r.End {
				count++
				break
			}
		}
	}
	return count
}

func part2(filename string) int {
	ranges, _ := parseInput(filename)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	merged := []Range{}
	curr := ranges[0]
	for _, r := range ranges[1:] {
		if curr.End < r.Start {
			merged = append(merged, curr)
			curr = r
		} else {
			curr.End = max(curr.End, r.End)
		}
	}
	merged = append(merged, curr)

	count := 0
	for _, r := range merged {
		count += r.End - r.Start + 1
	}

	return count
}

type Range struct {
	Start, End int
}

func parseInput(filename string) (ranges []Range, ingredients []int) {
	contents, _ := os.ReadFile(filename)
	parts := strings.Split(string(contents), "\n\n")

	ranges = []Range{}
	for _, line := range strings.Split(parts[0], "\n") {
		var r Range
		fmt.Sscanf(line, "%d-%d", &r.Start, &r.End)
		ranges = append(ranges, r)
	}

	ingredients = []int{}
	for _, line := range strings.Split(parts[1], "\n") {
		ingredient, _ := strconv.Atoi(line)
		ingredients = append(ingredients, ingredient)
	}

	return
}
