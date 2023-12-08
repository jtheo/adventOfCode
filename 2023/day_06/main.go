package main

import (
	"fmt"
	"strconv"
)

func main() {
	r := run()

	res := part1(r)
	fmt.Printf("Day 6, part 1. Result: %d\n\n", res)

	res = part2(r)
	fmt.Printf("Day 6, part 2. Result: %d\n\n", res)
}

func part1(r races) int {
	solutions := make([]int, len(r))
	for c, x := range r {
		for i := 1; i < x.time; i++ {
			timeToRun := x.time - i
			distance := timeToRun * i
			if distance > x.distance {
				solutions[c]++
			}
		}
	}
	res := 1
	for _, s := range solutions {
		res *= s
	}
	return res
}

func part2(r races) int {
	tmpTime := ""
	tmpDist := ""
	for _, x := range r {
		tmpTime += fmt.Sprintf("%d", x.time)
		tmpDist += fmt.Sprintf("%d", x.distance)
	}
	time, _ := strconv.Atoi(tmpTime)
	dist, _ := strconv.Atoi(tmpDist)
	newRace := races{
		{time: time, distance: dist},
	}
	res := part1(newRace)
	return res
}
