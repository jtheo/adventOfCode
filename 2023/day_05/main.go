package main

import (
	"fmt"
)

func main() {
	gardener := run()

	res := part1(gardener)
	fmt.Printf("Day 4, part 1. Result: %d\n\n", res)

	// res = part2(gardener)
	// fmt.Printf("Day 4, part 2. Result: %d\n", res)
}

func part1(a almanac) int {
	min := a.seeds[0]

	for _, s := range a.seeds {
		fmt.Printf("Seed: %d\n", s)
		src := s
		for _, gn := range guides {
			fmt.Printf("Passing %d to %s\n", src, gn)
			src = findRange(src, a.guides[gn])
			fmt.Println()
		}
		fmt.Printf("Last value: %d\n\n", src)

		if src < min {
			min = src
		}
		break
	}

	return min
}

func findRange(src int, gu []entry) int {
	for _, entry := range gu {
		fmt.Printf("testing %d vs %+v\n", src, entry)
		if src > entry.srcMin && src < entry.srcMax {
			d := src - entry.srcMin
			fmt.Printf("Found: Src: %2d | Dst: %d | Group: %+v\n", src, entry.dstMin+d, entry)
			return entry.dstMin + d
		}
	}
	fmt.Println("Found shit")
	return -1
}
