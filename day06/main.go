package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// t1 := part1("input-test.txt")
	// if t1 != 4277556 {
	// 	log.Fatalf("Failed part 1 test %d != ", t1)
	// }
	// fmt.Println("part 1", part1("input.txt"))

	t2 := part2("input-test.txt")
	if t2 != 3263827 {
		log.Fatalf("Failed part 2 test %d != 3263827", t2)
	}
	fmt.Println("part 1", part2("input.txt"))
}

func part1(filename string) int {
	contents, _ := os.ReadFile(filename)
	worksheet := [][]string{}
	for _, line := range strings.Split(string(contents), "\n") {
		worksheet = append(worksheet, strings.Fields(line))
	}

	total := 0
	for col := 0; col < len(worksheet[0]); col++ {
		values := []int{}
		for row := 0; row < len(worksheet)-1; row++ {
			val, _ := strconv.Atoi(worksheet[row][col])
			values = append(values, val)
		}

		var result int
		if worksheet[len(worksheet)-1][col] == "*" {
			result = 1
			for _, v := range values {
				result = result * v
			}
		} else {
			result = 0
			for _, v := range values {
				result = result + v
			}
		}
		fmt.Printf("Column %d result: %d\n", col+1, result)
		total += result
	}

	return total
}

func part2(filename string) int {
	contents, _ := os.ReadFile(filename)
	worksheet := [][]string{}
	for _, line := range strings.Split(string(contents), "\n") {
		worksheet = append(worksheet, strings.Split(line, ""))
	}

	total := 0
	var calculate func([]int) int
	values := []int{}
	for col := 0; col < len(worksheet[0]); col++ {
		value := ""
		for row := 0; row < len(worksheet)-1; row++ {
			if worksheet[row][col] != " " {
				value += worksheet[row][col]
			}
		}

		if len(value) > 0 {
			i, _ := strconv.Atoi(value)
			values = append(values, i)
		} else {
			result := calculate(values)
			values = []int{}
			total += result
			fmt.Printf("Column %d result: %d values: %v\n", col+1, result, values)
		}

		if worksheet[len(worksheet)-1][col] == "*" {
			calculate = func(values []int) int {
				result := 1
				for _, v := range values {
					if v != 0 {
						result *= v
					}
					// log.Println(v, result)
				}
				return result
			}
		} else if worksheet[len(worksheet)-1][col] == "+" {
			calculate = func(values []int) int {
				result := 0
				for _, v := range values {
					result += v
				}
				return result
			}
		}
	}

	result := calculate(values)
	values = []int{}
	total += result
	// fmt.Printf("Column %d result: %d values: %v\n", col+1, result, values)

	return total
}
