package main

import (
	"maxmil/aoc2025/util"
	"os"
	"strings"
)

func main() {
	// util.RunAndCheck(part1, "input-test-part1.txt",5)
	// util.Run(part1, "input.txt")
	util.RunAndCheck(part2, "input-test-part2.txt", 2)
	util.Run(part2, "input.txt")
}

func part1(filename string) int {
	devices := parseInput(filename)

	// start with you
	// countPaths - keep track of seen devices and devices that lead to out
		// if device leads to out, increment
		// if seen then stop
		// if device is out, increment and add
		// else recurse and add to paths
		// add to solutions if it's a solution


	found := make(map[string]int)
	paths := countPaths(devices, "you", make(map[string]struct{}), found)
	return paths
}

func part2(filename string) int {
	devices := parseInput(filename)
	found := make(map[Path]int)
	start := Path{"svr", false, false}
	paths := countPaths2(devices, start, make(map[Path]struct{}), found)
	return paths
}

func countPaths2(devices map[string][]string, path Path, seen map[Path]struct{}, found map[Path]int) int {
	if paths, exists := found[path]; exists {
		return paths
	} else if _, exists := seen[path]; exists {
		return 0
	} else if path.Start == "out" {
		if (path.Dac && path.Fft) {
			return 1
		} else {
			return 0
		}
	}
	seen[path] = struct{}{}

	paths := 0
	for _, d := range devices[path.Start] {
		next := Path{d, path.Dac || d == "dac", path.Fft || d == "fft"}
		p := countPaths2(devices, next, seen, found)
		if (p > 0) {
			found[next] = p
		}
		paths += p
	}
	return paths
}

type Path struct {
	Start string
	Dac bool
	Fft bool
}

func countPaths(devices map[string][]string, start string, seen map[string]struct{}, found map[string]int) int {
	if paths, exists := found[start]; exists {
		return paths
	} else if _, exists := seen[start]; exists {
		return 0
	} else if start == "out" {
		return 1
	}
	seen[start] = struct{}{}

	paths := 0
	for _, d := range devices[start] {
		p := countPaths(devices, d, seen, found)
		if (p > 0) {
			found[d] = p
		}
		paths += p
	}
	return paths
}

func parseInput(filename string) map[string][]string {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	devices := make(map[string][]string, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ":")
		devices[parts[0]] = strings.Fields(parts[1])
	}

	return devices
}