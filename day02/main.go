package main

import (
	"fmt"
	"log"
	"maxmil/aoc2025/utils"
	"strconv"
	"strings"
)

func main() {
	// if part1("input-test.txt") != 1227775554 {
	// 	log.Fatal("test failed")
	// }
	// if part1("input.txt") != 12586854255 {
	// 	log.Fatal("part1 failed")
	// }
	if part2("input-test.txt") != 4174379265 {
		log.Fatal("test failed")
	}
	part2("input.txt")
}

func part1(filename string) int {
	content := utils.ReadContent(filename)
	invalid := 0
	for _, r := range strings.Split(content, ",") {
		start, end, _ := strings.Cut(r, "-")
		i := sumInvalid(start, end, 2)
		fmt.Println(start, end, i)
		invalid += i
	}
	fmt.Println("total invalid", invalid)
	return invalid
}

func part2(filename string) int {
	content := utils.ReadContent(filename)
	invalid := []int{}
	for _, r := range strings.Split(content, ",") {
		start, end, _ := strings.Cut(r, "-")
		fmt.Println(start)
		invalidInRange := make(map[int]bool)
		for repetitions := 2; repetitions <= len(end); repetitions++ {
			found := findInvalidIds(start, end, repetitions)
			fmt.Println(start, end, repetitions, found)
			for k := range found {
				invalidInRange[found[k]] = true
			}
		}
		fmt.Println("invalid ids:", invalidInRange)
		fmt.Println()

		for k := range invalidInRange {
			invalid = append(invalid, k)
		}
	}
	sum := 0
	for _, v := range invalid {
		sum += v
	}
	fmt.Println("total invalid", sum)
	return sum
}

func findInvalidIds(start string, end string, repetitions int) []int {
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	invalid := []int{}

	first, _ := strconv.Atoi(start[:len(start)/repetitions])
	// fmt.Println("first", first)
	i := 0
	for {
		next, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(first+i), repetitions))
		// fmt.Println("next", next)
		if next < startInt {
			i++
			continue
		}
		if next > endInt {
			break
		}
		fmt.Println("invalid", next)
		invalid = append(invalid, next)
		i++
	}
	return invalid
}

func sumInvalid(start string, end string, repetitions int) int {
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	// if startInt > endInt {
	// 	return 0
	// }
	// if len(start)%2 != 0 {
	// 	return sumInvalid("1"+strings.Repeat("0", len(start)), end)
	// }
	invalid := 0

	first, _ := strconv.Atoi(start[:len(start)/repetitions])
	// fmt.Println("first", first)
	// second, _ := strconv.Atoi(start[len(start)/2:])
	i := 0
	for {
		next, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(first+i), repetitions))
		// fmt.Println("next", next)
		if next < startInt {
			i++
			continue
		}
		if next > endInt {
			break
		}
		fmt.Println("invalid", next)
		invalid += next
		i++
	}
	return invalid
}

// func sumInvalid(start string, end string) int {
// 	startInt, _:=strconv.Atoi(start)
// 	endInt, _:=strconv.Atoi(end)
// 	// if startInt > endInt {
// 	// 	return 0
// 	// }
// 	// if len(start)%2 != 0 {
// 	// 	return sumInvalid("1"+strings.Repeat("0", len(start)), end)
// 	// }
// 	invalid := 0

// 	first, _:= strconv.Atoi(start[:len(start)/2])
// 	// second, _ := strconv.Atoi(start[len(start)/2:])
// 	i := 0
// 	for {
// 		next, _ := strconv.Atoi(fmt.Sprintf("%d%d", first+i, first+i))
// 		// fmt.Println("next", next)
// 		if next < startInt {
// 			i++
// 			continue
// 		}
// 		if next > endInt {
// 			break
// 		}
// 		fmt.Println("invalid", next)
// 		invalid += next
// 		i++
// 	}
// 	return invalid
// }

//5293867403
//12586854255
