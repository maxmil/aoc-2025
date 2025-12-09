package main

import (
	"maxmil/aoc2025/util"
	"strconv"
	"strings"
)

func main() {
	util.RunAndCheck(part1, "input-test.txt", 1227775554)
	util.RunAndCheck(part1, "input.txt", 12586854255)
	util.RunAndCheck(part2, "input-test.txt", 4174379265)
	util.RunAndCheck(part2, "input.txt", 17298174201)
}

func part1(filename string) int {
	content := util.ReadContent(filename)
	sum := 0
	for _, r := range strings.Split(content, ",") {
		start, end, _ := strings.Cut(r, "-")
		for _, id := range findInvalidIds(start, end, 2) {
			sum += id
		}
	}
	return sum
}

func part2(filename string) int {
	content := util.ReadContent(filename)
	sum := 0
	for _, r := range strings.Split(content, ",") {
		start, end, _ := strings.Cut(r, "-")
		uniqueIds := make(map[int]struct{})
		for repetitions := 2; repetitions <= len(end); repetitions++ {
			for _, id := range findInvalidIds(start, end, repetitions) {
				uniqueIds[id] = struct{}{}
			}
		}

		for k := range uniqueIds {
			sum += k
		}
	}

	return sum
}

func findInvalidIds(start string, end string, repetitions int) []int {
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	invalid := []int{}
	base, _ := strconv.Atoi(start[:len(start)/repetitions])
	i := 0
	for {
		next, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(base+i), repetitions))
		if next < startInt {
			i++
			continue
		}
		if next > endInt {
			break
		}
		invalid = append(invalid, next)
		i++
	}
	return invalid
}
