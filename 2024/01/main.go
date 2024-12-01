package main

import (
	"fmt"
)

func main() {
	o := run()

	fmt.Printf("Running Part %d\n", o.part)
	switch o.part {
	case 1:
		part1(o)
	case 2:
		part2(o)
	}
}

func part1(o opts) {
	left, right := split(o)
	sum := 0
	sums := []int{}
	for i, l := range left {
		dist := l - right[i]
		if dist < 0 {
			dist = dist * -1
		}
		sum += dist
		sums = append(sums, dist)
	}
	fmt.Printf("Sum is %d\n", sum)
	if o.verbose {
		fmt.Printf("Sums: %v\n", sums)
	}
}

func part2(o opts) {
	left, right := split(o)
	similarity := []int{}
	for i, l := range left {
		if o.verbose {
			fmt.Printf("Index %d => %d\n", i, l)
		}
		found := false
		ntimes := 0
		for _, r := range right {
			if l != r && !found {
				continue
			}
			if l != r && found {
				break
			}
			ntimes++
		}
		simm := l * ntimes
		if o.verbose {
			fmt.Printf("%d -> l: %d - s: %d\n", i, l, simm)
		}
		similarity = append(similarity, simm)
	}
	sum := 0
	for _, s := range similarity {
		sum += s
	}
	fmt.Printf("Sum is %d\n", sum)
}
