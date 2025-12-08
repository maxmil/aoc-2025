package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	t1 := part1("input-test.txt", 10)
	if t1 != 40 {
		log.Fatalf("Part 1 test failed with %d != 40", t1)
	}
	p1 := part1("input.txt", 1000)
	if p1 != 163548 {
		log.Fatalf("Part 1 test failed with %d != 163548", p1)
	}
	// fmt.Println("part 1", part1("input.txt", 1000))

	t2 := part2("input-test.txt")
	if t2 != 25272 {
		log.Fatalf("Part 2 test failed with %d != 25272", t2)
	}
	fmt.Println("part 2", part2("input.txt"))

}

func part1(filename string, nConnections int) int {
	jumpBoxes, distances := calulateDistances(filename)

	connections := distances[:nConnections]
	adj := adjacencyMatrix(connections)

	visited := make(map[JumpBox]struct{})
	circuits := [][]JumpBox{}

	for _, jumpBox := range jumpBoxes {
		if _, exists := visited[jumpBox]; !exists {
			circuit := connect(adj, visited, jumpBox, []JumpBox{})
			circuits = append(circuits, circuit)
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	// for _, circuit := range circuits {
	// 	fmt.Println(circuit)
	// }

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2(filename string) int {
	jumpBoxes, distances := calulateDistances(filename)

	// for _, distance := range distances {
	// 	fmt.Println(distance)
	// }

	for i, _ := range distances {
		adj := adjacencyMatrix(distances[:i])
		circuit := connect(adj, make(map[JumpBox]struct{}), jumpBoxes[0], []JumpBox{})
		// fmt.Println(i, len(circuit), len(jumpBoxes))
		if len(circuit) == len(jumpBoxes) {
			// fmt.Println(distance, distances[i])
			return distances[i - 1].from.X * distances[i-1].to.X
		}		
	}
	return 0
}

func calulateDistances(filename string) ([]JumpBox, []Distance) {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	jumpBoxes := []JumpBox{}
	for _, line := range lines {
		jumpBox := lineToJumpBox(line)
		jumpBoxes = append(jumpBoxes, jumpBox)
	}

	distances := []Distance{}
	for i, from := range jumpBoxes {
		for j := i + 1; j < len(jumpBoxes); j++ {
			to := jumpBoxes[j]
			val := math.Sqrt(float64((from.X-to.X)*(from.X-to.X) + (from.Y-to.Y)*(from.Y-to.Y) + (from.Z-to.Z)*(from.Z-to.Z)))
			distance := Distance{from: from, to: to, val: val}
			distances = append(distances, distance)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].val < distances[j].val
	})
	return jumpBoxes, distances
}

func adjacencyMatrix(connections []Distance) map[JumpBox][]JumpBox {
	adj := make(map[JumpBox][]JumpBox)
	for _, connection := range connections {

		if _, exists := adj[connection.from]; !exists {
			adj[connection.from] = []JumpBox{}
		}
		adj[connection.from] = append(adj[connection.from], connection.to)

		if _, exists := adj[connection.to]; !exists {
			adj[connection.to] = []JumpBox{}
		}
		adj[connection.to] = append(adj[connection.to], connection.from)
	}
	return adj
}

func connect(adj map[JumpBox][]JumpBox, visited map[JumpBox]struct{}, jumpBox JumpBox, circuit []JumpBox) []JumpBox {
	if _, exists := visited[jumpBox]; exists {
		return circuit
	}
	visited[jumpBox] = struct{}{}
	circuit = append(circuit, jumpBox)
	for _, connected := range adj[jumpBox] {
		circuit = connect(adj, visited, connected, circuit)
	}
	return circuit
}

func lineToJumpBox(line string) JumpBox {
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	jumpBox := JumpBox{X: x, Y: y, Z: z}
	return jumpBox
}

type Distance struct {
	from JumpBox
	to   JumpBox
	val  float64
}

type JumpBox struct {
	X int
	Y int
	Z int
}
