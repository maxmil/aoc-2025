package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part 1", part1("input.txt", 1000))
	fmt.Println("part 2", part2("input.txt"))
}

func part1(filename string, nConnections int) int {
	boxes, allConnections := getConnections(filename)

	connections := allConnections[:nConnections]
	adj := adjacencyMatrix(connections)

	visited := make(map[JunctionBox]struct{})
	circuits := [][]JunctionBox{}

	for _, jumpBox := range boxes {
		if _, exists := visited[jumpBox]; !exists {
			circuit := connect(adj, visited, jumpBox, []JunctionBox{})
			circuits = append(circuits, circuit)
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2(filename string) int {
	boxes, connections := getConnections(filename)

	for i := range connections {
		adj := adjacencyMatrix(connections[:i])
		circuit := connect(adj, make(map[JunctionBox]struct{}), boxes[0], []JunctionBox{})
		if len(circuit) == len(boxes) {
			return connections[i-1].from.X * connections[i-1].to.X
		}
	}
	return 0
}

func getConnections(filename string) ([]JunctionBox, []Connection) {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	boxes := []JunctionBox{}
	for _, line := range lines {
		box := lineToJunctionBox(line)
		boxes = append(boxes, box)
	}

	connections := []Connection{}
	for i, from := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			to := boxes[j]
			val := math.Sqrt(float64((from.X-to.X)*(from.X-to.X) + (from.Y-to.Y)*(from.Y-to.Y) + (from.Z-to.Z)*(from.Z-to.Z)))
			connection := Connection{from: from, to: to, distance: val}
			connections = append(connections, connection)
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})
	return boxes, connections
}

func adjacencyMatrix(connections []Connection) map[JunctionBox][]JunctionBox {
	adj := make(map[JunctionBox][]JunctionBox)
	for _, connection := range connections {

		if _, exists := adj[connection.from]; !exists {
			adj[connection.from] = []JunctionBox{}
		}
		adj[connection.from] = append(adj[connection.from], connection.to)

		if _, exists := adj[connection.to]; !exists {
			adj[connection.to] = []JunctionBox{}
		}

		adj[connection.to] = append(adj[connection.to], connection.from)
	}
	return adj
}

func connect(adj map[JunctionBox][]JunctionBox, visited map[JunctionBox]struct{}, junctionBox JunctionBox, circuit []JunctionBox) []JunctionBox {
	if _, exists := visited[junctionBox]; exists {
		return circuit
	}
	visited[junctionBox] = struct{}{}
	circuit = append(circuit, junctionBox)
	for _, connected := range adj[junctionBox] {
		circuit = connect(adj, visited, connected, circuit)
	}
	return circuit
}

func lineToJunctionBox(line string) JunctionBox {
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return JunctionBox{X: x, Y: y, Z: z}
}

type Connection struct {
	from JunctionBox
	to   JunctionBox
	distance  float64
}

type JunctionBox struct {
	X int
	Y int
	Z int
}
