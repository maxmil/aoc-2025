package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	sections := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	regions := strings.Split(sections[len(sections)-1], "\n")

	shapes := make([]int, 6)
	for i := 0; i < 6; i++ {
		shapes[i] = strings.Count(sections[i], "#")
	}

	fits := 0
	for _, region := range regions {
		parts := strings.Split(region, ":")
		var w, h int
		fmt.Sscanf(parts[0], "%dx%d", &w, &h)
		area := w * h

		max := 0
		for _, s := range strings.Fields(parts[1]) {
			n, _ := strconv.Atoi(s)
			max += 9 * n
		}

		if area >= max {
			fits++
		}
	}

	fmt.Println(fits)
}
