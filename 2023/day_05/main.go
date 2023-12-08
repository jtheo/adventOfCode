package main

import (
	"fmt"
)

func main() {
	gardener := run()

	res := part1(gardener)
	fmt.Printf("Day 5, part 1. Result: %d\n\n", res)

	res = part2(gardener)
	fmt.Printf("Day 5, part 2. Result: %d\n", res)
}

func part1(a almanac) int {
	min := a.seeds[0]

	for _, s := range a.seeds {
		src := s
		for _, gn := range guides {
			src = findRange(src, a.guides[gn])
		}
		if src < min {
			min = src
		}
	}

	return min
}

func part2(a almanac) int {
	min := a.seeds[0]
	fmt.Println("len seeds: ", len(a.seeds))

	i := 0
	for i < len(a.seeds) {
		from := a.seeds[i]
		to := a.seeds[i+1]

		for s := 0; s < to; s++ {
			src := from + s
			for _, gn := range guides {
				src = findRange(src, a.guides[gn])
			}
			if src < min {
				min = src
			}
		}
		i += 2
	}

	return min
}

func findRange(src int, gu []entry) int {
	res := src
	for _, entry := range gu {
		if src >= entry.srcMin && src <= entry.srcMax {
			d := src - entry.srcMin
			res = entry.dstMin + d
			break
		}
	}
	return res
}
