package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part 1", part1("input.txt"))
	fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string) int {
	ranges, ingredients := parseInput(filename)
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

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if last.End < r.Start {
			merged = append(merged, r)
		} else {
			last.End = max(last.End, r.End)
		}
	}

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
