package main

import (
	"fmt"
	"maxmil/aoc2025/util"
	"strconv"
)

func main() {
	lines := util.ReadLines("input.txt")
	position := 50
	zeros := 0
	passes := 0
	dirs := map[string]int{"L": -1, "R": 1}
	for _, line := range lines {
		dir := string(line[0])
		num, _ := strconv.Atoi(line[1:])
		next := position + num*dirs[dir]
		if next%100 == 0 {
			zeros++
		}

		passes += dirs[dir] * ((next - dirs[dir] * 1) / 100)
		if next < 0 && position > 0 {
			passes += 1
		}
		position = (next + (100 * (abs(next)/100 + 1))) % 100
	}
	fmt.Println(zeros, passes, zeros+passes)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
