package main

import (
	"maxmil/aoc2025/util"
	"os"
	"strings"
)

func main() {
	util.RunAndCheck(part1, "input-test-part1.txt", 5)
	util.Run(part1, "input.txt")
	util.RunAndCheck(part2, "input-test-part2.txt", 2)
	util.Run(part2, "input.txt")
}

func part1(filename string) int {
	return countPaths( parseInput(filename), Path{Start: "you"}, make(map[Path]struct{}), make(map[Path]int), false)
}

func part2(filename string) int {
	return countPaths( parseInput(filename), Path{Start: "svr"}, make(map[Path]struct{}), make(map[Path]int), true)
}

func countPaths(devices map[string][]string, path Path, seen map[Path]struct{}, found map[Path]int, part2 bool) int {
	if paths, exists := found[path]; exists {
		return paths
	} else if _, exists := seen[path]; exists {
		return 0
	} else if path.Start == "out" {
		if part2 && (!path.Dac || !path.Fft) {
			return 0
		} else {
			return 1
		}
	}
	seen[path] = struct{}{}

	paths := 0
	for _, d := range devices[path.Start] {
		next := Path{d, path.Dac || d == "dac", path.Fft || d == "fft"}
		p := countPaths(devices, next, seen, found, part2)
		if p > 0 {
			found[next] = p
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

type Path struct {
	Start string
	Dac   bool
	Fft   bool
}
