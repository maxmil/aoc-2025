package main

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"math"
	"maxmil/aoc2025/util"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	// util.RunAndCheck(part1, "input-test.txt", 7)
	// util.Run(part1, "input.txt")
	util.RunAndCheck(part2, "input-test.txt", 33)
	util.Run(part2, "input-l14.txt")
}

func part1(filename string) int {
	machines := parseInput(filename)

	totalPressed := 0
	for m, machine := range machines {
		start := make([]int, len(machine.Lights))
		states := [][]int{start}
		for pressed := 1; true; pressed++ {
			found := false
			next := [][]int{}
			for _, button := range machine.Buttons {
				for _, state := range states {
					lights := make([]int, len(state))
					copy(lights, state)
					for _, i := range button {
						lights[i] = 1 - state[i]
					}
					// fmt.Println(button, state, lights)
					if slices.Equal(lights, machine.Lights) {
						if !found {
							totalPressed += pressed
						}
						found = true
						fmt.Printf("Found solution for machine %d after %d presses\n", m, pressed)
					}
					next = append(next, lights)
				}
			}
			if found {
				break
			}
			states = next
			fmt.Println()
		}
	}

	// fmt.Println(machines)
	return totalPressed
}

func part2(filename string) int {
	presses := 0
	for m, machine := range parseInput(filename) {
		slices.SortFunc(machine.Buttons, func(i, j Button) int {
			return cmp.Compare(j.sum(), i.sum())
		})
		p, _ := findMin(machine.Buttons, machine.Joltages, 0, math.MaxInt, make(map[string]int, 10000000), make([]int, len(machine.Joltages)))
		fmt.Printf("Solved for machine %d with %d presses\n", m, p)
		presses += p
	}
	return presses
}

// for each button subtract from joltages to slice,
// if any value < 0 don't copy joltage or button
// sort by distance from 0
// recurse

// terminal operation, if joltage is all zeros return presses
// if min presses defined
// if max joltage component + presses > min presses return nil

var start = time.Now()
var lapStart = start

func findMin(buttons []Button, joltages Joltages, pressed int, maxPresses int, seen map[string]int, end Joltages) (int, bool) {
	key := joltages.toKey()
	if cache, exists := seen[key]; exists && cache <= pressed {
		return 0, false
	} else if maxPresses < math.MaxInt && joltages.max()+pressed > maxPresses-1 {
		return 0, false
	} else if slices.Equal(joltages, end) {
		fmt.Println("Found", pressed)
		return pressed, true
	}

	seen[key] = pressed
	if len(seen)%1000000 == 0 {
		fmt.Printf("Seen %d in %f minutes (lap time %fs)\n", len(seen), time.Since(start).Minutes(), time.Since(lapStart).Seconds())
		lapStart = time.Now()
	}
	// fmt.Println(pressed, joltages)

	nextButtons := []Button{}
	nextJoltages := []Joltages{}
	for _, button := range buttons {
		skip := false
		joltageAfterPress := make([]int, len(joltages))
		copy(joltageAfterPress, joltages)
		for _, ind := range button {
			if joltageAfterPress[ind] == 0 {
				skip = true
				break
			}
			joltageAfterPress[ind] -= 1
		}
		if !skip {
			nextButtons = append(nextButtons, button)
			nextJoltages = append(nextJoltages, joltageAfterPress)
		}
	}

	// slices.SortFunc(nextJoltages, func(i, j Joltages) int {
	// 	return cmp.Compare(i.max(), j.max())
	// })

	minimum := maxPresses
	configured := false
	for _, joltages := range nextJoltages {
		m, c := findMin(nextButtons, joltages, pressed+1, minimum, seen, end)
		if c {
			configured = true
			minimum = min(minimum, m)
		}
	}

	return minimum, configured
}

func (j Joltages) toKey() string {
	b := make([]byte, len(j)*8)
	for i, v := range j {
		binary.LittleEndian.PutUint64(b[i*8:], uint64(v))
	}
	return string(b)
}

func (b Button) sum() int {
	sum := 0
	for _, value := range b {
		sum += value
	}
	return sum
}

func (b Joltages) max() int {
	value := 0
	for _, joltage := range b {
		value = max(value, joltage)
	}
	return value
}

func (b Joltages) min() int {
	value := math.MaxInt
	for _, joltage := range b {
		value = min(value, joltage)
	}
	return value
}

func (j Joltages) sum() int {
	sum := 0
	for _, joltage := range j {
		sum += joltage
	}
	return sum
}

func parseInput(filename string) []Machine {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	machines := make([]Machine, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")

		lightsStr := parts[0][1 : len(parts[0])-1]
		machines[i].Lights = make([]int, len(lightsStr))
		for j, ch := range lightsStr {
			if ch == '#' {
				machines[i].Lights[j] = 1
			}
		}

		for j := 1; j < len(parts)-1; j++ {
			btnStr := parts[j][1 : len(parts[j])-1]
			button := []int{}
			for _, p := range strings.Split(btnStr, ",") {
				var val int
				fmt.Sscanf(p, "%d", &val)
				button = append(button, val)
			}
			machines[i].Buttons = append(machines[i].Buttons, button)
		}

		lastPart := parts[len(parts)-1]
		joltsStr := lastPart[1 : len(lastPart)-1]
		for _, p := range strings.Split(joltsStr, ",") {
			var val int
			fmt.Sscanf(p, "%d", &val)
			machines[i].Joltages = append(machines[i].Joltages, val)
		}
	}

	return machines
}

type Joltages []int
type Button []int

type Machine struct {
	Lights   []int
	Buttons  []Button
	Joltages Joltages
}
