package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part 1", maxJoltage(2))
	fmt.Println("part 2", maxJoltage(12))
}

func maxJoltage(turnOn int) int {
	content, _ := os.ReadFile("input.txt")
	sum := 0
	for _, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		batteries := []rune(line)

		inds := make([]int, turnOn)
		for i := range inds {
			inds[i] = -1
		}

		for b := range batteries {
			for i := max(0, len(inds)-len(batteries)+b); i < len(inds); i++ {
				if inds[i] == -1 || batteries[b] > batteries[inds[i]] {
					inds[i] = b
					for j := i + 1; j < len(inds); j++ {
						inds[j] = -1
					}
					break
				}
			}
		}
		bateries := ""
		for _, ind := range inds {
			bateries += string(batteries[ind])
		}
		power, _ := strconv.Atoi(bateries)
		sum += power
	}

	return sum
}