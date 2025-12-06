package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part 1", part1("input.txt"))
	fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string) int {
	ws := parseInput(filename, strings.Fields)
	total := 0

	for col := range ws[0] {
		values := make([]int, len(ws)-1)
		for row := range values {
			values[row], _ = strconv.Atoi(ws[row][col])
		}

		if ws[len(ws)-1][col] == "*" {
			total += mult(values)
		} else {
			total += sum(values)
		}
	}
	return total
}

func part2(filename string) int {
	ws := parseInput(filename, func(s string) []string { return strings.Split(s, "") })
	total, values := 0, []int{}
	var op func([]int) int

	for col := range ws[0] {
		switch ws[len(ws)-1][col] {
		case "*":
			op = mult
		case "+":
			op = sum
		}

		value := ""
		for row := 0; row < len(ws)-1; row++ {
			if ws[row][col] != " " {
				value += ws[row][col]
			}
		}

		if value != "" {
			v, _ := strconv.Atoi(value)
			values = append(values, v)
		}

		if value == "" || col == len(ws[0])-1 {
			total += op(values)
			values = []int{}
		}
	}
	return total
}

func parseInput(filename string, splitter func(string) []string) [][]string {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	worksheet := make([][]string, len(lines))
	for i, line := range lines {
		worksheet[i] = splitter(line)
	}
	return worksheet
}

func mult(values []int) int {
	result := 1
	for _, v := range values {
		if v != 0 {
			result *= v
		}
	}
	return result
}

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
