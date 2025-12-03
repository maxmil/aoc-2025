package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if part1_2("day03/input-test.txt") != 357 {
		log.Fatal("part1 test failed")
	}
	if part1_2("day03/input.txt") != 17087 {
		log.Fatal("part1 failed")
	}
	if part2("day03/input-test.txt") != 3121910778619 {
		log.Fatal("part2 test failed")
	}
	fmt.Println()
	fmt.Println(part2("day03/input.txt"))
}

func part1(filename string) int {
	content, _ := os.ReadFile(filename)
	sum := 0
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		runes := []rune(line)
		max := []int{-1, -1}
		for i := 0; i < len(runes); i++ {
			if max[0] == -1 || (i < len(runes)-1 && runes[i] > runes[max[0]]) {
				max[0] = i
				max[1] = -1
			} else if max[1] == -1 || runes[i] > runes[max[1]] {
				max[1] = i
			}
		}
		top, _ := strconv.Atoi(string(runes[max[0]]) + string(runes[max[1]]))
		sum += top
	}
	return sum
}

func part1_2(filename string) int {
	return maxJoltage(filename, 2)
}

func part2(filename string) int {
	return maxJoltage(filename, 12)
}

func maxJoltage(filename string, batteries int) int {
	content, _ := os.ReadFile(filename)
	sum := 0
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		runes := []rune(line)

		inds := make([]int, batteries)
		for i := range inds {
			inds[i] = -1
		}

		for i := 0; i < len(runes); i++ {
			for j := max(0, len(inds)-len(runes)+i); j < len(inds); j++ {
				// fmt.Println(i, j, string(runes[i]))
				if inds[j] == -1 || runes[i] > runes[inds[j]] {
					inds[j] = i
					for k := j + 1; k < len(inds); k++ {
						inds[k] = -1
					}
					break
				}
			}
			// fmt.Println(getBateries(inds, runes))
		}
		bateries := getBateries(inds, runes)
		// fmt.Println(bateries)
		power, _ := strconv.Atoi(bateries)
		sum += power
	}

	return sum
}

func getBateries(inds []int, runes []rune) string {
	bateries := ""
	for _, ind := range inds {
		if ind == -1 {
			bateries += "_"
			continue
		}
		bateries += string(runes[ind])
	}
	return bateries
}


//169298745606378 to high
//169019504359949