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

		on := make([]int, turnOn)
		for i := range on {
			on[i] = -1
		}

		for b := range batteries {
			for i := max(0, len(on)-len(batteries)+b); i < len(on); i++ {
				if on[i] == -1 || batteries[b] > batteries[on[i]] {
					on[i] = b
					for j := i + 1; j < len(on); j++ {
						on[j] = -1
					}
					break
				}
			}
		}
		bateries := ""
		for _, ind := range on {
			bateries += string(batteries[ind])
		}
		power, _ := strconv.Atoi(bateries)
		sum += power
	}

	return sum
}