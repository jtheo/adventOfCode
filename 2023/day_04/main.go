package main

import (
	"fmt"
)

func main() {
	games := run()

	res := part1(games)
	fmt.Printf("Day 4, part 1. Result: %d\n\n", res)

	res = part2(games)
	fmt.Printf("Day 4, part 2. Result: %d\n", res)
}

func part1(games data) int {
	res := 0
	for _, c := range games {
		tmp := 0
		for _, n := range c.winning {
			if contains(n, c.played) {
				if tmp == 0 {
					tmp = 1
				} else {
					tmp *= 2
				}
				continue
			}
		}
		res += tmp
	}

	return res
}

func part2(m data) int {
	res := 0

	return res
}

func contains(n int, played []int) bool {
	for _, i := range played {
		if i == n {
			return true
		}
	}
	return false
}
